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
	"strings"
	"testing"
)

func TestDeleteSslCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/ssl/19", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.DeleteSslCheckV2(19)
	if err != nil {
		t.Fatal(err)
	}
	if resp != http.StatusNoContent {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, http.StatusNoContent)
	}
}

func TestDeleteSslCheckV2ReturnsErrorOnNon2xxStatus(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/ssl/19", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusMultipleChoices)
	})

	resp, err := testClient.DeleteSslCheckV2(19)

	if err == nil {
		t.Fatal("expected error on non-2xx response, got nil")
	}
	if !strings.Contains(err.Error(), "Response code") {
		t.Errorf("expected 'Response code' in error message, got: %v", err)
	}
	if resp != http.StatusMultipleChoices {
		t.Errorf("expected status code %d, got %d", http.StatusMultipleChoices, resp)
	}
}

func TestDeleteSslCheckV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	_, err := unreachableClient.DeleteSslCheckV2(19)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
