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

func TestValidateNewPortCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := PortCheckV2Input{}
	err := json.Unmarshal([]byte(createPortCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/port/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewPortCheckV2(&input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateNewPortCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := PortCheckV2Input{}
	err := json.Unmarshal([]byte(createPortCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/port/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"port":["is not a number"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewPortCheckV2(&input)
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
	if len(fieldErrors["port"]) != 1 || fieldErrors["port"][0] != "is not a number" {
		t.Errorf("FieldErrors()[\"port\"] = %#v, want [\"is not a number\"]", fieldErrors["port"])
	}
}

func TestValidatePortCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := PortCheckV2Input{}
	err := json.Unmarshal([]byte(createPortCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/port/33/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidatePortCheckV2(33, &input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidatePortCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := PortCheckV2Input{}
	err := json.Unmarshal([]byte(createPortCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/port/33/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"host":["can't be blank"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidatePortCheckV2(33, &input)
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
	if len(fieldErrors["host"]) != 1 || fieldErrors["host"][0] != "can't be blank" {
		t.Errorf("FieldErrors()[\"host\"] = %#v, want [\"can't be blank\"]", fieldErrors["host"])
	}
}
