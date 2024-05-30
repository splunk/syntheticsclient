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
	createDowntimeConfigurationV2Body = `{"downtimeConfiguration":{"name":"dc test","description":"My super awesome test downtimeConfiguration","rule":"augment_data","testIds":[482],"startTime":"2024-05-16T20:23:00.000Z","endTime":"2024-05-16T20:38:00.000Z"}}`
	inputDowntimeConfigurationV2Data  = DowntimeConfigurationV2Input{}
)

func TestCreateDowntimeConfigurationV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(createDowntimeConfigurationV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createDowntimeConfigurationV2Body), &inputDowntimeConfigurationV2Data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateDowntimeConfigurationV2(&inputDowntimeConfigurationV2Data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.DowntimeConfiguration.ID, inputDowntimeConfigurationV2Data.DowntimeConfiguration.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.ID, inputDowntimeConfigurationV2Data.DowntimeConfiguration.ID)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Name, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Name, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Name)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Description, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Description, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Description)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Rule, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Rule) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Rule, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Rule)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Starttime, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Starttime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Starttime, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Starttime)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Endtime, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Endtime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Endtime, inputDowntimeConfigurationV2Data.DowntimeConfiguration.Endtime)
	}

}
