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
	updatePortCheckV2Body  = `{"test":{"name":"splunk - port 443","type":"port","url":"","port":80,"protocol":"tcp","host":"www.splunk.com","location_ids":["aws-us-east-1"],"frequency":10,"scheduling_strategy":"round_robin","active":true}}`
	inputPortCheckV2Update = PortCheckV2Input{}
)

func TestUpdatePortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/port/1650", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.Write([]byte(updatePortCheckV2Body))
	})

	json.Unmarshal([]byte(updatePortCheckV2Body), &inputPortCheckV2Update)

	resp, _, err := testClient.UpdatePortCheckV2(1650, &inputPortCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputPortCheckV2Update.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputPortCheckV2Update.Test.Name)
	}

}
