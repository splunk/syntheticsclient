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
	getApiCheckV2Body  = `{"test":{"active":true,"createdAt":"0001-01-01T00:00:00Z","device":{"id":1,"label":"Desktop","networkConnection":{}},"frequency":5,"id":489,"name":"Appinspect login API","requests":[{"configuration":{"body":"","headers":{},"name":"Login","requestMethod":"GET","url":"https://api.splunk.com/2.0/rest/login/splunk"}}],"type":"api","updatedAt":"0001-01-01T00:00:00Z"}}`
	inputGetApiCheckV2 = verifyApiCheckV2Input(string(getApiCheckV2Body))
)

func TestGetApiCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/api/489", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(getApiCheckV2Body))
	})

	resp, _, err := testClient.GetApiCheckV2(489)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Test.ID, inputGetApiCheckV2.Test.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ID, inputGetApiCheckV2.Test.ID)
	}

	if !reflect.DeepEqual(resp.Test.Name, inputGetApiCheckV2.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputGetApiCheckV2.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Type, inputGetApiCheckV2.Test.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Type, inputGetApiCheckV2.Test.Type)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputGetApiCheckV2.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputGetApiCheckV2.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputGetApiCheckV2.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputGetApiCheckV2.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.Createdat, inputGetApiCheckV2.Test.Createdat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Createdat, inputGetApiCheckV2.Test.Createdat)
	}

	if !reflect.DeepEqual(resp.Test.Updatedat, inputGetApiCheckV2.Test.Updatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Updatedat, inputGetApiCheckV2.Test.Updatedat)
	}

	if !reflect.DeepEqual(resp.Test.Device, inputGetApiCheckV2.Test.Device) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Device, inputGetApiCheckV2.Test.Device)
	}

	if !reflect.DeepEqual(resp.Test.Requests, inputGetApiCheckV2.Test.Requests) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Requests, inputGetApiCheckV2.Test.Requests)
	}

}

func verifyApiCheckV2Input(stringInput string) *ApiCheckV2Response {
	check := &ApiCheckV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestLiveGetApiCheckV2(t *testing.T) {
	setup()
	defer teardown()

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
  res, _, err := c.GetApiCheckV2(489)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
