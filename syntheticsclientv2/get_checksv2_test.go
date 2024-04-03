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
	getChecksV2Body   = `{"testType":"","page":1,"perPage":50,"search":"","orderBy":"id"}`
	inputGetChecksV2  = verifyChecksV2Input(string(getChecksV2Body))
	getChecksV2Output = `{"tests":[{"id":482,"name":"Test of Splunk.com","active":true, "automaticRetries": 1, "frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-15T16:05:25.815Z","updated_at":"2022-09-29T19:13:13.853Z","location_ids":["aws-us-east-1"],"type":"browser"},{"id":489,"name":"Appinspect login API","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-16T15:47:43.730Z","updated_at":"2022-08-16T15:47:43.741Z","location_ids":["aws-us-east-1"],"type":"api"},{"id":490,"name":"Arch Linux Packages","active":true,"frequency":10,"scheduling_strategy":"round_robin","created_at":"2022-08-16T16:48:42.119Z","updated_at":"2022-08-16T16:48:42.131Z","location_ids":["aws-us-east-1"],"type":"http"},{"id":492,"name":"Test of Splunkbase","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-16T19:35:54.014Z","updated_at":"2022-09-29T19:13:13.907Z","location_ids":["aws-us-east-1"],"type":"browser"},{"id":493,"name":"Brewery API","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-16T19:44:15.626Z","updated_at":"2022-08-16T19:44:15.635Z","location_ids":["aws-us-east-1"],"type":"api"},{"id":495,"name":"Multi-step test of legacy Splunkbase","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-17T01:24:44.579Z","updated_at":"2022-09-29T19:13:13.203Z","location_ids":["aws-us-east-1"],"type":"browser"},{"id":496,"name":"Multi-step Test of new Splunkbase","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-08-17T01:33:27.771Z","updated_at":"2022-09-29T19:13:13.997Z","location_ids":["aws-us-east-1"],"type":"browser"},{"id":935,"name":"This test does test stuff","active":true,"frequency":30,"scheduling_strategy":"round_robin","created_at":"2022-10-26T14:48:36.026Z","updated_at":"2022-10-26T14:48:36.037Z","location_ids":["aws-us-east-1"],"type":"api"},{"id":1116,"name":"boop-test","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-11-16T19:18:59.603Z","updated_at":"2022-11-16T19:20:58.911Z","location_ids":["aws-us-east-1","aws-ap-northeast-1"],"type":"api"},{"id":1128,"name":"boopbeep","active":true,"frequency":5,"scheduling_strategy":"round_robin","created_at":"2022-11-17T14:19:49.564Z","updated_at":"2022-11-17T14:19:49.571Z","location_ids":["aws-us-east-1"],"type":"browser"}],"page":1,"per_page":50,"next_page_link":null,"total_count":10}`
	output            = &ChecksV2Response{}
)

func TestGetChecksV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getChecksV2Output))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(getChecksV2Output), output)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.GetChecksV2(inputGetChecksV2)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Tests, output.Tests) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Tests, output.Tests)
	}

}

func verifyChecksV2Input(stringInput string) *GetChecksV2Options {
	check := &GetChecksV2Options{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}
