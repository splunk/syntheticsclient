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
	getDevicesV2Body  = `{"devices":[{"id":1,"label":"Desktop","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Splunk Synthetics) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36","viewport_width":1366,"viewport_height":768,"network_connection":{"description":"Standard Cable","upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":null}},{"id":2,"label":"iPad","user_agent":"Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X; Splunk Synthetics) AppleWebKit/604.1.25 (KHTML, like Gecko) Version/11.0 Mobile/15A5304j Safari/604.1","viewport_width":1024,"viewport_height":768,"network_connection":{"description":"Legacy Cable","upload_bandwidth":1000,"download_bandwidth":5000,"latency":28,"packet_loss":null}},{"id":3,"label":"iPhone","user_agent":"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X; Splunk Synthetics) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A356 Safari/604.1","viewport_width":375,"viewport_height":844,"network_connection":{"description":"Mobile LTE","upload_bandwidth":12000,"download_bandwidth":12000,"latency":70,"packet_loss":null}},{"id":4,"label":"Samsung Galaxy","user_agent":"Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U Build/R16NW; Splunk Synthetics) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Mobile Safari/537.36","viewport_width":360,"viewport_height":640,"network_connection":{"description":"Mobile 3G","upload_bandwidth":768,"download_bandwidth":1600,"latency":150,"packet_loss":null}}],"default_device_id":null}`
	inputGetDevicesV2 = verifyDevicesV2Input(getDevicesV2Body)
)

func TestGetDevicesV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getDevicesV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetDevicesV2()

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Devices[1].ID, inputGetDevicesV2.Devices[1].ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Devices[1].ID, inputGetDevicesV2.Devices[1].ID)
	}

	if !reflect.DeepEqual(resp.Devices[1].Label, inputGetDevicesV2.Devices[1].Label) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Devices[1].Label, inputGetDevicesV2.Devices[1].Label)
	}

	if !reflect.DeepEqual(resp.Devices[1].UserAgent, inputGetDevicesV2.Devices[1].UserAgent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Devices[1].UserAgent, inputGetDevicesV2.Devices[1].UserAgent)
	}

	if !reflect.DeepEqual(resp.Devices[1].ViewportWidth, inputGetDevicesV2.Devices[1].ViewportWidth) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Devices[1].ViewportWidth, inputGetDevicesV2.Devices[1].ViewportWidth)
	}

	if !reflect.DeepEqual(resp.Devices[1].ViewportHeight, inputGetDevicesV2.Devices[1].ViewportHeight) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Devices[1].ViewportHeight, inputGetDevicesV2.Devices[1].ViewportHeight)
	}

	if !reflect.DeepEqual(resp.Devices[1].Networkconnection, inputGetDevicesV2.Devices[1].Networkconnection) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Devices[1].Networkconnection, inputGetDevicesV2.Devices[1].Networkconnection)
	}

}

func verifyDevicesV2Input(stringInput string) *DevicesV2Response {
	check := &DevicesV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}
