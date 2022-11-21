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
	createPortCheckV2Body = `{"test":{"name":"splunk - port 443","type":"port","url":"","port":443,"protocol":"tcp","host":"www.splunk.com","location_ids":["aws-us-east-1"],"frequency":10,"scheduling_strategy":"round_robin","active":true}}`
	inputPortCheckV2Data  = PortCheckV2Input{}
)

func TestCreatePortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/port", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(createPortCheckV2Body))
	})

	json.Unmarshal([]byte(createPortCheckV2Body), &inputPortCheckV2Data)

	resp, _, err := testClient.CreatePortCheckV2(&inputPortCheckV2Data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

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

func TestLiveCreatePortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	json.Unmarshal([]byte(createPortCheckV2Body), &inputPortCheckV2Data)

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
	fmt.Println(inputPortCheckV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreatePortCheckV2(&inputPortCheckV2Data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reqDetail)
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
