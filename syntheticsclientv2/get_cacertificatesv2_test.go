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
	"testing"
)

var (
	getCaCertificatesV2Body  = `{"cacerts":[{"id":1,"name":"foo","description":"My CA certificate","content":"<REDACTED>","fileExtension":"pem","filename":"ca_cert_file","expiresAt":"2026-09-14T14:35:37.801Z","createdAt":"2022-09-14T14:35:37.801Z","createdBy":"abcdefgh1234","updatedAt":"2022-09-14T14:35:38.099Z","updatedBy":"abcdefgh1234"},{"id":2,"name":"bar","description":"Another CA certificate","content":"<REDACTED>","fileExtension":"crt","filename":"ca_cert_file_2","expiresAt":"2027-09-14T14:35:37.801Z","createdAt":"2022-09-14T14:35:37.801Z","createdBy":"abcdefgh1234","updatedAt":"2022-09-14T14:35:38.099Z","updatedBy":"abcdefgh1234"}]}`
	inputGetCaCertificatesV2 = verifyCaCertificatesV2Input(string(getCaCertificatesV2Body))
)

func TestGetCaCertificatesV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/cacerts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getCaCertificatesV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetCaCertificatesV2()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.CaCerts, inputGetCaCertificatesV2.CaCerts) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.CaCerts, inputGetCaCertificatesV2.CaCerts)
	}
}

func verifyCaCertificatesV2Input(stringInput string) *CaCertificatesV2Response {
	check := &CaCertificatesV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}
