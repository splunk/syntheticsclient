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
	"os"
	"reflect"
	"testing"
)

var (
	getPortCheckV2Body  = `{"test":{"id":1647,"name":"splunk - port 443","active":true,"frequency":10,"scheduling_strategy":"round_robin","created_at":"2022-11-21T15:38:54.546Z","updated_at":"2022-11-21T15:38:54.554Z","location_ids":["aws-us-east-1"],"type":"port","protocol":"tcp","host":"www.splunk.com","port":443}}`
	inputGetPortCheckV2 = verifyPortCheckV2Input(string(getPortCheckV2Body))
)

func TestGetPortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/port/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(getPortCheckV2Body))
	})

	resp, _, err := testClient.GetPortCheckV2(1)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.Test.ID, inputGetPortCheckV2.Test.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ID, inputGetPortCheckV2.Test.ID)
	}

	if !reflect.DeepEqual(resp.Test.Name, inputGetPortCheckV2.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputGetPortCheckV2.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Type, inputGetPortCheckV2.Test.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Type, inputGetPortCheckV2.Test.Type)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputGetPortCheckV2.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputGetPortCheckV2.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputGetPortCheckV2.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputGetPortCheckV2.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.CreatedAt, inputGetPortCheckV2.Test.CreatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.CreatedAt, inputGetPortCheckV2.Test.CreatedAt)
	}

	if !reflect.DeepEqual(resp.Test.UpdatedAt, inputGetPortCheckV2.Test.UpdatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.UpdatedAt, inputGetPortCheckV2.Test.UpdatedAt)
	}

	if !reflect.DeepEqual(resp.Test.Protocol, inputGetPortCheckV2.Test.Protocol) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Protocol, inputGetPortCheckV2.Test.Protocol)
	}

	if !reflect.DeepEqual(resp.Test.Host, inputGetPortCheckV2.Test.Host) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Host, inputGetPortCheckV2.Test.Host)
	}

	if !reflect.DeepEqual(resp.Test.Port, inputGetPortCheckV2.Test.Port) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, inputGetPortCheckV2.Test.Port)
	}

}

func verifyPortCheckV2Input(stringInput string) *PortCheckV2Response {
	check := &PortCheckV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestLiveGetPortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, _, err := c.GetPortCheckV2(1647)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
