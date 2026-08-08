//go:build unit_tests
// +build unit_tests

// Copyright 2026 Splunk, Inc.
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
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

var (
	updateClientCertificateV2Body  = `{"certificate":{"description":"updated certificate","domain":"api.example.com","publicKey":{"content":"base64-public","fileExtension":"pem","filename":"client.crt"},"privateKey":{"content":"base64-private","fileExtension":"pem","filename":"client.key","password":"key-password"}}}`
	inputClientCertificateV2Update = ClientCertificateV2UpdateInput{}
)

func TestUpdateClientCertificateV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/certificates/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		assertClientCertificateUpdateRequestBody(t, r)
		w.WriteHeader(http.StatusOK)
	})

	if err := json.Unmarshal([]byte(updateClientCertificateV2Body), &inputClientCertificateV2Update); err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateClientCertificateV2(123, &inputClientCertificateV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}

func TestUpdateClientCertificateV2DoesNotIncludeName(t *testing.T) {
	setup()
	defer teardown()

	description := "updated certificate"
	inputClientCertificateV2Update := ClientCertificateV2UpdateInput{}
	inputClientCertificateV2Update.Certificate.Description = &description

	testMux.HandleFunc("/certificates/124", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, requestCertificateFields := readClientCertificateRequestFields(t, r)
		if _, ok := requestCertificateFields["name"]; ok {
			t.Fatal("request body should not include certificate.name")
		}
		w.WriteHeader(http.StatusOK)
	})

	resp, _, err := testClient.UpdateClientCertificateV2(124, &inputClientCertificateV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}

func assertClientCertificateUpdateRequestBody(t *testing.T, r *http.Request) {
	t.Helper()

	requestBody, requestCertificateFields := readClientCertificateRequestFields(t, r)

	expectedFields := map[string]string{
		"description": `"updated certificate"`,
		"domain":      `"api.example.com"`,
		"publicKey":   `{"content":"base64-public","filename":"client.crt","fileExtension":"pem"}`,
		"privateKey":  `{"content":"base64-private","filename":"client.key","fileExtension":"pem","password":"key-password"}`,
	}
	if len(requestCertificateFields) != len(expectedFields) {
		t.Errorf("request body certificate field count %d want %d: %s", len(requestCertificateFields), len(expectedFields), requestBody)
	}
	for field, expected := range expectedFields {
		rawValue, ok := requestCertificateFields[field]
		if !ok {
			t.Errorf("request body missing certificate.%s", field)
			continue
		}
		if string(rawValue) != expected {
			t.Errorf("request body certificate.%s %s want %s", field, rawValue, expected)
		}
	}
	for _, field := range []string{"name", "id", "expiresAt", "createdAt", "createdBy", "updatedAt", "updatedBy"} {
		if _, ok := requestCertificateFields[field]; ok {
			t.Errorf("request body should not include certificate.%s", field)
		}
	}
}

func TestUpdateClientCertificateV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/certificates/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	if err := json.Unmarshal([]byte(updateClientCertificateV2Body), &inputClientCertificateV2Update); err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateClientCertificateV2(123, &inputClientCertificateV2Update)

	if err == nil {
		t.Fatal("expected error on malformed JSON response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on error, got %#v", resp)
	}
	if !strings.Contains(err.Error(), "invalid character") {
		t.Errorf("expected JSON unmarshal error, got: %v", err)
	}
}

func TestUpdateClientCertificateV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	if err := json.Unmarshal([]byte(updateClientCertificateV2Body), &inputClientCertificateV2Update); err != nil {
		t.Fatal(err)
	}

	_, _, err := unreachableClient.UpdateClientCertificateV2(123, &inputClientCertificateV2Update)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestUpdateClientCertificateV2HandlesEmptyResponseBody(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/certificates/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
		// Write empty response body
	})

	if err := json.Unmarshal([]byte(updateClientCertificateV2Body), &inputClientCertificateV2Update); err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateClientCertificateV2(123, &inputClientCertificateV2Update)

	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for empty body")
	}
}
