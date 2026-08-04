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

func TestValidateNewHttpCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := HttpCheckV2Input{}
	err := json.Unmarshal([]byte(createHttpCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/http/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewHttpCheckV2(&input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateNewHttpCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := HttpCheckV2Input{}
	err := json.Unmarshal([]byte(createHttpCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/http/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"url":["is not a valid URL"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewHttpCheckV2(&input)
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
	if len(fieldErrors["url"]) != 1 || fieldErrors["url"][0] != "is not a valid URL" {
		t.Errorf("FieldErrors()[\"url\"] = %#v, want [\"is not a valid URL\"]", fieldErrors["url"])
	}
}

func TestValidateHttpCheckV2Success(t *testing.T) {
	setup()
	defer teardown()

	input := HttpCheckV2Input{}
	err := json.Unmarshal([]byte(createHttpCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/http/21/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateHttpCheckV2(21, &input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateHttpCheckV2Failure(t *testing.T) {
	setup()
	defer teardown()

	input := HttpCheckV2Input{}
	err := json.Unmarshal([]byte(createHttpCheckV2Body), &input)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/http/21/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"requestMethod":["is not included in the list"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateHttpCheckV2(21, &input)
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
	if len(fieldErrors["requestMethod"]) != 1 || fieldErrors["requestMethod"][0] != "is not included in the list" {
		t.Errorf("FieldErrors()[\"requestMethod\"] = %#v, want [\"is not included in the list\"]", fieldErrors["requestMethod"])
	}
}

func TestValidateNewHttpCheckV2WithNullablePortSuccess(t *testing.T) {
	setup()
	defer teardown()

	input := minimalHttpCheckV2InputWithNullablePort()
	input.Test.Port = *NewNullInt()

	testMux.HandleFunc("/tests/http/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fields := readHttpV2NullablePortRequestFields(t, r)
		rawPort, ok := fields["port"]
		if !ok {
			t.Fatal("request body missing test.port")
		}
		if string(rawPort) != "null" {
			t.Fatalf("request body test.port = %s, want null", rawPort)
		}
		_, err := w.Write([]byte(`{"valid":true,"message":"Test is valid","details":[]}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateNewHttpCheckV2WithNullablePort(&input)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Errorf("resp.Valid = %#v, want true", resp.Valid)
	}
}

func TestValidateHttpCheckV2WithNullablePortFailure(t *testing.T) {
	setup()
	defer teardown()

	input := minimalHttpCheckV2InputWithNullablePort()
	input.Test.Port = *NewNullableInt(0)

	testMux.HandleFunc("/tests/http/22/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fields := readHttpV2NullablePortRequestFields(t, r)
		rawPort, ok := fields["port"]
		if !ok {
			t.Fatal("request body missing test.port")
		}
		if string(rawPort) != "0" {
			t.Fatalf("request body test.port = %s, want 0", rawPort)
		}
		_, err := w.Write([]byte(`{"valid":false,"message":"Test is invalid","details":{"port":["is not included in the list"]}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.ValidateHttpCheckV2WithNullablePort(22, &input)
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
