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
	createBrowserCheckV2Body  = `{"test":{"name":"browser-beep-test","business_transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://www.splunk.com","action":"go_to_url","wait_for_nav":true},{"name":"Nexter step","type":"click_element","selector_type":"id","wait_for_nav":false,"selector":"free-splunk-click-desktop"}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","location_ids":["aws-us-east-1"],"device_id":1,"frequency":5,"scheduling_strategy":"round_robin","active":true,"advanced_settings":{"verify_certificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}]}}}`
	inputBrowserCheckV2Data = BrowserCheckV2Input{}
)

func TestCreateBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/browser", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(createBrowserCheckV2Body))
	})

	json.Unmarshal([]byte(createBrowserCheckV2Body), &inputBrowserCheckV2Data)

	resp, _, err := testClient.CreateBrowserCheckV2(&inputBrowserCheckV2Data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputBrowserCheckV2Data.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputBrowserCheckV2Data.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputBrowserCheckV2Data.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputBrowserCheckV2Data.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.Locationids, inputBrowserCheckV2Data.Test.LocationIds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Locationids, inputBrowserCheckV2Data.Test.LocationIds)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputBrowserCheckV2Data.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputBrowserCheckV2Data.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.BusinessTransactions, inputBrowserCheckV2Data.Test.BusinessTransactions) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.BusinessTransactions, inputBrowserCheckV2Data.Test.BusinessTransactions)
	}

	if !reflect.DeepEqual(resp.Test.Advancedsettings, inputBrowserCheckV2Data.Test.Advancedsettings) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Advancedsettings, inputBrowserCheckV2Data.Test.Advancedsettings)
	}

	if !reflect.DeepEqual(resp.Test.Schedulingstrategy, inputBrowserCheckV2Data.Test.Schedulingstrategy) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Schedulingstrategy, inputBrowserCheckV2Data.Test.Schedulingstrategy)
	}

}


func TestLiveCreateBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	json.Unmarshal([]byte(createBrowserCheckV2Body), &inputBrowserCheckV2Data)

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)
	
	fmt.Println(c)
	fmt.Println(inputBrowserCheckV2Data)

	// Make the request with your check settings and print result
  res, reqDetail, err := c.CreateBrowserCheckV2(&inputBrowserCheckV2Data)
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