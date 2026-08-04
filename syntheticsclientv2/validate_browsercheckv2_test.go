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
	"testing"
)

func TestValidateNewBrowserCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := BrowserCheckV2Input{}
	err := json.Unmarshal([]byte(createBrowserCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/browser/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewBrowserCheckV2(&input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateNewBrowserCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := BrowserCheckV2Input{}
	err := json.Unmarshal([]byte(createBrowserCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/browser/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"startUrl":["can't be blank"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewBrowserCheckV2(&input)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Valid {
		t.Errorf("resp.Valid = %#v, want false", resp.Valid)
	}

	fieldErrors, err := resp.FieldErrors()
	if err != nil {
		t.Fatal(err)
	}
	if len(fieldErrors["startUrl"]) != 1 || fieldErrors["startUrl"][0] != "can't be blank" {
		t.Errorf("FieldErrors()[\"startUrl\"] = %#v, want [\"can't be blank\"]", fieldErrors["startUrl"])
	}
}

func TestValidateBrowserCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := BrowserCheckV2Input{}
	err := json.Unmarshal([]byte(createBrowserCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/browser/77/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateBrowserCheckV2(77, &input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateBrowserCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := BrowserCheckV2Input{}
	err := json.Unmarshal([]byte(createBrowserCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/browser/77/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"transactions":["is invalid"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateBrowserCheckV2(77, &input)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Valid {
		t.Errorf("resp.Valid = %#v, want false", resp.Valid)
	}

	fieldErrors, err := resp.FieldErrors()
	if err != nil {
		t.Fatal(err)
	}
	if len(fieldErrors["transactions"]) != 1 || fieldErrors["transactions"][0] != "is invalid" {
		t.Errorf("FieldErrors()[\"transactions\"] = %#v, want [\"is invalid\"]", fieldErrors["transactions"])
	}
}
