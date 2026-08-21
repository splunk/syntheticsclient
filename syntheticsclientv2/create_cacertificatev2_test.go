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
	createCaCertificateV2Body         = `{"cacert":{"name":"foo","description":"My CA certificate","content":"Q2VydGlmaWNhdGU=","fileExtension":"pem","filename":"ca_cert_file"}}`
	createCaCertificateV2ResponseBody = `{"cacert":{"id":1,"name":"foo","description":"My CA certificate","content":"<REDACTED>","fileExtension":"pem","filename":"ca_cert_file","expiresAt":"2026-09-14T14:35:37.801Z","createdAt":"2022-09-14T14:35:37.801Z","createdBy":"abcdefgh1234","updatedAt":"2022-09-14T14:35:38.099Z","updatedBy":"abcdefgh1234"}}`
	inputCaCertificateV2Data          = CaCertificateV2Input{}
	outputCaCertificateV2Data         = CaCertificateV2Response{}
)

func TestCreateCaCertificateV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/cacerts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		var requestEnvelope map[string]json.RawMessage
		err = json.Unmarshal(requestBody, &requestEnvelope)
		if err != nil {
			t.Fatal(err)
		}
		if len(requestEnvelope) != 1 {
			t.Fatalf("request body envelope keys \n\n%#v want only cacert", requestEnvelope)
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
		writableFields := []string{"name", "description", "content", "fileExtension", "filename"}
		if len(requestCaCertFields) != len(writableFields) {
			t.Errorf("request body cacert field count \n\n%#v want \n\n%#v", len(requestCaCertFields), len(writableFields))
		}
		for _, field := range writableFields {
			if _, ok := requestCaCertFields[field]; !ok {
				t.Errorf("request body missing cacert.%s", field)
			}
		}
		for _, field := range []string{"id", "expiresAt", "createdAt", "createdBy", "updatedAt", "updatedBy"} {
			if _, ok := requestCaCertFields[field]; ok {
				t.Errorf("request body includes response-only cacert.%s", field)
			}
		}

		requestCaCertificateV2Data := CaCertificateV2Input{}
		err = json.Unmarshal(requestBody, &requestCaCertificateV2Data)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(requestCaCertificateV2Data.CaCert.Name, inputCaCertificateV2Data.CaCert.Name) {
			t.Errorf("request body name \n\n%#v want \n\n%#v", requestCaCertificateV2Data.CaCert.Name, inputCaCertificateV2Data.CaCert.Name)
		}
		if !reflect.DeepEqual(requestCaCertificateV2Data.CaCert.Description, inputCaCertificateV2Data.CaCert.Description) {
			t.Errorf("request body description \n\n%#v want \n\n%#v", requestCaCertificateV2Data.CaCert.Description, inputCaCertificateV2Data.CaCert.Description)
		}
		if !reflect.DeepEqual(requestCaCertificateV2Data.CaCert.Content, inputCaCertificateV2Data.CaCert.Content) {
			t.Errorf("request body content \n\n%#v want \n\n%#v", requestCaCertificateV2Data.CaCert.Content, inputCaCertificateV2Data.CaCert.Content)
		}
		if !reflect.DeepEqual(requestCaCertificateV2Data.CaCert.FileExtension, inputCaCertificateV2Data.CaCert.FileExtension) {
			t.Errorf("request body file extension \n\n%#v want \n\n%#v", requestCaCertificateV2Data.CaCert.FileExtension, inputCaCertificateV2Data.CaCert.FileExtension)
		}
		if !reflect.DeepEqual(requestCaCertificateV2Data.CaCert.Filename, inputCaCertificateV2Data.CaCert.Filename) {
			t.Errorf("request body filename \n\n%#v want \n\n%#v", requestCaCertificateV2Data.CaCert.Filename, inputCaCertificateV2Data.CaCert.Filename)
		}

		_, err = w.Write([]byte(createCaCertificateV2ResponseBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createCaCertificateV2Body), &inputCaCertificateV2Data)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal([]byte(createCaCertificateV2ResponseBody), &outputCaCertificateV2Data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateCaCertificateV2(&inputCaCertificateV2Data)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.CaCert, outputCaCertificateV2Data.CaCert) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.CaCert, outputCaCertificateV2Data.CaCert)
	}
}

func TestCreateCaCertificateV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/cacerts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createCaCertificateV2Body), &inputCaCertificateV2Data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateCaCertificateV2(&inputCaCertificateV2Data)

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

func TestCreateCaCertificateV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	err := json.Unmarshal([]byte(createCaCertificateV2Body), &inputCaCertificateV2Data)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = unreachableClient.CreateCaCertificateV2(&inputCaCertificateV2Data)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
