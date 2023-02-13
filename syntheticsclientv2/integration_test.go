//go:build integration
// +build integration

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
	"os"
	"testing"
)

var (
	token                    = os.Getenv("API_ACCESS_TOKEN")
	realm                    = os.Getenv("REALM")
	getChecksV2Body          = `{"testType":"","page":1,"perPage":50,"search":"","orderBy":"id"}`
	inputGetChecksV2         = GetChecksV2Options{}
	createVariableV2Body     = `{"variable":{"description":"My super awesome test variable","name":"foodz","secret":false,"value":"bar"}}`
	inputVariableV2Data      = VariableV2Input{}
	createHttpCheckV2Body    = `{"test":{"name":"morebeeps-test","type":"http","url":"https://www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true,"requestMethod":"GET","body":null,"headers":[{"name":"boop","value":"beep"}]}}`
	inputHttpCheckV2Data     = HttpCheckV2Input{}
	createBrowserCheckV2Body = `{"test":{"name":"browser-beep-test","transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://www.splunk.com","action":"go_to_url","wait_for_nav":true},{"name":"Nexter step","type":"click_element","selectorType":"id","wait_for_nav":false,"selector":"free-splunk-click-desktop"}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":5,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}]}}}`
	inputBrowserCheckV2Data  = BrowserCheckV2Input{}
	createPortCheckV2Body    = `{"test":{"name":"splunk - port 443","type":"port","url":"","port":443,"protocol":"tcp","host":"www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true}}`
	inputPortCheckV2Data     = PortCheckV2Input{}
	createApiV2Body          = `{"test":{"active":true,"deviceId":1,"frequency":5,"locationIds":["aws-us-east-1"],"name":"boop-test","schedulingStrategy":"round_robin","requests":[{"configuration":{"name":"Get-Test","requestMethod": "GET","url":"https://api.us1.signalfx.com/v2/synthetics/tests/api/489","headers":{"X-SF-TOKEN":"jinglebellsbatmanshells", "beep":"boop"},"body":null},"setup":[{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"$.requests","variable":"custom-varz"}],"validations":[{"name":"Assert response code equals 200","type":"assert_numeric","actual":"{{response.code}}","expected":"200","comparator":"equals"}]}]}}`
	inputApiCheckV2Data      = ApiCheckV2Input{}
)

// You will need to fill in values for the get and delete tests
// as the check ids will vary from organization to organization

func TestLiveGetChecksV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	err := json.Unmarshal([]byte(getChecksV2Body), &inputGetChecksV2)
	if err != nil {
		t.Fatal(err)
	}

	// Make the request with your check settings and print result
	res, _, err := c.GetChecksV2(&inputGetChecksV2)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreateVariableV2(t *testing.T) {

	err := json.Unmarshal([]byte(createVariableV2Body), &inputVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(inputVariableV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateVariableV2(&inputVariableV2Data)
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

func TestLiveGetVariableV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, _, err := c.GetVariableV2(246)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteVariableV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, err := c.DeleteVariableV2(397)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreateHttpCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createHttpCheckV2Body), &inputHttpCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

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

func TestLiveGetHttpCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, _, err := c.GetHttpCheckV2(1528)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteHttpCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, err := c.DeleteHttpCheckV2(2099)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreateBrowserCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createBrowserCheckV2Body), &inputBrowserCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(inputBrowserCheckV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateBrowserCheckV2(&inputBrowserCheckV2Data)
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

func TestLiveGetBrowserCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, _, err := c.GetBrowserCheckV2(495)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteBrowserCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, err := c.DeleteBrowserCheckV2(2101)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreateApiCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createApiV2Body), &inputApiCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateApiCheckV2(&inputApiCheckV2Data)
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

func TestLiveGetApiCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, _, err := c.GetApiCheckV2(489)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteApiCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, err := c.DeleteApiCheckV2(1093)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreatePortCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createPortCheckV2Body), &inputPortCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

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

func TestLiveGetPortCheckV2(t *testing.T) {

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

func TestLiveDeletePortCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, err := c.DeletePortCheckV2(1649)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}