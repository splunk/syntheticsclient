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
	"testing"
)

func TestNewNullableString(t *testing.T) {
	n := NewNullableString("hello")
	if n.Value == nil || *n.Value != "hello" {
		t.Errorf("returned \n\n%#v want value \n\n%#v", n.Value, "hello")
	}

	body, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `"hello"` {
		t.Errorf("returned \n\n%#v want \n\n%#v", string(body), `"hello"`)
	}
}

func TestNullableStringUnmarshalJSONError(t *testing.T) {
	var n NullableString
	err := n.UnmarshalJSON([]byte(`123`))
	if err == nil {
		t.Fatal("expected an error unmarshalling a non-string value into NullableString")
	}
}

func TestNullableIntUnmarshalJSONError(t *testing.T) {
	var n NullableInt
	err := n.UnmarshalJSON([]byte(`"not-an-int"`))
	if err == nil {
		t.Fatal("expected an error unmarshalling a non-numeric value into NullableInt")
	}
}
