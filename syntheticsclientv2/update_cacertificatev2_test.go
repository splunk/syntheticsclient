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
	"io"
	"net/http"
	"testing"
)

var (
	updateCaCertificateV2Body  = `{"cacert":{"description":"Updated CA certificate","content":"Q2VydGlmaWNhdGU=","fileExtension":"pem","filename":"updated_ca_cert_file"}}`
	inputCaCertificateV2Update = CaCertificateV2UpdateInput{}
)

func TestUpdateCaCertificateV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/cacerts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		assertCaCertificateUpdateRequestBody(t, r)
		w.WriteHeader(http.StatusOK)
	})

	err := json.Unmarshal([]byte(updateCaCertificateV2Body), &inputCaCertificateV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateCaCertificateV2(1, &inputCaCertificateV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}

func TestUpdateCaCertificateV2AllowsEmptyDescription(t *testing.T) {
	setup()
	defer teardown()

	description := ""
	inputCaCertificateV2Update := CaCertificateV2UpdateInput{}
	inputCaCertificateV2Update.CaCert.Description = &description

	testMux.HandleFunc("/cacerts/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		assertSingleCaCertificateUpdateField(t, r, "description", `""`)
		w.WriteHeader(http.StatusOK)
	})

	resp, _, err := testClient.UpdateCaCertificateV2(2, &inputCaCertificateV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}

func TestUpdateCaCertificateV2DoesNotIncludeName(t *testing.T) {
	setup()
	defer teardown()

	description := "Updated CA certificate"
	inputCaCertificateV2Update := CaCertificateV2UpdateInput{}
	inputCaCertificateV2Update.CaCert.Description = &description

	testMux.HandleFunc("/cacerts/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		assertSingleCaCertificateUpdateField(t, r, "description", `"Updated CA certificate"`)
		w.WriteHeader(http.StatusOK)
	})

	resp, _, err := testClient.UpdateCaCertificateV2(3, &inputCaCertificateV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}

func assertCaCertificateUpdateRequestBody(t *testing.T, r *http.Request) {
	t.Helper()

	_, requestCaCertFields := readCaCertificateUpdateRequestFields(t, r)

	expectedFields := map[string]string{
		"description":   `"Updated CA certificate"`,
		"content":       `"Q2VydGlmaWNhdGU="`,
		"fileExtension": `"pem"`,
		"filename":      `"updated_ca_cert_file"`,
	}
	if len(requestCaCertFields) != len(expectedFields) {
		t.Errorf("request body cacert field count %d want %d", len(requestCaCertFields), len(expectedFields))
	}
	for field, expected := range expectedFields {
		rawValue, ok := requestCaCertFields[field]
		if !ok {
			t.Errorf("request body missing cacert.%s", field)
			continue
		}
		if string(rawValue) != expected {
			t.Errorf("request body cacert.%s %s want %s", field, rawValue, expected)
		}
	}
	for _, field := range []string{"name", "id", "expiresAt", "createdAt", "createdBy", "updatedAt", "updatedBy"} {
		if _, ok := requestCaCertFields[field]; ok {
			t.Errorf("request body should not include cacert.%s", field)
		}
	}
}

func assertSingleCaCertificateUpdateField(t *testing.T, r *http.Request, field string, expected string) {
	t.Helper()

	requestBody, requestCaCertFields := readCaCertificateUpdateRequestFields(t, r)
	if _, ok := requestCaCertFields["name"]; ok {
		t.Fatalf("request body should not include cacert.name: %s", requestBody)
	}
	if len(requestCaCertFields) != 1 {
		t.Fatalf("request body cacert field count %d want 1: %s", len(requestCaCertFields), requestBody)
	}

	rawValue, ok := requestCaCertFields[field]
	if !ok {
		t.Fatalf("request body missing cacert.%s", field)
	}
	if string(rawValue) != expected {
		t.Fatalf("request body cacert.%s %s want %s", field, rawValue, expected)
	}
}

func readCaCertificateUpdateRequestFields(t *testing.T, r *http.Request) ([]byte, map[string]json.RawMessage) {
	t.Helper()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	var requestEnvelope map[string]json.RawMessage
	err = json.Unmarshal(requestBody, &requestEnvelope)
	if err != nil {
		t.Fatal(err)
	}
	rawCaCert, ok := requestEnvelope["cacert"]
	if !ok {
		t.Fatal("request body missing cacert envelope")
	}

	var requestCaCertFields map[string]json.RawMessage
	err = json.Unmarshal(rawCaCert, &requestCaCertFields)
	if err != nil {
		t.Fatal(err)
	}

	return requestBody, requestCaCertFields
}
