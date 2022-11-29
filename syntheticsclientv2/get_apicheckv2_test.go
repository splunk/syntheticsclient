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
	getApiCheckV2Body  = `{"test":{"id":489,"name":"Appinspect login API","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-16T15:47:43.730Z","updated_at":"2022-08-16T15:47:43.741Z","location_ids":["aws-us-east-1"],"type":"api","device":{"id":1,"label":"Desktop","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Splunk Synthetics) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36","viewport_width":1366,"viewport_height":768,"network_connection":{"description":"Standard Cable","upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":null}},"requests":[{"configuration":{"name":"Login","url":"https://api.splunk.com/2.0/rest/login/splunk","requestMethod":"GET","headers":{},"body":null},"setup":[],"validations":[]}]}}`
	inputGetApiCheckV2 = verifyApiCheckV2Input(string(getApiCheckV2Body))
)

func TestGetApiCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/api/489", func(w http.ResponseWriter, r *http.Request) {
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

	if !reflect.DeepEqual(resp.Test.Device, inputGetApiCheckV2.Test.Device) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Device, inputGetApiCheckV2.Test.Device)
	}

	if !reflect.DeepEqual(resp.Test.Device.Viewportheight, inputGetApiCheckV2.Test.Device.Viewportheight) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Device.Viewportheight, inputGetApiCheckV2.Test.Device.Viewportheight)
	}

	if !reflect.DeepEqual(resp.Test.Requests, inputGetApiCheckV2.Test.Requests) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Requests, inputGetApiCheckV2.Test.Requests)
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
