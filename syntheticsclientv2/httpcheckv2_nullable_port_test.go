//go:build unit_tests
// +build unit_tests

// Copyright 2021 Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syntheticsclientv2

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestCreateHttpCheckV2WithNullablePortSendsNullPort(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fields := readHttpV2NullablePortRequestFields(t, r)
		rawPort, ok := fields["port"]
		if !ok {
			t.Fatal("request body missing test.port")
		}
		if string(rawPort) != "null" {
			t.Fatalf("request body test.port = %s, want null", rawPort)
		}
		_, err := w.Write([]byte(`{"test":{"id":11,"name":"nullable-port","type":"http","url":"https://example.com","requestMethod":"GET","port":null}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	input := minimalHttpCheckV2InputWithNullablePort()
	input.Test.Port = *NewNullInt()

	resp, _, err := testClient.CreateHttpCheckV2WithNullablePort(&input)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Test.Port.Value != nil {
		t.Fatalf("response port = %#v, want nil", resp.Test.Port.Value)
	}
}

func TestUpdateHttpCheckV2WithNullablePortSendsZeroPort(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fields := readHttpV2NullablePortRequestFields(t, r)
		rawPort, ok := fields["port"]
		if !ok {
			t.Fatal("request body missing test.port")
		}
		if string(rawPort) != "0" {
			t.Fatalf("request body test.port = %s, want 0", rawPort)
		}
		_, err := w.Write([]byte(`{"test":{"id":12,"name":"nullable-port","type":"http","url":"https://example.com","requestMethod":"GET","port":0}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	input := minimalHttpCheckV2InputWithNullablePort()
	input.Test.Port = *NewNullableInt(0)

	resp, _, err := testClient.UpdateHttpCheckV2WithNullablePort(12, &input)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Test.Port.Value == nil || *resp.Test.Port.Value != 0 {
		t.Fatalf("response port = %#v, want 0", resp.Test.Port.Value)
	}
}

func TestGetHttpCheckV2WithNullablePortPreservesNullAndValue(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantNil   bool
		wantValue int
	}{
		{
			name:    "null",
			body:    `{"test":{"id":13,"name":"nullable-port","type":"http","url":"https://example.com","requestMethod":"GET","port":null}}`,
			wantNil: true,
		},
		{
			name:      "value",
			body:      `{"test":{"id":13,"name":"nullable-port","type":"http","url":"https://example.com","requestMethod":"GET","port":443}}`,
			wantValue: 443,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
			defer teardown()

			testMux.HandleFunc("/tests/http/13", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				_, err := w.Write([]byte(tt.body))
				if err != nil {
					t.Fatal(err)
				}
			})

			resp, _, err := testClient.GetHttpCheckV2WithNullablePort(13)
			if err != nil {
				t.Fatal(err)
			}
			if tt.wantNil {
				if resp.Test.Port.Value != nil {
					t.Fatalf("response port = %#v, want nil", resp.Test.Port.Value)
				}
				return
			}
			if resp.Test.Port.Value == nil || *resp.Test.Port.Value != tt.wantValue {
				t.Fatalf("response port = %#v, want %d", resp.Test.Port.Value, tt.wantValue)
			}
		})
	}
}

func minimalHttpCheckV2InputWithNullablePort() HttpCheckV2InputWithNullablePort {
	input := HttpCheckV2InputWithNullablePort{}
	input.Test.Name = "nullable-port"
	input.Test.Type = "http"
	input.Test.URL = "https://example.com"
	input.Test.LocationIds = []string{"aws-us-east-1"}
	input.Test.Frequency = 5
	input.Test.SchedulingStrategy = "round_robin"
	input.Test.Active = true
	input.Test.RequestMethod = "GET"
	input.Test.Verifycertificates = true
	input.Test.Validations = []Validations{}
	input.Test.Customproperties = []CustomProperties{}
	return input
}

func readHttpV2NullablePortRequestFields(t *testing.T, r *http.Request) map[string]json.RawMessage {
	t.Helper()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(requestBody, &envelope); err != nil {
		t.Fatalf("request body JSON parse failed: %v", err)
	}

	rawTest, ok := envelope["test"]
	if !ok {
		t.Fatalf("request body missing test envelope: %s", requestBody)
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(rawTest, &fields); err != nil {
		t.Fatalf("request body test JSON parse failed: %v", err)
	}

	return fields
}
