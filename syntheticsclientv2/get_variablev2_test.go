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
	"net/http"
	"reflect"
	"testing"
)

var (
	getVariableV2Body   = `{"variable":{"id":1,"name":"beep-var","description":"","value":"<REDACTED>","secret":true,"created_at":"2022-11-16T14:55:09.480Z","updated_at":"2022-11-16T14:55:09.480Z"}}`
	inputGetVariableV2  = verifyVariableV2Input(string(getVariableV2Body))
	getVariablesV2Body  = `{"variables":[{"id":286,"name":"terraform-test-foo-301","description":"The most awesome variable. Full of snakes.","value":"barv3--oopsasdasd","secret":false,"created_at":"2022-12-07T16:48:10.105Z","updated_at":"2022-12-07T16:48:10.105Z"},{"id":398,"name":"foodz","description":"My super awesome test variable","value":"bar","secret":false,"created_at":"2023-01-12T18:51:26.790Z","updated_at":"2023-01-12T18:51:26.790Z"},{"id":246,"name":"beep-var","description":"","value":"<REDACTED>","secret":true,"created_at":"2022-11-16T14:55:09.480Z","updated_at":"2022-11-16T14:55:09.480Z"},{"id":255,"name":"food","description":"My super awesome test variable","value":"bar","secret":false,"created_at":"2022-11-18T14:32:35.135Z","updated_at":"2022-11-18T14:32:35.135Z"},{"id":268,"name":"foo-var-two-barv3","description":"The most awesome variable. Full of snakes and bats (and bugs!).","value":"barv3","secret":false,"created_at":"2022-12-06T14:34:10.524Z","updated_at":"2022-12-06T14:34:10.524Z"},{"id":278,"name":"foo-var-two-barv-30","description":"The most awesome variable. Full of snakes.","value":"barv3--oopsasdasd","secret":false,"created_at":"2022-12-06T19:33:59.589Z","updated_at":"2022-12-06T19:33:59.589Z"}]}`
	inputGetVariablesV2 = verifyVariablesV2Input(string(getVariablesV2Body))
)

func TestGetVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getVariableV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetVariableV2(1)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Variable.ID, inputGetVariableV2.Variable.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.ID, inputGetVariableV2.Variable.ID)
	}

	if !reflect.DeepEqual(resp.Variable.Name, inputGetVariableV2.Variable.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Name, inputGetVariableV2.Variable.Name)
	}

	if !reflect.DeepEqual(resp.Variable.Description, inputGetVariableV2.Variable.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Description, inputGetVariableV2.Variable.Description)
	}

	if !reflect.DeepEqual(resp.Variable.Value, inputGetVariableV2.Variable.Value) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Value, inputGetVariableV2.Variable.Value)
	}

	if !reflect.DeepEqual(resp.Variable.Secret, inputGetVariableV2.Variable.Secret) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Secret, inputGetVariableV2.Variable.Secret)
	}

	if !reflect.DeepEqual(resp.Variable.Createdat, inputGetVariableV2.Variable.Createdat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Createdat, inputGetVariableV2.Variable.Createdat)
	}

	if !reflect.DeepEqual(resp.Variable.Updatedat, inputGetVariableV2.Variable.Updatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Updatedat, inputGetVariableV2.Variable.Updatedat)
	}

}

func verifyVariableV2Input(stringInput string) *VariableV2Response {
	check := &VariableV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestGetVariablesV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getVariablesV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetVariablesV2()

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Variable[1].ID, inputGetVariablesV2.Variable[1].ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable[1].ID, inputGetVariablesV2.Variable[1].ID)
	}

	if !reflect.DeepEqual(resp.Variable[1].Name, inputGetVariablesV2.Variable[1].Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable[1].Name, inputGetVariablesV2.Variable[1].Name)
	}

	if !reflect.DeepEqual(resp.Variable[1].Description, inputGetVariablesV2.Variable[1].Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable[1].Description, inputGetVariablesV2.Variable[1].Description)
	}

	if !reflect.DeepEqual(resp.Variable[1].Value, inputGetVariablesV2.Variable[1].Value) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable[1].Value, inputGetVariablesV2.Variable[1].Value)
	}

	if !reflect.DeepEqual(resp.Variable[1].Secret, inputGetVariablesV2.Variable[1].Secret) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable[1].Secret, inputGetVariablesV2.Variable[1].Secret)
	}

	if !reflect.DeepEqual(resp.Variable[1].Createdat, inputGetVariablesV2.Variable[1].Createdat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable[1].Createdat, inputGetVariablesV2.Variable[1].Createdat)
	}

	if !reflect.DeepEqual(resp.Variable[1].Updatedat, inputGetVariablesV2.Variable[1].Updatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable[1].Updatedat, inputGetVariablesV2.Variable[1].Updatedat)
	}

}

func verifyVariablesV2Input(stringInput string) *VariablesV2Response {
	check := &VariablesV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}
