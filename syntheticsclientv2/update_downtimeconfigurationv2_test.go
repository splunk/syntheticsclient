//go:build unit_tests
// +build unit_tests

// Copyright 2024 Splunk, Inc.
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
	updateDowntimeConfigurationV2Body  = `{"downtimeConfiguration":{"name":"dc test","description":"My super awesome test downtimeConfiguration","rule":"pause_tests","testIds":[29976],"startTime":"2024-05-16T20:23:00.000Z","endTime":"2024-05-16T20:38:00.000Z"}}`
	inputDowntimeConfigurationV2Update = DowntimeConfigurationV2Input{}
)

func TestUpdateDowntimeConfigurationV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateDowntimeConfigurationV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateDowntimeConfigurationV2Body), &inputDowntimeConfigurationV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateDowntimeConfigurationV2(10, &inputDowntimeConfigurationV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Name, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Name, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Name)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Description, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Description, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Description)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Rule, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Rule) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Rule, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Rule)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Testids, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Testids) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Testids, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Testids)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Starttime, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Starttime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Starttime, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Starttime)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Endtime, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Endtime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Endtime, inputDowntimeConfigurationV2Update.DowntimeConfiguration.Endtime)
	}
}
