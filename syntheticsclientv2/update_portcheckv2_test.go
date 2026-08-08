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
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var (
	updatePortCheckV2Body  = `{"test":{"name":"splunk - port 443", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "type":"port","url":"","port":80,"protocol":"tcp","host":"www.splunk.com","location_ids":["aws-us-east-1"],"frequency":10,"scheduling_strategy":"round_robin","active":true}}`
	inputPortCheckV2Update = PortCheckV2Input{}
)

func TestUpdatePortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/port/1650", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updatePortCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updatePortCheckV2Body), &inputPortCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdatePortCheckV2(1650, &inputPortCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputPortCheckV2Update.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputPortCheckV2Update.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputPortCheckV2Update.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputPortCheckV2Update.Test.Customproperties)
	}
}

func TestUpdatePortCheckV2HandlesEmptyResponseBody(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/port/1650", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
		// Write empty body
		_, err := w.Write([]byte(""))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updatePortCheckV2Body), &inputPortCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdatePortCheckV2(1650, &inputPortCheckV2Update)

	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Errorf("expected non-nil response on empty body, got nil")
	}
}

func TestUpdatePortCheckV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/port/1650", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updatePortCheckV2Body), &inputPortCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdatePortCheckV2(1650, &inputPortCheckV2Update)

	if err == nil {
		t.Fatal("expected error on malformed JSON response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on error, got %#v", resp)
	}
	if !strings.Contains(err.Error(), "invalid character") {
		t.Errorf("expected JSON unmarshal error, got: %v", err)
	}
}

func TestUpdatePortCheckV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	err := json.Unmarshal([]byte(updatePortCheckV2Body), &inputPortCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = unreachableClient.UpdatePortCheckV2(1650, &inputPortCheckV2Update)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
