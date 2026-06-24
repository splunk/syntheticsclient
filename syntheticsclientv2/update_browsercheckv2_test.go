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
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var (
	updateBrowserCheckV2Body  = `{"test":{"automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "name":"browser-beep-test","transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url","options":{"url":"https://splunk.com"}},{"name":"click","type":"click_element","selectors":[{"type":"id","value":"clicky"}],"waitForNav":true,"waitForNavTimeout":2000},{"name":"fill in fieldz","type":"enter_value","selectors":[{"type":"id","value":"beep"}],"value":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"accept---Alert","type":"accept_alert"},{"name":"Select-Val-Index","type":"select_option","selectors":[{"type":"id","value":"selectionz"}],"optionSelectorType":"index","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-val-text","type":"select_option","selectors":[{"type":"id","value":"textzz"}],"optionSelectorType":"text","optionSelector":"sdad","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-Val-Val","type":"select_option","selectors":[{"type":"id","value":"valz"}],"optionSelectorType":"value","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Run JS","type":"run_javascript","value":"beeeeeeep","waitForNav":true,"waitForNavTimeout":2000},{"name":"Save as text","type":"store_variable_from_element","selectors":[{"type":"link","value":"beepval"}],"variableName":"{{env.terraform-test-foo-301}}"},{"name":"Save JS return Val","type":"store_variable_from_javascript","value":"sdasds","variableName":"{{env.terraform-test-foo-301}}","waitForNav":true,"waitForNavTimeout":2000}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":15,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"chromeFlags":[{"name":"--proxy-bypass-list","value":"127.0.0.1:8080"}],"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}],"excludedFiles":[{"type":"google_analytics"},{"type":"custom","regex":"some-domain.com"},{"type":"all_except","regex":"another-domain.com"}]}}}`
	inputBrowserCheckV2Update = BrowserCheckV2Input{}
)

func captureUpdateBrowserCheckV2RequestBody(t *testing.T, input BrowserCheckV2Input) string {
	t.Helper()

	setup()
	defer teardown()

	var requestBody string
	testMux.HandleFunc("/v2/tests/browser/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		requestBody = string(bodyBytes)

		_, err = w.Write([]byte(updateBrowserCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	_, _, err := testClient.UpdateBrowserCheckV2(10, &input)
	if err != nil {
		t.Fatal(err)
	}

	return requestBody
}

func TestUpdateBrowserCheckV2OmitsUnsetCertificateIDs(t *testing.T) {
	input := BrowserCheckV2Input{}
	input.Test.Name = "browser-without-certs"

	requestBody := captureUpdateBrowserCheckV2RequestBody(t, input)

	if strings.Contains(requestBody, "certificateIds") {
		t.Fatalf("expected request body to omit unset certificateIds, but saw: %s", requestBody)
	}
}

func TestUpdateBrowserCheckV2SerializesEmptyCertificateIDs(t *testing.T) {
	input := BrowserCheckV2Input{}
	input.Test.Name = "browser-clear-certs"
	certificateIDs := []int{}
	input.Test.CertificateIDs = certificateIDs

	requestBody := captureUpdateBrowserCheckV2RequestBody(t, input)

	if !strings.Contains(requestBody, `"certificateIds":[]`) {
		t.Fatalf("expected request body to include empty certificateIds array, but saw: %s", requestBody)
	}
	if strings.Contains(requestBody, `"certificateIds":null`) {
		t.Fatalf("expected empty certificateIds array, not null, but saw: %s", requestBody)
	}
}

func TestUpdateBrowserCheckV2SerializesCertificateIDs(t *testing.T) {
	input := BrowserCheckV2Input{}
	input.Test.Name = "browser-attach-certs"
	certificateIDs := []int{123}
	input.Test.CertificateIDs = certificateIDs

	requestBody := captureUpdateBrowserCheckV2RequestBody(t, input)

	if !strings.Contains(requestBody, `"certificateIds":[123]`) {
		t.Fatalf("expected request body to include populated certificateIds array, but saw: %s", requestBody)
	}
}

func TestUpdateBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/tests/browser/10", func(w http.ResponseWriter, r *http.Request) {
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

	if !reflect.DeepEqual(resp.Test.Advancedsettings, inputBrowserCheckV2Update.Test.Advancedsettings) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Advancedsettings, inputBrowserCheckV2Update.Test.Advancedsettings)
	}
}
