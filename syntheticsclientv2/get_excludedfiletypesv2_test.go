//go:build unit_tests
// +build unit_tests

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
	"net/http"
	"reflect"
	"testing"
)

var getExcludedFileTypesV2Body = `{"excludedFileTypes":["chartbeat","google_analytics"]}`

func TestGetExcludedFileTypesV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/excluded_file_types", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if r.URL.Path != "/excluded_file_types" {
			t.Fatalf("path = %q, want /excluded_file_types", r.URL.Path)
		}
		_, err := w.Write([]byte(getExcludedFileTypesV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.GetExcludedFileTypesV2()
	if err != nil {
		t.Fatal(err)
	}
	if details.StatusCode != http.StatusOK {
		t.Fatalf("status code = %d, want %d", details.StatusCode, http.StatusOK)
	}

	expected := []string{"chartbeat", "google_analytics"}
	if !reflect.DeepEqual(resp.ExcludedFileTypes, expected) {
		t.Fatalf("ExcludedFileTypes = %#v, want %#v", resp.ExcludedFileTypes, expected)
	}
}

func TestParseExcludedFileTypesV2Response(t *testing.T) {
	resp, err := parseExcludedFileTypesV2Response(getExcludedFileTypesV2Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"chartbeat", "google_analytics"}
	if !reflect.DeepEqual(resp.ExcludedFileTypes, expected) {
		t.Fatalf("ExcludedFileTypes = %#v, want %#v", resp.ExcludedFileTypes, expected)
	}
}

func TestGetExcludedFileTypesV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/excluded_file_types", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte("{not valid json"))
	})

	resp, details, err := testClient.GetExcludedFileTypesV2()
	if err == nil {
		t.Fatal("expected a parse error, but got none")
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	if resp != nil {
		t.Errorf("expected nil response, got %#v", resp)
	}
}

func TestGetExcludedFileTypesV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})
	_, _, err := unreachableClient.GetExcludedFileTypesV2()
	if err == nil {
		t.Fatal("expected a connection error")
	}
}

func TestParseExcludedFileTypesV2ResponseReturnsErrorOnMalformedJSON(t *testing.T) {
	resp, err := parseExcludedFileTypesV2Response("{not valid json")
	if err == nil {
		t.Fatal("expected a parse error, but got none")
	}
	if resp != nil {
		t.Errorf("expected nil response, got %#v", resp)
	}
}
