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
	"reflect"
	"strings"
	"testing"
)

var (
	getCaCertificateV2Body  = `{"cacert":{"id":1,"name":"foo","description":"My CA certificate","content":"<REDACTED>","fileExtension":"pem","filename":"ca_cert_file","expiresAt":"2026-09-14T14:35:37.801Z","createdAt":"2022-09-14T14:35:37.801Z","createdBy":"abcdefgh1234","updatedAt":"2022-09-14T14:35:38.099Z","updatedBy":"abcdefgh1234"}}`
	inputGetCaCertificateV2 = verifyCaCertificateV2Input(string(getCaCertificateV2Body))
)

func TestGetCaCertificateV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/cacerts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getCaCertificateV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetCaCertificateV2(1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.CaCert, inputGetCaCertificateV2.CaCert) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.CaCert, inputGetCaCertificateV2.CaCert)
	}
}

func verifyCaCertificateV2Input(stringInput string) *CaCertificateV2Response {
	check := &CaCertificateV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestGetCaCertificateV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/cacerts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetCaCertificateV2(1)

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

func TestGetCaCertificateV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	_, _, err := unreachableClient.GetCaCertificateV2(1)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
