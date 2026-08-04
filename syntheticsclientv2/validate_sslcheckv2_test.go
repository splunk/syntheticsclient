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

func TestValidateNewSslCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := SslCheckV2Input{}
	err := json.Unmarshal([]byte(createSslCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/ssl/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewSslCheckV2(&input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateNewSslCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := SslCheckV2Input{}
	err := json.Unmarshal([]byte(createSslCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/ssl/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"host":["can't be blank"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewSslCheckV2(&input)
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

func TestValidateSslCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	active := false
	input := SslCheckV2UpdateInput{}
	input.Test.Active = &active

	testMux.HandleFunc("/tests/ssl/1655/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateSslCheckV2(1655, &input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateSslCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	port := 999999
	input := SslCheckV2UpdateInput{}
	input.Test.Port = &port

	testMux.HandleFunc("/tests/ssl/1656/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"port":["is not included in the list"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateSslCheckV2(1656, &input)
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
	if len(fieldErrors["port"]) != 1 || fieldErrors["port"][0] != "is not included in the list" {
		t.Errorf("FieldErrors()[\"port\"] = %#v, want [\"is not included in the list\"]", fieldErrors["port"])
	}
}
