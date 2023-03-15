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
	updateLocationV2Body  = `{"location":{"id":"private-data-center","label":"Data Center"}}`
	inputLocationV2Update = LocationV2Input{}
)

func TestUpdateLocationV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateLocationV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateLocationV2Body), &inputLocationV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateLocationV2("private-data-center", &inputLocationV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Location.ID, inputLocationV2Data.Location.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Location.ID, inputLocationV2Data.Location.ID)
	}

	if !reflect.DeepEqual(resp.Location.Label, inputLocationV2Data.Location.Label) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Location.Label, inputLocationV2Data.Location.Label)
	}

}
