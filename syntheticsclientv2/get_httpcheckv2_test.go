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
)

var (
	getHttpCheckV2Body  = `{"test":{"customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "active":true,"advancedSettings":{"authentication":{"password":"password123","username":"myuser"},"cookies":[{"key":"qux","value":"qux","domain":"splunk.com","path":"/qux"}],"headers":[{"name":"Accept","value":"application/json","domain":"splunk.com"}],"verifyCertificates":true},"createdAt":"2022-09-14T14:35:37.801Z","device":{"id":1,"label":"iPhone","networkConnection":{"description":"Mobile LTE","downloadBandwidth":12000,"latency":70,"packetLoss":0,"uploadBandwidth":12000},"viewportHeight":844,"viewportWidth":375},"frequency":5,"id":1,"locationIds":["na-us-virginia"],"name":"My Test","schedulingStrategy":"round_robin","transactions":[{"name":"Example transaction","steps":[{"name":"element step","selector":".main","selectorType":"css","type":"click_element","waitForNav":true}]}],"type":"browser","updatedAt":"2022-09-14T14:35:38.099Z","lastRunAt":"2024-03-07T00:47:43.741Z","lastRunStatus":"success"}}`
	inputGetHttpCheckV2 = verifyHttpCheckV2Input(string(getHttpCheckV2Body))
)

func TestGetHttpCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getHttpCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetHttpCheckV2(1)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Test.ID, inputGetHttpCheckV2.Test.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ID, inputGetHttpCheckV2.Test.ID)
	}

	if !reflect.DeepEqual(resp.Test.Name, inputGetHttpCheckV2.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputGetHttpCheckV2.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Type, inputGetHttpCheckV2.Test.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Type, inputGetHttpCheckV2.Test.Type)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputGetHttpCheckV2.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputGetHttpCheckV2.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputGetHttpCheckV2.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputGetHttpCheckV2.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.HttpHeaders, inputGetHttpCheckV2.Test.HttpHeaders) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.HttpHeaders, inputGetHttpCheckV2.Test.HttpHeaders)
	}

	if !reflect.DeepEqual(resp.Test.Body, inputGetHttpCheckV2.Test.Body) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Body, inputGetHttpCheckV2.Test.Body)
	}

	if !reflect.DeepEqual(resp.Test.URL, inputGetHttpCheckV2.Test.URL) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.URL, inputGetHttpCheckV2.Test.URL)
	}

	if !reflect.DeepEqual(resp.Test.SchedulingStrategy, inputGetHttpCheckV2.Test.SchedulingStrategy) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.SchedulingStrategy, inputGetHttpCheckV2.Test.SchedulingStrategy)
	}

	if !reflect.DeepEqual(resp.Test.RequestMethod, inputGetHttpCheckV2.Test.RequestMethod) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.RequestMethod, inputGetHttpCheckV2.Test.RequestMethod)
	}

	if !reflect.DeepEqual(resp.Test.LocationIds, inputGetHttpCheckV2.Test.LocationIds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.LocationIds, inputGetHttpCheckV2.Test.LocationIds)
	}

	if !reflect.DeepEqual(resp.Test.UpdatedAt, inputGetHttpCheckV2.Test.UpdatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.UpdatedAt, inputGetHttpCheckV2.Test.UpdatedAt)
	}

	if !reflect.DeepEqual(resp.Test.CreatedAt, inputGetHttpCheckV2.Test.CreatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.CreatedAt, inputGetHttpCheckV2.Test.CreatedAt)
	}

}

func verifyHttpCheckV2Input(stringInput string) *HttpCheckV2Response {
	check := &HttpCheckV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}
