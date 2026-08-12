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
	"reflect"
	"strings"
	"testing"
)

var (
	createClientCertificateV2Body         = `{"certificate":{"name":"mtls_api_example","description":"mTLS certificate","domain":"api.example.com","publicKey":{"content":"base64-public","fileExtension":"pem","filename":"client.crt"},"privateKey":{"content":"base64-private","fileExtension":"pem","filename":"client.key","password":"key-password"}}}`
	createClientCertificateV2ResponseBody = `{"certificate":{"id":123,"name":"mtls_api_example","description":"mTLS certificate","domain":"api.example.com","expiresAt":"2027-01-02T03:04:05Z","createdAt":"2026-01-02T03:04:05Z","createdBy":"boris@example.com","updatedAt":"2026-01-03T03:04:05Z","updatedBy":"boris@example.com","publicKey":{"id":501,"content":"<REDACTED>","filename":"client.crt","fileExtension":"pem"},"privateKey":{"id":502,"content":"<REDACTED>","filename":"client.key","fileExtension":"pem","password":"<REDACTED>"}}}`
	inputClientCertificateV2Data          = ClientCertificateV2Input{}
	outputClientCertificateV2Data         = ClientCertificateV2Response{}
)

func TestCreateClientCertificateV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		requestBody, requestCertificateFields := readClientCertificateRequestFields(t, r)

		writableFields := []string{"name", "description", "domain", "publicKey", "privateKey"}
		if len(requestCertificateFields) != len(writableFields) {
			t.Errorf("request body certificate field count %d want %d: %s", len(requestCertificateFields), len(writableFields), requestBody)
		}
		for _, field := range writableFields {
			if _, ok := requestCertificateFields[field]; !ok {
				t.Errorf("request body missing certificate.%s", field)
			}
		}
		for _, field := range []string{"id", "expiresAt", "createdAt", "createdBy", "updatedAt", "updatedBy", "public_key", "private_key", "file_extension"} {
			if _, ok := requestCertificateFields[field]; ok {
				t.Errorf("request body includes invalid certificate.%s", field)
			}
		}

		requestClientCertificateV2Data := ClientCertificateV2Input{}
		if err := json.Unmarshal(requestBody, &requestClientCertificateV2Data); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(requestClientCertificateV2Data.Certificate, inputClientCertificateV2Data.Certificate) {
			t.Errorf("request body certificate \n\n%#v want \n\n%#v", requestClientCertificateV2Data.Certificate, inputClientCertificateV2Data.Certificate)
		}

		_, err := w.Write([]byte(createClientCertificateV2ResponseBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	if err := json.Unmarshal([]byte(createClientCertificateV2Body), &inputClientCertificateV2Data); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(createClientCertificateV2ResponseBody), &outputClientCertificateV2Data); err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateClientCertificateV2(&inputClientCertificateV2Data)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Certificate, outputClientCertificateV2Data.Certificate) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Certificate, outputClientCertificateV2Data.Certificate)
	}
}

func readClientCertificateRequestFields(t *testing.T, r *http.Request) ([]byte, map[string]json.RawMessage) {
	t.Helper()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	var requestEnvelope map[string]json.RawMessage
	if err := json.Unmarshal(requestBody, &requestEnvelope); err != nil {
		t.Fatal(err)
	}
	if len(requestEnvelope) != 1 {
		t.Fatalf("request body envelope keys %#v want only certificate", requestEnvelope)
	}
	rawCertificate, ok := requestEnvelope["certificate"]
	if !ok {
		t.Fatal("request body missing certificate envelope")
	}

	var requestCertificateFields map[string]json.RawMessage
	if err := json.Unmarshal(rawCertificate, &requestCertificateFields); err != nil {
		t.Fatal(err)
	}

	return requestBody, requestCertificateFields
}

func TestCreateClientCertificateV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/certificates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	if err := json.Unmarshal([]byte(createClientCertificateV2Body), &inputClientCertificateV2Data); err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateClientCertificateV2(&inputClientCertificateV2Data)

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

func TestCreateClientCertificateV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	if err := json.Unmarshal([]byte(createClientCertificateV2Body), &inputClientCertificateV2Data); err != nil {
		t.Fatal(err)
	}

	_, _, err := unreachableClient.CreateClientCertificateV2(&inputClientCertificateV2Data)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
