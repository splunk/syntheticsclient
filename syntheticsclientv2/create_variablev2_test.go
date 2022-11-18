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
	"net/http"
	"reflect"
	"testing"
	"fmt"
	"encoding/json"
	"os"
)

var (
	createVariableV2Body  = `{"variable":{"description":"My super awesome test variable","name":"food","secret":false,"value":"bar"}}`
	inputVariableV2Data = VariableV2Input{}
)

func TestCreateVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(createVariableV2Body))
	})

	json.Unmarshal([]byte(createVariableV2Body), &inputVariableV2Data)

	resp, _, err := testClient.CreateVariableV2(&inputVariableV2Data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Variable.ID, inputVariableV2Data.Variable.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.ID, inputVariableV2Data.Variable.ID)
	}

	if !reflect.DeepEqual(resp.Variable.Name, inputVariableV2Data.Variable.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Name, inputVariableV2Data.Variable.Name)
	}

	if !reflect.DeepEqual(resp.Variable.Description, inputVariableV2Data.Variable.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Description, inputVariableV2Data.Variable.Description)
	}

	if !reflect.DeepEqual(resp.Variable.Value, inputVariableV2Data.Variable.Value) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Value, inputVariableV2Data.Variable.Value)
	}

	if !reflect.DeepEqual(resp.Variable.Secret, inputVariableV2Data.Variable.Secret) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Variable.Secret, inputVariableV2Data.Variable.Secret)
	}

}


func TestLiveCreateVariableV2(t *testing.T) {
	setup()
	defer teardown()

	json.Unmarshal([]byte(createVariableV2Body), &inputVariableV2Data)

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)
	
	fmt.Println(c)
	fmt.Println(inputVariableV2Data)

	// Make the request with your check settings and print result
  res, reqDetail, err := c.CreateVariableV2(&inputVariableV2Data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reqDetail)
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}