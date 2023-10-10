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
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var (
	createPortCheckV2Body = `{"test":{"customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "name":"splunk - port 443","type":"port","url":"","port":443,"protocol":"tcp","host":"www.splunk.com","location_ids":["aws-us-east-1"],"frequency":10,"scheduling_strategy":"round_robin","active":true}}`
	inputPortCheckV2Data  = PortCheckV2Input{}
)

func TestCreatePortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/port", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(createPortCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createPortCheckV2Body), &inputPortCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreatePortCheckV2(&inputPortCheckV2Data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputPortCheckV2Data.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputPortCheckV2Data.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Type, inputPortCheckV2Data.Test.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Type, inputPortCheckV2Data.Test.Type)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputPortCheckV2Data.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputPortCheckV2Data.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputPortCheckV2Data.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputPortCheckV2Data.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.Protocol, inputPortCheckV2Data.Test.Protocol) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Protocol, inputPortCheckV2Data.Test.Protocol)
	}

	if !reflect.DeepEqual(resp.Test.Host, inputPortCheckV2Data.Test.Host) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Host, inputPortCheckV2Data.Test.Host)
	}

	if !reflect.DeepEqual(resp.Test.Port, inputPortCheckV2Data.Test.Port) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, inputPortCheckV2Data.Test.Port)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputPortCheckV2Data.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputPortCheckV2Data.Test.Customproperties)
	}
}
