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
	getBrowserCheckV2Body  = `{"test":{"active":true,"advancedSettings":{"authentication":{"password":"password123","username":"myuser"},"cookies":[{"key":"qux","value":"qux","domain":"splunk.com","path":"/qux"}],"headers":[{"name":"Accept","value":"application/json","domain":"splunk.com"}],"verifyCertificates":true},"createdAt":"2022-09-14T14:35:37.801Z","device":{"id":1,"label":"iPhone","networkConnection":{"description":"Mobile LTE","downloadBandwidth":12000,"latency":70,"packetLoss":0,"uploadBandwidth":12000},"viewportHeight":844,"viewportWidth":375},"frequency":5,"id":1,"locationIds":["na-us-virginia"],"name":"My Test","schedulingStrategy":"round_robin","transactions":[{"name":"Example transaction","steps":[{"name":"element step","selector":".main","selectorType":"css","type":"click_element","waitForNav":true}]}],"type":"browser","updatedAt":"2022-09-14T14:35:38.099Z"}}`
	inputGetBrowserCheckV2 = verifyBrowserCheckV2Input(string(getBrowserCheckV2Body))
)

func TestGetBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/browser/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(getBrowserCheckV2Body))
	})

	resp, _, err := testClient.GetBrowserCheckV2(1)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Test.ID, inputGetBrowserCheckV2.Test.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ID, inputGetBrowserCheckV2.Test.ID)
	}

	if !reflect.DeepEqual(resp.Test.Name, inputGetBrowserCheckV2.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputGetBrowserCheckV2.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Type, inputGetBrowserCheckV2.Test.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Type, inputGetBrowserCheckV2.Test.Type)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputGetBrowserCheckV2.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputGetBrowserCheckV2.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputGetBrowserCheckV2.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputGetBrowserCheckV2.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.Createdat, inputGetBrowserCheckV2.Test.Createdat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Createdat, inputGetBrowserCheckV2.Test.Createdat)
	}

	if !reflect.DeepEqual(resp.Test.Updatedat, inputGetBrowserCheckV2.Test.Updatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Updatedat, inputGetBrowserCheckV2.Test.Updatedat)
	}

	if !reflect.DeepEqual(resp.Test.Device, inputGetBrowserCheckV2.Test.Device) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Device, inputGetBrowserCheckV2.Test.Device)
	}

	if !reflect.DeepEqual(resp.Test.Advancedsettings, inputGetBrowserCheckV2.Test.Advancedsettings) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Advancedsettings, inputGetBrowserCheckV2.Test.Advancedsettings)
	}

	if !reflect.DeepEqual(resp.Test.Transactions, inputGetBrowserCheckV2.Test.Transactions) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Transactions, inputGetBrowserCheckV2.Test.Transactions)
	}

}

func verifyBrowserCheckV2Input(stringInput string) *BrowserCheckV2Response {
	check := &BrowserCheckV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestLiveGetBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
  res, _, err := c.GetBrowserCheckV2(495)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
