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

func TestValidateNewApiCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := ApiCheckV2Input{}
	err := json.Unmarshal([]byte(createApiV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/api/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.ValidateNewApiCheckV2(&input)
	if err != nil {
		t.Fatal(err)
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateNewApiCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := ApiCheckV2Input{}
	err := json.Unmarshal([]byte(createApiV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/api/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"name":["can't be blank"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewApiCheckV2(&input)
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
	if len(fieldErrors["name"]) != 1 || fieldErrors["name"][0] != "can't be blank" {
		t.Errorf("FieldErrors()[\"name\"] = %#v, want [\"can't be blank\"]", fieldErrors["name"])
	}
}

func TestValidateApiCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := ApiCheckV2Input{}
	err := json.Unmarshal([]byte(createApiV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/api/489/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateApiCheckV2(489, &input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateApiCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := ApiCheckV2Input{}
	err := json.Unmarshal([]byte(createApiV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/v2/tests/api/489/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"requests":["is invalid"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateApiCheckV2(489, &input)
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
	if len(fieldErrors["requests"]) != 1 || fieldErrors["requests"][0] != "is invalid" {
		t.Errorf("FieldErrors()[\"requests\"] = %#v, want [\"is invalid\"]", fieldErrors["requests"])
	}
}
