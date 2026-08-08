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
	getClientCertificateV2Body  = `{"certificate":{"id":123,"name":"mtls_api_example","description":"mTLS certificate","domain":"api.example.com","expiresAt":"2027-01-02T03:04:05Z","createdAt":"2026-01-02T03:04:05Z","createdBy":"boris@example.com","updatedAt":"2026-01-03T03:04:05Z","updatedBy":"boris@example.com","publicKey":{"id":501,"content":"<REDACTED>","filename":"client.crt","fileExtension":"pem"},"privateKey":{"id":502,"content":"<REDACTED>","filename":"client.key","fileExtension":"pem","password":"<REDACTED>"}}}`
	inputGetClientCertificateV2 = verifyClientCertificateV2Input(getClientCertificateV2Body)
)

func TestGetClientCertificateV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/certificates/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getClientCertificateV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetClientCertificateV2(123)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Certificate, inputGetClientCertificateV2.Certificate) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Certificate, inputGetClientCertificateV2.Certificate)
	}
}

func verifyClientCertificateV2Input(stringInput string) *ClientCertificateV2Response {
	check := &ClientCertificateV2Response{}
	if err := json.Unmarshal([]byte(stringInput), check); err != nil {
		panic(err)
	}
	return check
}

func TestGetClientCertificateV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/certificates/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetClientCertificateV2(123)

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

func TestGetClientCertificateV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	_, _, err := unreachableClient.GetClientCertificateV2(123)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
