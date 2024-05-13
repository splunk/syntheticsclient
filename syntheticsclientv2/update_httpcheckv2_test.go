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
  updateHttpCheckV2Body  = `{"test":{"customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "name":"morebeeps-test","type":"http","url":"https://www.splunk.com","location_ids":["aws-us-east-1"],"frequency":10,"scheduling_strategy":"round_robin","active":true,"request_method":"GET","body":null,"port":443,"headers":[{"name":"boop","value":"beep"}]}}`
	inputHttpCheckV2Update = HttpCheckV2Input{}
)

func TestUpdateHttpCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateHttpCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateHttpCheckV2Body), &inputHttpCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateHttpCheckV2(10, &inputHttpCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputHttpCheckV2Update.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputHttpCheckV2Update.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputHttpCheckV2Update.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputHttpCheckV2Update.Test.Customproperties)
	}

	if !reflect.DeepEqual(resp.Test.Port, inputHttpCheckV2Update.Test.Port) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, inputHttpCheckV2Update.Test.Port)
	}
}
