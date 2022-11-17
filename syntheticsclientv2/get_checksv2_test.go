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
	getChecksV2Body  = ``
	inputGetChecksV2 = verifyChecksV2Input(string(getChecksV2Body))
)

// func TestGetChecksV2(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	testMux.HandleFunc("/tests/browser/1", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		w.Write([]byte(getChecksV2Body))
// 	})

// 	resp, _, err := testClient.GetChecksV2(1)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !reflect.DeepEqual(resp.Test.ID, inputGetChecksV2.Test.ID) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ID, inputGetChecksV2.Test.ID)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Name, inputGetChecksV2.Test.Name) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputGetChecksV2.Test.Name)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Type, inputGetChecksV2.Test.Type) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Type, inputGetChecksV2.Test.Type)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Frequency, inputGetChecksV2.Test.Frequency) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputGetChecksV2.Test.Frequency)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Active, inputGetChecksV2.Test.Active) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputGetChecksV2.Test.Active)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Createdat, inputGetChecksV2.Test.Createdat) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Createdat, inputGetChecksV2.Test.Createdat)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Updatedat, inputGetChecksV2.Test.Updatedat) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Updatedat, inputGetChecksV2.Test.Updatedat)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Device, inputGetChecksV2.Test.Device) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Device, inputGetChecksV2.Test.Device)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Advancedsettings, inputGetChecksV2.Test.Advancedsettings) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Advancedsettings, inputGetChecksV2.Test.Advancedsettings)
// 	}

// 	if !reflect.DeepEqual(resp.Test.Transactions, inputGetChecksV2.Test.Transactions) {
// 		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Transactions, inputGetChecksV2.Test.Transactions)
// 	}

// }

func verifyChecksV2Input(stringInput string) *ChecksV2Response {
	check := &ChecksV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestLiveGetChecksV2(t *testing.T) {
	setup()
	defer teardown()

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
  res, _, err := c.GetChecksV2(495)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
