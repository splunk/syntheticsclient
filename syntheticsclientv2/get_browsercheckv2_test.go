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
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	getBrowserCheckV2Body  = `{"test":{"automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "active":true,"advancedSettings":{"authentication":{"password":"password123","username":"myuser"},"cookies":[{"key":"qux","value":"qux","domain":"splunk.com","path":"/qux"}],"headers":[{"name":"Accept","value":"application/json","domain":"splunk.com"}],"verifyCertificates":true,"excludedFiles": [{"type": "google_analytics"},{"type": "custom","regex": "some-domain.com"},{"type": "all_except","regex": "another-domain.com"}]},"createdAt":"2022-09-14T14:35:37.801Z","deviceId":1,"frequency":5,"id":1,"locationIds":["na-us-virginia"],"name":"My Test","schedulingStrategy":"round_robin","transactions":[{"name":"Example transaction","steps":[{"name":"element step","selectors":[{"type":"css","value":".main"}],"type":"click_element","waitForNav":true,"waitForNavTimeout":2000,"waitForNavTimeoutDefault":true,"maxWaitTime":10000,"maxWaitTimeDefault":true}]}],"type":"browser","updatedAt":"2022-09-14T14:35:38.099Z","lastRunAt":"2024-03-07T00:47:43.741Z","lastRunStatus":"success","createdBy":"abc1234","updatedBy":"abc1234"}}`
	inputGetBrowserCheckV2 = verifyBrowserCheckV2Input(string(getBrowserCheckV2Body))
	expectedBrowserCheckV2 = BrowserCheckV2Response{
		Test: BrowserCheckV2ResponseTest{
			Automaticretries: 1,
			Customproperties: []CustomProperties{
				{Key: "Test_Key", Value: "Test Custom Properties"},
			},
			Active: true,
			Advancedsettings: Advancedsettings{
				Authentication: &Authentication{
					Password: "password123",
					Username: "myuser",
				},
				Cookiesv2: []Cookiesv2{
					{Key: "qux", Value: "qux", Domain: "splunk.com", Path: "/qux"},
				},
				BrowserHeaders: []BrowserHeaders{
					{Name: "Accept", Value: "application/json", Domain: "splunk.com"},
				},
				ExcludedFiles: []ExcludedFile{
					{Type: "google_analytics"},
					{Type: "custom", Regex: "some-domain.com"},
					{Type: "all_except", Regex: "another-domain.com"},
				},
				Verifycertificates: true,
			},
			Createdat: time.Date(2022, 9, 14, 14, 35, 37, 801000000, time.UTC),
			Deviceid:  1,
			Frequency: 5,
			ID:                 1,
			Locationids:        []string{"na-us-virginia"},
			Name:               "My Test",
			Schedulingstrategy: "round_robin",
			Transactions: []Transactions{
				{
					Name: "Example transaction",
					StepsV2: []StepsV2{
						{
							Name:                     "element step",
							Selectors:                []Selector{{Type: "css", Value: ".main"}},
							Type:                     "click_element",
							WaitForNav:               true,
							WaitForNavTimeout:        2000,
							WaitForNavTimeoutDefault: true,
							MaxWaitTime:              10000,
							MaxWaitTimeDefault:       true,
						},
					},
				},
			},
			Type:          "browser",
			Updatedat:     time.Date(2022, 9, 14, 14, 35, 38, 99000000, time.UTC),
			Lastrunat:     time.Date(2024, 3, 7, 0, 47, 43, 741000000, time.UTC),
			Lastrunstatus: "success",
			Createdby:     "abc1234",
			Updatedby:     "abc1234",
		},
	}
)

func TestGetBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/tests/browser/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getBrowserCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetBrowserCheckV2(1)

	if err != nil {
		t.Fatal(err)
	}
	// verify the json is unmarshalled correctly
	if !reflect.DeepEqual(*inputGetBrowserCheckV2, expectedBrowserCheckV2) {
		t.Errorf("returned \n\n%#v want \n\n%#v", *inputGetBrowserCheckV2, expectedBrowserCheckV2)
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

	if !reflect.DeepEqual(resp.Test.Deviceid, inputGetBrowserCheckV2.Test.Deviceid) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Deviceid, inputGetBrowserCheckV2.Test.Deviceid)
	}

	if !reflect.DeepEqual(resp.Test.Advancedsettings, inputGetBrowserCheckV2.Test.Advancedsettings) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Advancedsettings, inputGetBrowserCheckV2.Test.Advancedsettings)
	}

	if !reflect.DeepEqual(resp.Test.Transactions, inputGetBrowserCheckV2.Test.Transactions) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Transactions, inputGetBrowserCheckV2.Test.Transactions)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputGetBrowserCheckV2.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputGetBrowserCheckV2.Test.Customproperties)
	}

	// verify the whole response
	if !reflect.DeepEqual(*resp, expectedBrowserCheckV2) {
		t.Errorf("returned \n\n%#v want \n\n%#v", *resp, expectedBrowserCheckV2)
	}
}

func verifyBrowserCheckV2Input(stringInput string) *BrowserCheckV2Response {
	check := &BrowserCheckV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	empty := BrowserCheckV2Response{}
	if reflect.DeepEqual(empty, *check) {
		panic("Unmarshal failed, empty struct returned")
	}
	return check
}
