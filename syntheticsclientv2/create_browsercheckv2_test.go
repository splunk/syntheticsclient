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
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var (
	createBrowserCheckV2Body = `{"test":{"automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "name":"browser-beep-test","transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url","options":{"url":"https://splunk.com"}},{"name":"click","type":"click_element","selectors":[{"type":"id","value":"clicky"}],"waitForNav":true,"waitForNavTimeout":2000},{"name":"fill in fieldz","type":"enter_value","selectors":[{"type":"id","value":"beep"}],"value":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"accept---Alert","type":"accept_alert"},{"name":"Select-Val-Index","type":"select_option","selectors":[{"type":"id","value":"selectionz"}],"optionSelectorType":"index","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-val-text","type":"select_option","selectors":[{"type":"id","value":"textzz"}],"optionSelectorType":"text","optionSelector":"sdad","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-Val-Val","type":"select_option","selectors":[{"type":"id","value":"valz"}],"optionSelectorType":"value","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Run JS","type":"run_javascript","value":"beeeeeeep","waitForNav":true,"waitForNavTimeout":2000},{"name":"Save as text","type":"store_variable_from_element","selectors":[{"type":"link","value":"beepval"}],"variableName":"{{env.terraform-test-foo-301}}"},{"name":"Save JS return Val","type":"store_variable_from_javascript","value":"sdasds","variableName":"{{env.terraform-test-foo-301}}","waitForNav":true,"waitForNavTimeout":2000}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":5,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"certificateIds":[123],"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"chromeFlags":[{"name":"--proxy-bypass-list","value":"127.0.0.1:8080"}],"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}],"excludedFiles":[{"type":"google_analytics"},{"type":"custom","regex":"some-domain.com"},{"type":"all_except","regex":"another-domain.com"}]}}}`
	inputBrowserCheckV2Data  = BrowserCheckV2Input{}
)

func TestCreateBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/tests/browser", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		requestBodyString := string(requestBody)
		testPayload := browserCheckV2RequestTestPayload(t, requestBodyString)
		if got := testPayload["name"]; got != "browser-beep-test" {
			t.Fatalf("expected request body to preserve browser test name, but saw %#v in body: %s", got, requestBodyString)
		}
		advancedSettings := browserCheckV2RequestAdvancedSettingsPayload(t, requestBodyString)
		assertBrowserCheckV2RequestCertificateIDs(t, advancedSettings, []int{123})
		if strings.Contains(requestBodyString, "certificate_ids") {
			t.Fatalf("request body contains internal certificate_ids field: %s", requestBodyString)
		}
		_, err = w.Write([]byte(createBrowserCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createBrowserCheckV2Body), &inputBrowserCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateBrowserCheckV2(&inputBrowserCheckV2Data)

	if err != nil {
		t.Fatal(err)
	}

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

	if !reflect.DeepEqual(resp.Test.Transactions, inputBrowserCheckV2Data.Test.Transactions) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Transactions, inputBrowserCheckV2Data.Test.Transactions)
	}

	if !reflect.DeepEqual(resp.Test.Advancedsettings, inputBrowserCheckV2Data.Test.Advancedsettings) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Advancedsettings, inputBrowserCheckV2Data.Test.Advancedsettings)
	}

	if !reflect.DeepEqual(resp.Test.Schedulingstrategy, inputBrowserCheckV2Data.Test.Schedulingstrategy) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Schedulingstrategy, inputBrowserCheckV2Data.Test.Schedulingstrategy)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputBrowserCheckV2Data.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputBrowserCheckV2Data.Test.Customproperties)
	}
}
