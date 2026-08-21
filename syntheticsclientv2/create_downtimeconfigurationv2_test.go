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
	createDowntimeConfigurationV2Body = `{"downtimeConfiguration":{"name":"dc test","description":"My super awesome test downtimeConfiguration","rule":"augment_data","testIds":[482],"startTime":"2024-05-16T20:23:00.000Z","endTime":"2024-05-16T20:38:00.000Z","timezone":"UTC","recurrence":{"repeats":{"type":"daily"},"end":{"type":"after","value":"10"}}}}`
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

	if !reflect.DeepEqual(resp.ID, inputDowntimeConfigurationV2Data.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ID, inputDowntimeConfigurationV2Data.ID)
	}

	if !reflect.DeepEqual(resp.Name, inputDowntimeConfigurationV2Data.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Name, inputDowntimeConfigurationV2Data.Name)
	}

	if !reflect.DeepEqual(resp.Description, inputDowntimeConfigurationV2Data.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Description, inputDowntimeConfigurationV2Data.Description)
	}

	if !reflect.DeepEqual(resp.Rule, inputDowntimeConfigurationV2Data.Rule) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Rule, inputDowntimeConfigurationV2Data.Rule)
	}

	if !reflect.DeepEqual(resp.Starttime, inputDowntimeConfigurationV2Data.Starttime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Starttime, inputDowntimeConfigurationV2Data.Starttime)
	}

	if !reflect.DeepEqual(resp.Endtime, inputDowntimeConfigurationV2Data.Endtime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Endtime, inputDowntimeConfigurationV2Data.Endtime)
	}

	if !reflect.DeepEqual(resp.Recurrence, inputDowntimeConfigurationV2Data.Recurrence) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Recurrence, inputDowntimeConfigurationV2Data.Recurrence)
	}

	if !reflect.DeepEqual(resp.Timezone, inputDowntimeConfigurationV2Data.Timezone) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Timezone, inputDowntimeConfigurationV2Data.Timezone)
	}

}

func TestCreateDowntimeConfigurationV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.CreateDowntimeConfigurationV2(&DowntimeConfigurationV2Input{})
	if err == nil {
		t.Fatal("expected a parse error, but got none")
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	if resp != nil {
		t.Errorf("expected nil response, got %#v", resp)
	}
}

func TestCreateDowntimeConfigurationV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})
	_, _, err := unreachableClient.CreateDowntimeConfigurationV2(&DowntimeConfigurationV2Input{})
	if err == nil {
		t.Fatal("expected a connection error")
	}
}
