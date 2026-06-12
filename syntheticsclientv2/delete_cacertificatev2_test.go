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
	"net/http"
	"testing"
)

func TestDeleteCaCertificateV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/cacerts/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.DeleteCaCertificateV2(1)
	if err != nil {
		t.Fatal(err)
	}
	if resp != http.StatusNoContent {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, http.StatusNoContent)
	}
}
