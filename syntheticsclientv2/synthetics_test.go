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
	"encoding/json"
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

func TestSanitizeRequestDumpRedactsHeadersAndSecretJSONFields(t *testing.T) {
	requestDump := "POST /certificates HTTP/1.1\r\nX-SF-TOKEN: token-123\r\nContent-Type: application/json\r\n\r\n" +
		`{"certificate":{"publicKey":{"content":"public-secret"},"privateKey":{"content":"private-secret","password":"password-secret"}}}`

	sanitized := sanitizeRequestDump(requestDump, "token-123", "/certificates")

	for _, secret := range []string{"token-123", "public-secret", "private-secret", "password-secret"} {
		if strings.Contains(sanitized, secret) {
			t.Fatalf("sanitized request dump leaked %q: %s", secret, sanitized)
		}
	}
	for _, redacted := range []string{`"content":"[REDACTED]"`, `"password":"[REDACTED]"`, "X-SF-TOKEN: [REDACTED]"} {
		if !strings.Contains(sanitized, redacted) {
			t.Fatalf("sanitized request dump missing %q: %s", redacted, sanitized)
		}
	}
}

func TestSanitizeRequestDumpRedactsHeaderCookieAndAuthenticationValues(t *testing.T) {
	requestDump := "PUT /tests/browser/123 HTTP/1.1\r\nX-Sf-Token: token-abc\r\n\r\n" +
		`{"test":{"advancedSettings":{"headers":[{"name":"Authorization","value":"Bearer secret"}],"cookies":[{"name":"session","value":"cookie-secret"}],"authentication":{"username":"auth-user-secret","password":"auth-secret"}}}}`

	sanitized := sanitizeRequestDump(requestDump, "token-abc", "/tests/browser/123")

	for _, secret := range []string{"token-abc", "Bearer secret", "cookie-secret", "auth-user-secret", "auth-secret"} {
		if strings.Contains(sanitized, secret) {
			t.Fatalf("sanitized request dump leaked %q: %s", secret, sanitized)
		}
	}
	for _, redacted := range []string{`"value":"[REDACTED]"`, `"username":"[REDACTED]"`, `"password":"[REDACTED]"`, "X-Sf-Token: [REDACTED]"} {
		if !strings.Contains(sanitized, redacted) {
			t.Fatalf("sanitized request dump missing %q: %s", redacted, sanitized)
		}
	}
}

func TestSanitizeRequestDumpRedactsAllAPIHeaderMapValues(t *testing.T) {
	requestDump := "POST /v2/tests/api HTTP/1.1\r\nHost: example.com\r\nX-SF-TOKEN: client-api-key\r\n\r\n" +
		`{"test":{"requests":[{"configuration":{"headers":{"X-SF-TOKEN":"request-token","Accept":"accept-json-secret","beep":"plain-header-secret"},"body":null}}]}}`

	sanitized := sanitizeRequestDump(requestDump, "client-api-key", "/v2/tests/api")

	for _, secret := range []string{"client-api-key", "request-token", "accept-json-secret", "plain-header-secret"} {
		if strings.Contains(sanitized, secret) {
			t.Fatalf("sanitized request dump leaked %q: %s", secret, sanitized)
		}
	}
	for _, redacted := range []string{`"X-SF-TOKEN":"[REDACTED]"`, `"Accept":"[REDACTED]"`, `"beep":"[REDACTED]"`} {
		if !strings.Contains(sanitized, redacted) {
			t.Fatalf("sanitized request dump missing %q: %s", redacted, sanitized)
		}
	}
}

func TestSanitizeRequestDumpRedactsAPIAndHTTPCheckBodyFields(t *testing.T) {
	requestDump := "POST /tests/api HTTP/1.1\r\nHost: example.com\r\nX-SF-TOKEN: client-api-key\r\n\r\n" +
		`{"test":{"requests":[{"configuration":{"name":"login","body":"client_secret=api-body-secret&password=api-password-secret"}}],"body":"password=http-body-secret"}}`

	sanitized := sanitizeRequestDump(requestDump, "client-api-key", "/tests/api")

	for _, secret := range []string{"client-api-key", "api-body-secret", "api-password-secret", "http-body-secret"} {
		if strings.Contains(sanitized, secret) {
			t.Fatalf("sanitized request dump leaked %q: %s", secret, sanitized)
		}
	}
	if got := strings.Count(sanitized, `"body":"[REDACTED]"`); got != 2 {
		t.Fatalf("expected both API and HTTP check body fields to be redacted, saw %d redactions in: %s", got, sanitized)
	}
}

func TestSanitizeRequestDumpRedactsURLQueryValues(t *testing.T) {
	requestDump := "POST /tests/http?trace=client-query-secret&limit=25 HTTP/1.1\r\nHost: example.com\r\nX-SF-TOKEN: client-api-key\r\n\r\n" +
		`{"test":{"startUrl":"https://browser-user:browser-password@browser.example/start?login_token=browser-start-token&continue=browser-start-continue","url":"https://target-user:target-password@target.example/login?token=url-token-secret&tenant=tenant-secret","callbackUrl":"https://callback-user:callback-password@callback.example/path","requests":[{"configuration":{"url":"https://api.example/search?api_key=api-url-secret&q=customer-secret"}}]}}`

	sanitized := sanitizeRequestDump(requestDump, "client-api-key", "/tests/http")

	for _, secret := range []string{"client-api-key", "client-query-secret", "25", "browser-user", "browser-password", "browser-start-token", "browser-start-continue", "target-user", "target-password", "url-token-secret", "tenant-secret", "callback-user", "callback-password", "api-url-secret", "customer-secret"} {
		if strings.Contains(sanitized, secret) {
			t.Fatalf("sanitized request dump leaked %q: %s", secret, sanitized)
		}
	}
	for _, redacted := range []string{
		"/tests/http?trace=[REDACTED]&limit=[REDACTED]",
		"[REDACTED]@browser.example",
		"[REDACTED]@target.example",
		"[REDACTED]@callback.example",
		"login_token=[REDACTED]",
		"continue=[REDACTED]",
		"token=[REDACTED]",
		"tenant=[REDACTED]",
		"api_key=[REDACTED]",
		"q=[REDACTED]",
	} {
		if !strings.Contains(sanitized, redacted) {
			t.Fatalf("sanitized request dump missing %q: %s", redacted, sanitized)
		}
	}
}

func TestMakePublicAPICallDoesNotExposeRawRequest(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	defer testServer.Close()

	testMux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	})

	apiKey := "secret-api-key"
	certificateContent := "private-certificate-material"
	testConfigurableClient := NewConfigurableClient(apiKey, "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})

	requestBody := `{"certificate":{"publicKey":{"content":"` + certificateContent + `"}}}`
	details, err := testConfigurableClient.makePublicAPICall("POST", "/certificates", bytes.NewBufferString(requestBody), nil)
	if err != nil {
		t.Fatalf("expected no error, but saw: %s", err.Error())
	}

	if details.RawRequest != nil {
		t.Fatal("expected RawRequest to be nil; RequestBody is the supported sanitized debug representation")
	}
	if strings.Contains(details.RequestBody, apiKey) {
		t.Fatalf("expected request details to redact API key, but found it in: %s", details.RequestBody)
	}
	if strings.Contains(details.RequestBody, certificateContent) {
		t.Fatalf("expected request details to redact certificate content, but found it in: %s", details.RequestBody)
	}
	if !strings.Contains(details.RequestBody, `"content":"[REDACTED]"`) {
		t.Fatalf("expected sanitized RequestBody to include redacted content field, but saw: %s", details.RequestBody)
	}
}

func TestMakePublicAPICallRedactsSensitiveFieldsInSanitizedResponseBody(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	defer testServer.Close()

	totpSecret := "totp-secret-material"
	certContent := "private-certificate-material"
	certPassword := "private-key-password"

	testMux.HandleFunc("/totps/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"totp":{"id":1,"name":"login-totp","secret":"` + totpSecret + `","digits":6,"privateKey":{"content":"` + certContent + `","password":"` + certPassword + `"}}}`))
	})

	testConfigurableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})

	details, err := testConfigurableClient.makePublicAPICall("GET", "/totps/1", bytes.NewBufferString("{}"), nil)
	if err != nil {
		t.Fatalf("expected no error, but saw: %s", err.Error())
	}

	for _, secret := range []string{totpSecret, certContent, certPassword} {
		if strings.Contains(details.SanitizedResponseBody, secret) {
			t.Fatalf("sanitized response body leaked %q: %s", secret, details.SanitizedResponseBody)
		}
	}
	for _, redacted := range []string{`"secret":"[REDACTED]"`, `"content":"[REDACTED]"`, `"password":"[REDACTED]"`} {
		if !strings.Contains(details.SanitizedResponseBody, redacted) {
			t.Fatalf("sanitized response body missing %q: %s", redacted, details.SanitizedResponseBody)
		}
	}
	if !strings.Contains(details.SanitizedResponseBody, `"digits":6`) {
		t.Fatalf("sanitized response body should preserve unrelated fields, but saw: %s", details.SanitizedResponseBody)
	}
	if !strings.Contains(details.SanitizedResponseBody, `"name":"login-totp"`) {
		t.Fatalf("sanitized response body should preserve unrelated fields, but saw: %s", details.SanitizedResponseBody)
	}

	if strings.Contains(details.ResponseBody, "[REDACTED]") {
		t.Fatalf("expected raw ResponseBody used for parsing to remain unredacted, but saw: %s", details.ResponseBody)
	}
	if !strings.Contains(details.ResponseBody, totpSecret) {
		t.Fatalf("expected raw ResponseBody to retain real secret for parse*Response call sites, but saw: %s", details.ResponseBody)
	}

	var parsed TotpVariableV2Response
	if err := json.Unmarshal([]byte(details.ResponseBody), &parsed); err != nil {
		t.Fatalf("expected raw ResponseBody to remain valid JSON for parsing, but saw error: %s", err.Error())
	}
	if parsed.Totp.Secret != totpSecret {
		t.Fatalf("expected parsed response to retain real secret, but saw: %s", parsed.Totp.Secret)
	}
}

func TestMakePublicAPICallSanitizedResponseBodyHandlesEmptyAndMalformedBodies(t *testing.T) {
	if got := sanitizeResponseBody(""); got != "" {
		t.Fatalf("expected empty response body to remain unchanged, but saw: %q", got)
	}
	if got := sanitizeResponseBody("   "); got != "   " {
		t.Fatalf("expected blank response body to remain unchanged, but saw: %q", got)
	}

	malformed := `{"totp":{"secret":"unterminated`
	if got := sanitizeResponseBody(malformed); got != "[REDACTED]" {
		t.Fatalf("expected malformed response body to be fully redacted, but saw: %q", got)
	}
}

func TestMakePublicAPICallRedactsErrorResponseDetails(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	defer testServer.Close()

	echoedPassword := "echoed-password-secret"
	echoedContent := "echoed-content-secret"

	testMux.HandleFunc("/tests", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"400","message":"validation failed","details":{"password":"` + echoedPassword + `","content":"` + echoedContent + `","field":"name"}}`))
	})

	testConfigurableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})

	_, err := testConfigurableClient.makePublicAPICall("POST", "/tests", bytes.NewBufferString(`{}`), nil)
	if err == nil {
		t.Fatal("expected an error for a 400 status code")
	}

	for _, secret := range []string{echoedPassword, echoedContent} {
		if strings.Contains(err.Error(), secret) {
			t.Fatalf("error response leaked %q: %s", secret, err.Error())
		}
	}
	for _, redacted := range []string{`"password":"[REDACTED]"`, `"content":"[REDACTED]"`} {
		if !strings.Contains(err.Error(), redacted) {
			t.Fatalf("error response missing %q: %s", redacted, err.Error())
		}
	}
	if !strings.Contains(err.Error(), `"field":"name"`) {
		t.Fatalf("error response should preserve unrelated fields, but saw: %s", err.Error())
	}
}
