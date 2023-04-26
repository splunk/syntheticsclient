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
	createVariableV2Body     = `{"variable":{"description":"beep-var","name":"foodz","secret":false,"value":"bar"}}`
	inputVariableV2Data      = VariableV2Input{}
	updateVariableV2Body     = `{"variable":{"description":"My super awesome test variable22","name":"foodz","secret":false,"value":"bar"}}`
	updateVariableV2Data      = VariableV2Input{}
	createLocationV2Body     = `{"location":{"id":"private-data-center-go-test","label":"Data Center place", "default":false}}`
	inputLocationV2Data      = LocationV2Input{}
	createHttpCheckV2Body    = `{"test":{"name":"morebeeps-test","type":"http","url":"https://www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true,"requestMethod":"GET","body":null,"headers":[{"name":"boop","value":"beep"}]}}`
	inputHttpCheckV2Data     = HttpCheckV2Input{}
	updateHttpCheckV2Body    = `{"test":{"name":"morebeeps-test","type":"http","url":"https://www.splunk.com/index/","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true,"requestMethod":"GET","body":null,"headers":[{"name":"boop","value":"beep"}]}}`
	updateHttpCheckV2Data     = HttpCheckV2Input{}
	createBrowserCheckV2Body = `{"test":{"name":"browser-beep-test","transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url","options":{"url":"https://splunk.com"},"waitForNav":true},{"name":"click","type":"click_element","selectorType":"id","selector":"clicky","waitForNav":true},{"name":"fill in fieldz","type":"enter_value","selectorType":"id","selector":"beep","value":"{{env.beep-var}}","waitForNav":false},{"name":"accept---Alert","type":"accept_alert"},{"name":"Select-Val-Index","type":"select_option","selectorType":"id","selector":"selectionz","optionSelectorType":"index","optionSelector":"{{env.beep-var}}","waitForNav":false},{"name":"Select-val-text","type":"select_option","selectorType":"id","selector":"textzz","optionSelectorType":"text","optionSelector":"sdad","waitForNav":false},{"name":"Select-Val-Val","type":"select_option","selectorType":"id","selector":"valz","optionSelectorType":"value","optionSelector":"{{env.beep-var}}","waitForNav":false},{"name":"Run JS","type":"run_javascript","value":"beeeeeeep","waitForNav":true},{"name":"Save as text","type":"store_variable_from_element","selectorType":"link","selector":"beepval","variableName":"{{env.terraform-test-foo-301}}"},{"name":"Save JS return Val","type":"store_variable_from_javascript","value":"sdasds","variableName":"{{env.terraform-test-foo-301}}","waitForNav":true}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":5,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}]}}}`
	inputBrowserCheckV2Data  = BrowserCheckV2Input{}
	updateBrowserCheckV2Body = `{"test":{"name":"browser-beep-test","transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url","options":{"url":"https://splunk.com"},"waitForNav":true},{"name":"click","type":"click_element","selectorType":"id","selector":"clicky","waitForNav":true},{"name":"fill in fieldz","type":"enter_value","selectorType":"id","selector":"beep","value":"{{env.beep-var}}","waitForNav":false},{"name":"accept---Alert","type":"accept_alert"},{"name":"Select-Val-Index","type":"select_option","selectorType":"id","selector":"selectionz","optionSelectorType":"index","optionSelector":"{{env.beep-var}}","waitForNav":false},{"name":"Select-val-text","type":"select_option","selectorType":"id","selector":"textzz","optionSelectorType":"text","optionSelector":"sdad","waitForNav":false},{"name":"Select-Val-Val","type":"select_option","selectorType":"id","selector":"valz","optionSelectorType":"value","optionSelector":"{{env.beep-var}}","waitForNav":false},{"name":"Run JS","type":"run_javascript","value":"beeeeeeep","waitForNav":true},{"name":"Save as text","type":"store_variable_from_element","selectorType":"link","selector":"beepval","variableName":"{{env.terraform-test-foo-301}}"},{"name":"Save JS return Val","type":"store_variable_from_javascript","value":"sdasds","variableName":"{{env.terraform-test-foo-301}}","waitForNav":true}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":15,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"cookies":[{"key":"super","value":"dooper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}]}}}`
	updateBrowserCheckV2Data  = BrowserCheckV2Input{}
	createPortCheckV2Body    = `{"test":{"name":"splunk - port 443","type":"port","url":"","port":443,"protocol":"tcp","host":"www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true}}`
	inputPortCheckV2Data     = PortCheckV2Input{}
	updatePortCheckV2Body    = `{"test":{"name":"splunk - port 448","type":"port","url":"","port":448,"protocol":"tcp","host":"www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true}}`
	updatePortCheckV2Data     = PortCheckV2Input{}
	createApiV2Body          = `{"test":{"active":true,"deviceId":1,"frequency":5,"locationIds":["aws-us-east-1"],"name":"boop-test","schedulingStrategy":"round_robin","requests":[{"configuration":{"name":"Get-Test","requestMethod": "GET","url":"https://api.us1.signalfx.com/v2/synthetics/tests/api/489","headers":{"X-SF-TOKEN":"jinglebellsbatmanshells", "beep":"boop"},"body":null},"setup":[{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"$.requests","variable":"custom-varz"}],"validations":[{"name":"Assert response code equals 200","type":"assert_numeric","actual":"{{response.code}}","expected":"200","comparator":"equals"}]}]}}`
	inputApiCheckV2Data      = ApiCheckV2Input{}
	updateApiCheckV2Body          = `{"test":{"active":true,"deviceId":1,"frequency":5,"locationIds":["aws-us-east-1"],"name":"boop-test","schedulingStrategy":"round_robin","requests":[{"configuration":{"name":"Get-Test","requestMethod": "GET","url":"https://api.us1.signalfx.com/v2/synthetics/tests/api/4892","headers":{"X-SF-TOKEN":"jinglebellsbatmanshells", "beep":"boop"},"body":null},"setup":[{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"$.requests","variable":"custom-varz"}],"validations":[{"name":"Assert response code equals 200","type":"assert_numeric","actual":"{{response.code}}","expected":"200","comparator":"equals"}]}]}}`
	updateApiCheckV2Data      = ApiCheckV2Input{}
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

func TestLiveUpdateVariableV2(t *testing.T) {

	err := json.Unmarshal([]byte(updateVariableV2Body), &updateVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Update your client with the token
	c := NewClient(token, realm)

	fmt.Println(updateVariableV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateVariableV2(859, &updateVariableV2Data)
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
	res, _, err := c.GetVariableV2(859)
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
	res, err := c.DeleteVariableV2(398)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveGetDevicesV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, _, err := c.GetDevicesV2()
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

func TestLiveUpdateHttpCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(updateHttpCheckV2Body), &inputHttpCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Update your client with the token
	c := NewClient(token, realm)

	fmt.Println(inputHttpCheckV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateHttpCheckV2(3420, &inputHttpCheckV2Data)
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
	res, _, err := c.GetBrowserCheckV2(5111)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveUpdateBrowserCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(updateBrowserCheckV2Body), &updateBrowserCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Update your client with the token
	c := NewClient(token, realm)

	fmt.Println(updateBrowserCheckV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateBrowserCheckV2(5111, &updateBrowserCheckV2Data)
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

func TestLiveDeleteBrowserCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, err := c.DeleteBrowserCheckV2(3434)
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
	res, _, err := c.GetApiCheckV2(3435)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveUpdateApiCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(updateApiCheckV2Body), &updateApiCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Update your client with the token
	c := NewClient(token, realm)

	fmt.Println(updateApiCheckV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateApiCheckV2(3435, &updateApiCheckV2Data)
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

func TestLiveDeleteApiCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, err := c.DeleteApiCheckV2(3435)
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

func TestLiveUpdatePortCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(updatePortCheckV2Body), &updatePortCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Update your client with the token
	c := NewClient(token, realm)

	fmt.Println(updatePortCheckV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdatePortCheckV2(3436, &updatePortCheckV2Data)
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

func TestLiveCreateLocationV2(t *testing.T) {

	err := json.Unmarshal([]byte(createLocationV2Body), &inputLocationV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(inputLocationV2Data)

	// Make the request with your location settings and print result
	res, reqDetail, err := c.CreateLocationV2(&inputLocationV2Data)
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

func TestLiveGetLocationV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your location settings and print result
	res, _, err := c.GetLocationV2("aws-eu-central-1")
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteLocationV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your location settings and print result
	res, err := c.DeleteLocationV2("private-data-center-go-test")
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
