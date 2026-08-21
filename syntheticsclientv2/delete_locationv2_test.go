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
	"fmt"
	"net/http"
	"strings"
	"testing"
)

var (
	deleteLocationV2RespBody = ``
)

func TestDeleteLocationV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/locations/beep", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		_, err := w.Write([]byte(deleteLocationV2RespBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, err := testClient.DeleteLocationV2("beep")
	if err != nil {
		fmt.Println(resp)
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestDeleteLocationV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})
	_, err := unreachableClient.DeleteLocationV2("beep")
	if err == nil {
		t.Fatal("expected a connection error")
	}
}

func TestDeleteLocationV2ReturnsErrorOnNon2xxStatus(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/locations/beep", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusMultipleChoices)
	})

	resp, err := testClient.DeleteLocationV2("beep")
	if err == nil {
		t.Fatalf("expected an error for non-2xx status, but got none")
	}
	if !strings.Contains(err.Error(), "Response code") {
		t.Errorf("expected error message to contain 'Response code', got: %v", err)
	}
	if resp != http.StatusMultipleChoices {
		t.Errorf("expected status code %d, got %d", http.StatusMultipleChoices, resp)
	}
}
