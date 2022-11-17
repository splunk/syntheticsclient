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
	"net/http"
	"reflect"
	"testing"
	"fmt"
	"encoding/json"
	"os"
)

var (
	createApiV2Body  = `{"test":{"active":true,"device_id":1,"frequency":5,"location_ids":["aws-us-east-1"],"name":"boop-test","scheduling_strategy":"round_robin","requests":[{"configuration":{"name":"Get-Test","requestMethod":"GET","url":"https://api.us1.signalfx.com/v2/synthetics/tests/api/489","headers":{"X-SF-TOKEN":"jinglebellsbatmanshells", "beep":"boop"},"body":null},"setup":[{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"$.requests","variable":"custom-varz"}],"validations":[{"name":"Assert response code equals 200","type":"assert_numeric","actual":"{{response.code}}","expected":"200","comparator":"equals"}]}]}}`
	inputData = ApiCheckV2Input{}
)

func TestCreateApiCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/api", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(createApiV2Body))
	})

	json.Unmarshal([]byte(createApiV2Body), &inputData)

	resp, _, err := testClient.CreateApiCheckV2(&inputData)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputData.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputData.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Active, inputData.Test.Active) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Active, inputData.Test.Active)
	}

	if !reflect.DeepEqual(resp.Test.Locationids, inputData.Test.Locationids) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Locationids, inputData.Test.Locationids)
	}

	if !reflect.DeepEqual(resp.Test.Frequency, inputData.Test.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Frequency, inputData.Test.Frequency)
	}

	if !reflect.DeepEqual(resp.Test.Requests, inputData.Test.Requests) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Requests, inputData.Test.Requests)
	}

	if !reflect.DeepEqual(resp.Test.Schedulingstrategy, inputData.Test.Schedulingstrategy) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Schedulingstrategy, inputData.Test.Schedulingstrategy)
	}

}


func TestLiveCreateApiCheckV2(t *testing.T) {
	setup()
	defer teardown()

	json.Unmarshal([]byte(createApiV2Body), &inputData)

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)
	
	fmt.Println(c)
	fmt.Println(inputData)

	// Make the request with your check settings and print result
  res, reqDetail, err := c.CreateApiCheckV2(&inputData)
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