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
	createHttpCheckV2Body = `{"test":{"name":"morebeeps-test","type":"http","url":"https://www.splunk.com","location_ids":["aws-us-east-1"],"frequency":10,"scheduling_strategy":"round_robin","active":true,"request_method":"GET","body":null,"headers":[{"name":"boop","value":"beep"}]}}`
	inputHttpCheckV2Data  = HttpCheckV2Input{}
)

func TestCreateHttpCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(createHttpCheckV2Body))
	})

	json.Unmarshal([]byte(createHttpCheckV2Body), &inputHttpCheckV2Data)

	resp, _, err := testClient.CreateHttpCheckV2(&inputHttpCheckV2Data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputHttpCheckV2Data.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputHttpCheckV2Data.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputHttpCheckV2Data.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputHttpCheckV2Data.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.LocationIds, inputHttpCheckV2Data.Test.LocationIds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.LocationIds, inputHttpCheckV2Data.Test.LocationIds)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputHttpCheckV2Data.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputHttpCheckV2Data.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.RequestMethod, inputHttpCheckV2Data.Test.RequestMethod) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.RequestMethod, inputHttpCheckV2Data.Test.RequestMethod)
	}

	if !reflect.DeepEqual(resp.Test.Body, inputHttpCheckV2Data.Test.Body) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Body, inputHttpCheckV2Data.Test.Body)
	}

	if !reflect.DeepEqual(resp.Test.SchedulingStrategy, inputHttpCheckV2Data.Test.SchedulingStrategy) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.SchedulingStrategy, inputHttpCheckV2Data.Test.SchedulingStrategy)
	}

}

func TestLiveCreateHttpCheckV2(t *testing.T) {
	setup()
	defer teardown()

	json.Unmarshal([]byte(createHttpCheckV2Body), &inputHttpCheckV2Data)

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
	fmt.Println(inputHttpCheckV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateHttpCheckV2(&inputHttpCheckV2Data)
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
