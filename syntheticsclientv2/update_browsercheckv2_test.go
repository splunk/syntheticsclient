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

func browserCheckV2RequestTestPayload(t *testing.T, requestBody string) map[string]interface{} {
	t.Helper()

	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(requestBody), &payload); err != nil {
		t.Fatalf("expected request body to be valid JSON, but saw error %s for body: %s", err, requestBody)
	}

	testPayload, ok := payload["test"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected request body to contain test object, but saw: %s", requestBody)
	}

	return testPayload
}

func browserCheckV2RequestAdvancedSettingsPayload(t *testing.T, requestBody string) map[string]interface{} {
	t.Helper()

	testPayload := browserCheckV2RequestTestPayload(t, requestBody)
	if _, ok := testPayload["certificateIds"]; ok {
		t.Fatalf("expected certificateIds under advancedSettings, not directly under test: %s", requestBody)
	}

	advancedSettings, ok := testPayload["advancedSettings"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected request body to contain advancedSettings object, but saw: %s", requestBody)
	}

	return advancedSettings
}

func assertBrowserCheckV2RequestCertificateIDs(t *testing.T, advancedSettings map[string]interface{}, expected []int) {
	t.Helper()

	certificateIDs, ok := advancedSettings["certificateIds"].([]interface{})
	if !ok {
		t.Fatalf("expected advancedSettings.certificateIds array, but saw: %#v", advancedSettings["certificateIds"])
	}
	if len(certificateIDs) != len(expected) {
		t.Fatalf("expected %d certificateIds, but saw %d: %#v", len(expected), len(certificateIDs), certificateIDs)
	}
	for i, expectedID := range expected {
		actualID, ok := certificateIDs[i].(float64)
		if !ok || int(actualID) != expectedID {
			t.Fatalf("expected certificateIds[%d] to be %d, but saw %#v", i, expectedID, certificateIDs[i])
		}
	}
}

func TestUpdateBrowserCheckV2OmitsUnsetCertificateIDs(t *testing.T) {
	input := BrowserCheckV2Input{}
	input.Test.Name = "browser-without-certs"

	requestBody := captureUpdateBrowserCheckV2RequestBody(t, input)

	testPayload := browserCheckV2RequestTestPayload(t, requestBody)
	if got := testPayload["name"]; got != "browser-without-certs" {
		t.Fatalf("expected request body to preserve browser test name, but saw %#v in body: %s", got, requestBody)
	}
	if _, ok := testPayload["certificateIds"]; ok {
		t.Fatalf("expected request body to omit certificateIds directly under test, but saw: %s", requestBody)
	}
	if advancedSettings, ok := testPayload["advancedSettings"].(map[string]interface{}); ok {
		if _, ok := advancedSettings["certificateIds"]; ok {
			t.Fatalf("expected request body to omit unset advancedSettings.certificateIds, but saw: %s", requestBody)
		}
	}
}

func TestUpdateBrowserCheckV2SerializesEmptyCertificateIDs(t *testing.T) {
	input := BrowserCheckV2Input{}
	input.Test.Name = "browser-clear-certs"
	certificateIDs := []int{}
	input.Test.CertificateIDs = certificateIDs

	requestBody := captureUpdateBrowserCheckV2RequestBody(t, input)

	testPayload := browserCheckV2RequestTestPayload(t, requestBody)
	if got := testPayload["name"]; got != "browser-clear-certs" {
		t.Fatalf("expected request body to preserve browser test name, but saw %#v in body: %s", got, requestBody)
	}
	advancedSettings := browserCheckV2RequestAdvancedSettingsPayload(t, requestBody)
	assertBrowserCheckV2RequestCertificateIDs(t, advancedSettings, []int{})
}

func TestUpdateBrowserCheckV2SerializesCertificateIDs(t *testing.T) {
	input := BrowserCheckV2Input{}
	input.Test.Name = "browser-attach-certs"
	certificateIDs := []int{123}
	input.Test.CertificateIDs = certificateIDs

	requestBody := captureUpdateBrowserCheckV2RequestBody(t, input)

	testPayload := browserCheckV2RequestTestPayload(t, requestBody)
	if got := testPayload["name"]; got != "browser-attach-certs" {
		t.Fatalf("expected request body to preserve browser test name, but saw %#v in body: %s", got, requestBody)
	}
	advancedSettings := browserCheckV2RequestAdvancedSettingsPayload(t, requestBody)
	assertBrowserCheckV2RequestCertificateIDs(t, advancedSettings, []int{123})
}

func TestUpdateBrowserCheckV2PreservesAdvancedSettingsFields(t *testing.T) {
	userAgent := "syntheticsclient-test-agent"
	input := BrowserCheckV2Input{}
	input.Test.Name = "browser-advanced-settings"
	input.Test.Authentication = &Authentication{
		Username: "browser-user",
		Password: "browser-password",
	}
	input.Test.Cookiesv2 = []Cookiesv2{
		{
			Key:    "session",
			Value:  "cookie-value",
			Domain: "example.com",
			Path:   "/",
		},
	}
	input.Test.BrowserHeaders = []BrowserHeaders{
		{
			Name:   "X-Test-Header",
			Value:  "header-value",
			Domain: "example.com",
		},
	}
	input.Test.HostOverrides = []HostOverrides{
		{
			Source:         "source.example.com",
			Target:         "target.example.com",
			KeepHostHeader: true,
		},
	}
	input.Test.UserAgent = &userAgent
	input.Test.CollectInteractiveMetrics = true
	input.Test.Verifycertificates = true
	input.Test.ChromeFlags = []ChromeFlag{
		{
			Name:  "--proxy-bypass-list",
			Value: "127.0.0.1:8080",
		},
	}
	input.Test.ExcludedFiles = []ExcludedFile{
		{
			Type:  "custom",
			Regex: "example.com",
		},
	}

	requestBody := captureUpdateBrowserCheckV2RequestBody(t, input)

	advancedSettings := browserCheckV2RequestAdvancedSettingsPayload(t, requestBody)
	expectedAdvancedSettingsJSON, err := json.Marshal(input.Test.Advancedsettings)
	if err != nil {
		t.Fatal(err)
	}
	var expectedAdvancedSettings map[string]interface{}
	if err := json.Unmarshal(expectedAdvancedSettingsJSON, &expectedAdvancedSettings); err != nil {
		t.Fatal(err)
	}

	for key, expected := range expectedAdvancedSettings {
		actual, ok := advancedSettings[key]
		if !ok {
			t.Fatalf("expected advancedSettings.%s to be preserved in request body: %s", key, requestBody)
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected advancedSettings.%s to be %#v, but saw %#v in body: %s", key, expected, actual, requestBody)
		}
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
