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
	getApiCheckV2Body  = `{"test":{"id":489,"name":"Appinspect login API","active":true, "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-16T15:47:43.730Z","updated_at":"2022-08-16T15:47:43.741Z","lastRunAt":"2024-03-07T00:47:43.741Z","lastRunStatus":"success","createdBy":"abc1234","updatedBy":"abc1234","location_ids":["aws-us-east-1"],"type":"api","deviceId":1,"requests":[{"configuration":{"name":"Login","url":"https://api.splunk.com/2.0/rest/login/splunk","requestMethod":"GET","headers":{},"body":null},"setup":[],"validations":[]}]}}`
	inputGetApiCheckV2 = verifyApiCheckV2Input(string(getApiCheckV2Body))
)

func TestGetApiCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/tests/api/489", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getApiCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
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

	if !reflect.DeepEqual(resp.Test.Deviceid, inputGetApiCheckV2.Test.Deviceid) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Deviceid, inputGetApiCheckV2.Test.Deviceid)
	}

	if !reflect.DeepEqual(resp.Test.Requests, inputGetApiCheckV2.Test.Requests) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Requests, inputGetApiCheckV2.Test.Requests)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputGetApiCheckV2.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputGetApiCheckV2.Test.Customproperties)
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
