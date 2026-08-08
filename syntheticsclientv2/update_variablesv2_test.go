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
	"testing"
)

var (
	updateVariableV2Body  = `{"variable":{"description":"My super awesome test variable","name":"foo2","secret":false,"value":"bar"}}`
	inputVariableV2Update = VariableV2Input{}
)

func TestUpdateVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateVariableV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateVariableV2Body), &inputVariableV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateVariableV2(10, &inputVariableV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Variable.Name, inputVariableV2Update.Variable.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Name, inputVariableV2Update.Variable.Name)
	}

	if !reflect.DeepEqual(resp.Variable.Description, inputVariableV2Update.Variable.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Description, inputVariableV2Update.Variable.Description)
	}

	if !reflect.DeepEqual(resp.Variable.Value, inputVariableV2Update.Variable.Value) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Value, inputVariableV2Update.Variable.Value)
	}

	if !reflect.DeepEqual(resp.Variable.Secret, inputVariableV2Update.Variable.Secret) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Secret, inputVariableV2Update.Variable.Secret)
	}

}

func TestUpdateVariableV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	err := json.Unmarshal([]byte(updateVariableV2Body), &inputVariableV2Update)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/variables/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{malformed response}"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.UpdateVariableV2(10, &inputVariableV2Update)
	if err == nil {
		t.Fatal("expected error on malformed response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on parse error, got %#v", resp)
	}
}

func TestUpdateVariableV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	inputData := VariableV2Input{
		Variable: Variable{
			Name:        "test-var",
			Value:       "test-value",
			Secret:      false,
			Description: "test description",
		},
	}

	_, _, err := unreachableClient.UpdateVariableV2(10, &inputData)
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestUpdateVariableV2ReturnsNonNilResponseOnEmptyBody(t *testing.T) {
	setup()
	defer teardown()

	err := json.Unmarshal([]byte(updateVariableV2Body), &inputVariableV2Update)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/variables/20", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(""))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.UpdateVariableV2(20, &inputVariableV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for empty body with 2xx status")
	}
}
