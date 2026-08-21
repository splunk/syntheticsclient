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
	createVariableV2Body = `{"variable":{"description":"My super awesome test variable","name":"food","secret":false,"value":"bar"}}`
	inputVariableV2Data  = VariableV2Input{}
)

func TestCreateVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(createVariableV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createVariableV2Body), &inputVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateVariableV2(&inputVariableV2Data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.ID, inputVariableV2Data.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ID, inputVariableV2Data.ID)
	}

	if !reflect.DeepEqual(resp.Name, inputVariableV2Data.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Name, inputVariableV2Data.Name)
	}

	if !reflect.DeepEqual(resp.Description, inputVariableV2Data.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Description, inputVariableV2Data.Description)
	}

	if !reflect.DeepEqual(resp.Value, inputVariableV2Data.Value) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Value, inputVariableV2Data.Value)
	}

	if !reflect.DeepEqual(resp.Secret, inputVariableV2Data.Secret) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Secret, inputVariableV2Data.Secret)
	}

}

func TestCreateVariableV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	err := json.Unmarshal([]byte(createVariableV2Body), &inputVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.CreateVariableV2(&inputVariableV2Data)
	if err == nil {
		t.Fatal("expected error on malformed response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on parse error, got %#v", resp)
	}
}

func TestCreateVariableV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	inputData := VariableV2Input{
		Variable: Variable{
			Name:        "test-var",
			Value:       "test-value",
			Secret:      false,
			Description: "test description",
		},
	}

	_, _, err := unreachableClient.CreateVariableV2(&inputData)
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
