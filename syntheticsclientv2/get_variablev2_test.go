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
	"os"
	"fmt"
)

var (
	getVariableV2Body  = `{"variable":{"id":1,"name":"beep-var","description":"","value":"<REDACTED>","secret":true,"created_at":"2022-11-16T14:55:09.480Z","updated_at":"2022-11-16T14:55:09.480Z"}}`
	inputGetVariableV2 = verifyVariableV2Input(string(getVariableV2Body))
)

func TestGetVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(getVariableV2Body))
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

func TestLiveGetVariableV2(t *testing.T) {
	setup()
	defer teardown()

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
  res, _, err := c.GetVariableV2(246)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
