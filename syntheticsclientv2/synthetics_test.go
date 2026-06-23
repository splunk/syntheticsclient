//go:build unit_tests
// +build unit_tests

// Copyright 2021 Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syntheticsclientv2

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	testClient *Client

	// server is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testClient = NewConfigurableClient("apiKey", "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})
	log.Printf("Client instantiated: %s", testClient.publicBaseURL)
}

func teardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestConfigurableClient(t *testing.T) {
	setup()
	defer teardown()

	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	args := ClientArgs{
		timeoutSeconds: 30,
		publicBaseUrl:  testServer.URL,
	}

	testConfigurableClient := NewConfigurableClient("snakedonut", "us0", args)
	log.Printf("Client instantiated: %s", testServer.URL)
	if testConfigurableClient.GetHTTPClient() == nil {
		t.Errorf("http client is nil")
	}
	if testConfigurableClient.apiKey != "snakedonut" {
		t.Errorf("returned \n\n%#v want \n\n%#v", testConfigurableClient.apiKey, "snakedonut")
	}
	if testConfigurableClient.realm != "us0" {
		t.Errorf("returned \n\n%#v want \n\n%#v", testConfigurableClient.realm, "us0")
	}
}

func TestConfigurableClientTimeout(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testMux.HandleFunc("/v2/tests/browser/12", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	})

	testConfigurableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{
		timeoutSeconds: 1,
		publicBaseUrl:  testServer.URL,
	})
	log.Printf("Client instantiated: %s", testServer.URL)
	_, _, err := testConfigurableClient.GetBrowserCheckV2(12)
	if !strings.Contains(err.Error(), "context deadline exceeded (Client.Timeout exceeded while awaiting headers)") {
		t.Errorf("expected to see timeout error, but saw: %s", err.Error())
	}
}

func TestConfigurableClientErrorStatusCode(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testMux.HandleFunc("/tests", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status":"404"}`))
	})

	testConfigurableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})
	details, err := testConfigurableClient.makePublicAPICall("GET", "/tests", nil, nil)

	if !strings.Contains(err.Error(), "404 Not Found") {
		t.Errorf("expected to see 404 error, but saw: %s", err.Error())
	}

	if details.StatusCode != http.StatusNotFound {
		t.Errorf("expected to see 404 status code, but saw: %d", details.StatusCode)
	}
}

func TestMakePublicAPICallRedactsAPIKeyFromRequestDetails(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	defer testServer.Close()

	testMux.HandleFunc("/tests", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	})

	apiKey := "secret-api-key"
	testConfigurableClient := NewConfigurableClient(apiKey, "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})

	details, err := testConfigurableClient.makePublicAPICall("POST", "/tests", bytes.NewBufferString(`{"name":"test"}`), nil)
	if err != nil {
		t.Fatalf("expected no error, but saw: %s", err.Error())
	}

	if strings.Contains(details.RequestBody, apiKey) {
		t.Fatalf("expected request details to redact API key, but found it in: %s", details.RequestBody)
	}

	if !strings.Contains(details.RequestBody, "[REDACTED]") {
		t.Fatalf("expected request details to contain redaction marker, but saw: %s", details.RequestBody)
	}
}

func TestMakePublicAPICallRedactsCaCertificateContentFromRequestDetails(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	defer testServer.Close()

	testMux.HandleFunc("/cacerts", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	})

	apiKey := "secret-api-key"
	caCertificateContent := "private-ca-material"
	testConfigurableClient := NewConfigurableClient(apiKey, "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})

	requestBody := `{"cacert":{"name":"test-ca","content":"` + caCertificateContent + `","metadata":{"Content":"` + caCertificateContent + `"}}}`
	details, err := testConfigurableClient.makePublicAPICall("POST", "/cacerts", bytes.NewBufferString(requestBody), nil)
	if err != nil {
		t.Fatalf("expected no error, but saw: %s", err.Error())
	}

	if strings.Contains(details.RequestBody, apiKey) {
		t.Fatalf("expected request details to redact API key, but found it in: %s", details.RequestBody)
	}
	if strings.Contains(details.RequestBody, caCertificateContent) {
		t.Fatalf("expected request details to redact CA certificate content, but found it in: %s", details.RequestBody)
	}
	if !strings.Contains(details.RequestBody, `"content":"[REDACTED]"`) {
		t.Fatalf("expected request details to preserve redacted content field, but saw: %s", details.RequestBody)
	}
	if !strings.Contains(details.RequestBody, `"Content":"[REDACTED]"`) {
		t.Fatalf("expected request details to redact content fields case-insensitively, but saw: %s", details.RequestBody)
	}
}

func TestMakePublicAPICallRedactsMalformedCaCertificateRequestBody(t *testing.T) {
	requestDump := "POST /cacerts HTTP/1.1\r\nHost: example.com\r\nX-Sf-Token: secret-api-key\r\n\r\n{\"cacert\":{\"content\":\"private-ca-material\""

	got := sanitizeRequestDump(requestDump, "secret-api-key", "/cacerts")

	if strings.Contains(got, "secret-api-key") {
		t.Fatalf("expected request details to redact API key, but found it in: %s", got)
	}
	if strings.Contains(got, "private-ca-material") {
		t.Fatalf("expected request details to redact malformed CA certificate body, but found it in: %s", got)
	}
	if !strings.Contains(got, "Host: example.com") {
		t.Fatalf("expected request details to preserve headers, but saw: %s", got)
	}
	if !strings.HasSuffix(got, "\r\n\r\n[REDACTED]") {
		t.Fatalf("expected request details to replace malformed body with redaction marker, but saw: %s", got)
	}
}

func TestMakePublicAPICallRedactsMalformedCaCertificateRequestBodyWithLFSeparator(t *testing.T) {
	requestDump := "POST /cacerts HTTP/1.1\nHost: example.com\nX-Sf-Token: secret-api-key\n\n{\"cacert\":{\"content\":\"private-ca-material\""

	got := sanitizeRequestDump(requestDump, "secret-api-key", "/cacerts")

	if strings.Contains(got, "secret-api-key") {
		t.Fatalf("expected request details to redact API key, but found it in: %s", got)
	}
	if strings.Contains(got, "private-ca-material") {
		t.Fatalf("expected request details to redact malformed CA certificate body, but found it in: %s", got)
	}
	if !strings.Contains(got, "Host: example.com") {
		t.Fatalf("expected request details to preserve headers, but saw: %s", got)
	}
	if !strings.HasSuffix(got, "\n\n[REDACTED]") {
		t.Fatalf("expected request details to preserve LF separator before redaction marker, but saw: %s", got)
	}
}

func TestCreateCaCertificateV2RedactsRequestDetails(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	defer testServer.Close()

	testMux.HandleFunc("/cacerts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"cacert":{"id":1,"name":"test-ca","description":"private test CA","content":"<REDACTED>","fileExtension":"pem","filename":"ca.pem"}}`))
	})

	apiKey := "secret-api-key"
	caCertificateContent := "private-ca-material"
	testConfigurableClient := NewConfigurableClient(apiKey, "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})
	input := CaCertificateV2Input{}
	input.CaCert.Name = "test-ca"
	input.CaCert.Description = "private test CA"
	input.CaCert.Content = caCertificateContent
	input.CaCert.FileExtension = "pem"
	input.CaCert.Filename = "ca.pem"

	_, details, err := testConfigurableClient.CreateCaCertificateV2(&input)
	if err != nil {
		t.Fatalf("expected no error, but saw: %s", err.Error())
	}
	if details == nil {
		t.Fatal("expected request details")
	}

	if strings.Contains(details.RequestBody, apiKey) {
		t.Fatalf("expected request details to redact API key, but found it in: %s", details.RequestBody)
	}
	if strings.Contains(details.RequestBody, caCertificateContent) {
		t.Fatalf("expected request details to redact CA certificate content, but found it in: %s", details.RequestBody)
	}
	if !strings.Contains(details.RequestBody, "X-Sf-Token: [REDACTED]") {
		t.Fatalf("expected request details to contain redacted API token header, but saw: %s", details.RequestBody)
	}
	if !strings.Contains(details.RequestBody, `"content":"[REDACTED]"`) {
		t.Fatalf("expected request details to contain redacted CA certificate content field, but saw: %s", details.RequestBody)
	}
}

func TestRedactSensitiveValueIgnoresEmptyValue(t *testing.T) {
	requestDump := "GET /tests HTTP/1.1\r\nX-Sf-Token: \r\n\r\n{}"

	if got := redactSensitiveValue(requestDump, ""); got != requestDump {
		t.Fatalf("expected empty sensitive value to leave request dump unchanged, but saw: %s", got)
	}
}
