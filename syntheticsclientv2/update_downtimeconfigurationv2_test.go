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

	if !reflect.DeepEqual(resp.Name, inputDowntimeConfigurationV2Update.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Name, inputDowntimeConfigurationV2Update.Name)
	}

	if !reflect.DeepEqual(resp.Description, inputDowntimeConfigurationV2Update.Description) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Description, inputDowntimeConfigurationV2Update.Description)
	}

	if !reflect.DeepEqual(resp.Rule, inputDowntimeConfigurationV2Update.Rule) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Rule, inputDowntimeConfigurationV2Update.Rule)
	}

	if !reflect.DeepEqual(resp.Testids, inputDowntimeConfigurationV2Update.Testids) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Testids, inputDowntimeConfigurationV2Update.Testids)
	}

	if !reflect.DeepEqual(resp.Starttime, inputDowntimeConfigurationV2Update.Starttime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Starttime, inputDowntimeConfigurationV2Update.Starttime)
	}

	if !reflect.DeepEqual(resp.Endtime, inputDowntimeConfigurationV2Update.Endtime) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Endtime, inputDowntimeConfigurationV2Update.Endtime)
	}
}

func TestUpdateDowntimeConfigurationV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.UpdateDowntimeConfigurationV2(10, &DowntimeConfigurationV2Input{})
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

func TestUpdateDowntimeConfigurationV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})
	_, _, err := unreachableClient.UpdateDowntimeConfigurationV2(10, &DowntimeConfigurationV2Input{})
	if err == nil {
		t.Fatal("expected a connection error")
	}
}

func TestUpdateDowntimeConfigurationV2ReturnsResponseOnEmptyBody(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
	})

	resp, details, err := testClient.UpdateDowntimeConfigurationV2(10, &DowntimeConfigurationV2Input{})
	if err != nil {
		t.Fatalf("expected no error for empty response, but got: %v", err)
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	if resp == nil {
		t.Fatal("expected non-nil response for empty body")
	}
}
