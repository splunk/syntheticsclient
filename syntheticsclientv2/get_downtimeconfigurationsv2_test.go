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
	"net/http"
	"reflect"
	"testing"
)

var (
	getDowntimeConfigurationV2Body   = `{"downtimeConfiguration":{"id":329,"name":"dc test","description":"My super awesome test downtimeConfiguration","rule":"pause_tests","startTime":"2024-05-16T20:23:00.000Z","endTime":"2024-05-16T20:38:00.000Z","status":"scheduled","createdAt":"2024-05-15T20:24:07.541Z","updatedAt":"2024-05-15T20:25:44.211Z","testsUpdatedAt":"2024-05-15T20:24:07.541Z","testIds":[29976]}}`
	inputGetDowntimeConfigurationV2  = verifyDowntimeConfigurationV2Input(string(getDowntimeConfigurationV2Body))
	getDowntimeConfigurationsV2Body  = `{"downtimeConfigurations":[{"id":329,"name":"dc test","description":"My super awesome test downtimeConfiguration","rule":"pause_tests","startTime":"2024-05-16T20:23:00.000Z","endTime":"2024-05-16T20:38:00.000Z","status":"scheduled","createdAt":"2024-05-15T20:24:07.541Z","updatedAt":"2024-05-15T20:25:44.211Z","testsUpdatedAt":"2024-05-15T20:24:07.541Z","testCount":1}],"page":1,"pageLimit":1,"totalCount":1}`
	inputGetDowntimeConfigurationsV2 = verifyDowntimeConfigurationsV2Input(string(getDowntimeConfigurationsV2Body))
)

func TestGetDowntimeConfigurationV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getDowntimeConfigurationV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetDowntimeConfigurationV2(1)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.DowntimeConfiguration.ID, inputGetDowntimeConfigurationV2.DowntimeConfiguration.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.ID, inputGetDowntimeConfigurationV2.DowntimeConfiguration.ID)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Name, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Name, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Name)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Description, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Description, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Description)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Rule, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Rule) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Rule, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Rule)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Starttime, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Starttime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Starttime, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Starttime)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Endtime, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Endtime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Endtime, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Endtime)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Status, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Status) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Status, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Status)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Createdat, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Createdat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Createdat, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Createdat)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Updatedat, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Updatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Updatedat, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Updatedat)
	}

	if !reflect.DeepEqual(resp.DowntimeConfiguration.Testsupdatedat, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Testsupdatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DowntimeConfiguration.Testsupdatedat, inputGetDowntimeConfigurationV2.DowntimeConfiguration.Testsupdatedat)
	}

}

func verifyDowntimeConfigurationV2Input(stringInput string) *DowntimeConfigurationV2Response {
	check := &DowntimeConfigurationV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestGetDowntimeConfigurationsV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getDowntimeConfigurationsV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetDowntimeConfigurationsV2()

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].ID, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].ID, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].ID)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Name, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Name, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Name)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Description, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Description, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Description)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Rule, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Rule) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Rule, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Rule)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Starttime, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Starttime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Starttime, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Starttime)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Endtime, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Endtime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Endtime, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Endtime)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Status, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Status) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Status, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Status)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Createdat, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Createdat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Createdat, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Createdat)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Updatedat, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Updatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Updatedat, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Updatedat)
	}

	if !reflect.DeepEqual(resp.Downtimeconfigurations[0].Testsupdatedat, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Testsupdatedat) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Downtimeconfigurations[0].Testsupdatedat, inputGetDowntimeConfigurationsV2.Downtimeconfigurations[0].Testsupdatedat)
	}

}

func verifyDowntimeConfigurationsV2Input(stringInput string) *DowntimeConfigurationsV2Response {
	check := &DowntimeConfigurationsV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}
