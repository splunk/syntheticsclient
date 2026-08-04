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
	"reflect"
	"testing"
)

func TestValidateResponseFieldErrorsSuccess(t *testing.T) {
	resp, err := parseValidateResponse(`{"valid":true,"message":"Test is valid","details":[]}`)
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Valid {
		t.Fatalf("resp.Valid = %#v, want true", resp.Valid)
	}

	fieldErrors, err := resp.FieldErrors()
	if err != nil {
		t.Fatal(err)
	}
	if len(fieldErrors) != 0 {
		t.Errorf("FieldErrors() = %#v, want empty map", fieldErrors)
	}
}

func TestValidateResponseFieldErrorsFailure(t *testing.T) {
	resp, err := parseValidateResponse(`{"valid":false,"message":"Test is invalid","details":{"name":["can't be blank"],"frequency":["is not included in the list"]}}`)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Valid {
		t.Fatalf("resp.Valid = %#v, want false", resp.Valid)
	}

	fieldErrors, err := resp.FieldErrors()
	if err != nil {
		t.Fatal(err)
	}

	want := map[string][]string{
		"name":      {"can't be blank"},
		"frequency": {"is not included in the list"},
	}
	if !reflect.DeepEqual(fieldErrors, want) {
		t.Errorf("FieldErrors() = %#v, want %#v", fieldErrors, want)
	}
}

func TestValidateResponseBlankBody(t *testing.T) {
	resp, err := parseValidateResponse("")
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank body")
	}

	fieldErrors, err := resp.FieldErrors()
	if err != nil {
		t.Fatal(err)
	}
	if len(fieldErrors) != 0 {
		t.Errorf("FieldErrors() = %#v, want empty map", fieldErrors)
	}
}
