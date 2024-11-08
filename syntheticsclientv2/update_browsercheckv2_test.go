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
	updateBrowserCheckV2Body  = `{"test":{"automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "name":"browser-beep-test","transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url","options":{"url":"https://splunk.com"}},{"name":"click","type":"click_element","selectorType":"id","selector":"clicky","waitForNav":true,"waitForNavTimeout":2000},{"name":"fill in fieldz","type":"enter_value","selectorType":"id","selector":"beep","value":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"accept---Alert","type":"accept_alert"},{"name":"Select-Val-Index","type":"select_option","selectorType":"id","selector":"selectionz","optionSelectorType":"index","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-val-text","type":"select_option","selectorType":"id","selector":"textzz","optionSelectorType":"text","optionSelector":"sdad","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-Val-Val","type":"select_option","selectorType":"id","selector":"valz","optionSelectorType":"value","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Run JS","type":"run_javascript","value":"beeeeeeep","waitForNav":true,"waitForNavTimeout":2000},{"name":"Save as text","type":"store_variable_from_element","selectorType":"link","selector":"beepval","variableName":"{{env.terraform-test-foo-301}}"},{"name":"Save JS return Val","type":"store_variable_from_javascript","value":"sdasds","variableName":"{{env.terraform-test-foo-301}}","waitForNav":true,"waitForNavTimeout":2000}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":15,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"chromeFlags":[{"name":"--proxy-bypass-list","value":"127.0.0.1:8080"}],"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}]}}}`
	inputBrowserCheckV2Update = BrowserCheckV2Input{}
)

func TestUpdateBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/browser/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateBrowserCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateBrowserCheckV2Body), &inputBrowserCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateBrowserCheckV2(10, &inputBrowserCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputBrowserCheckV2Update.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputBrowserCheckV2Update.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputBrowserCheckV2Update.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputBrowserCheckV2Update.Test.Customproperties)
	}
}
