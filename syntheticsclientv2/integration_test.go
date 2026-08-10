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
	"log"
	"os"
	"testing"
	"time"
)

var (
	token                                = os.Getenv("API_ACCESS_TOKEN")
	realm                                = os.Getenv("REALM")
	liveGetChecksV2Body                  = `{"testType":"","page":1,"perPage":50,"search":"","orderBy":"id"}`
	liveInputGetChecksV2                 = GetChecksV2Options{}
	liveCreateVariableV2Body             = `{"variable":{"description":"beep-var","name":"a-variable-named-foodz","secret":false,"value":"bar"}}`
	liveInputVariableV2Data              = VariableV2Input{}
	liveUpdateVariableV2Body             = `{"variable":{"description":"My super awesome test variable22","name":"a-variable-named-foodz","secret":false,"value":"bar"}}`
	updateVariableV2Data                 = VariableV2Input{}
	liveCreateLocationV2Body             = `{"location":{"id":"private-data-center-go-test","label":"Data Center place", "default":false}}`
	liveInputLocationV2Data              = LocationV2Input{}
	liveCreateHttpCheckV2Body            = `{"test":{"name":"a-minimal-http-integration-test","type":"http", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"url":"https://www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true,"requestMethod":"GET","body":null,"userAgent":null,"authentication":null,"verifyCertificates":false}}`
	liveInputHttpCheckV2Data             = HttpCheckV2Input{}
	liveUpdateHttpCheckV2Body            = `{"test":{"name":"a-maximal-http-integration-test-update","type":"http", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"url":"https://www.splunk.com/updated","locationIds":["aws-us-east-1","aws-ap-southeast-2","aws-ap-southeast-4"],"frequency":30,"schedulingStrategy":"round_robin","active":true,"requestMethod":"GET","body":null,"headers":[{"name":"header-1","value":"value-1"},{"name":"header_2","value":"value_2"}],"validations":[{"type":"assert_string","actual":"{{response.first_byte_time}}","expected":"100","comparator":"equals"},{"type":"assert_string","actual":"{{headers.Content-Length}}","expected":"100","comparator":"does_not_equal"}],"userAgent":"user-agent_standards met","authentication":{"username":"beepusers","password":"{{env.terraform-test-foo-301}}"},"verifyCertificates":true}}`
	updateHttpCheckV2Data                = HttpCheckV2Input{}
	createMaximalBrowserCheckV2Body      = `{"test":{"name":"a-maximal-browser-beep-test", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url","options":{"url":"https://splunk.com"}},{"name":"click","type":"click_element","selectors":[{"type":"id","value":"clicky"}],"waitForNav":true,"waitForNavTimeout":2000},{"name":"fill in fieldz","type":"enter_value","selectors":[{"type":"id","value":"beep"}],"value":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"accept---Alert","type":"accept_alert"},{"name":"Select-Val-Index","type":"select_option","selectors":[{"type":"id","value":"selectionz"}],"optionSelectorType":"index","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-val-text","type":"select_option","selectors":[{"type":"id","value":"textzz"}],"optionSelectorType":"text","optionSelector":"sdad","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-Val-Val","type":"select_option","selectors":[{"type":"id","value":"valz"}],"optionSelectorType":"value","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Run JS","type":"run_javascript","value":"beeeeeeep","waitForNav":true,"waitForNavTimeout":2000},{"name":"Save as text","type":"store_variable_from_element","selectors":[{"type":"link","value":"beepval"}],"variableName":"{{env.terraform-test-foo-301}}"},{"name":"Wait","type":"wait","duration":1312},{"name":"Save JS return Val","type":"store_variable_from_javascript","value":"sdasds","variableName":"{{env.terraform-test-foo-301}}","waitForNav":true,"waitForNavTimeout":2000}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":5,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}]}}}`
	createMinimalBrowserCheckV2Body      = `{"test":{"name":"a-minimal-browser-beep-test", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url"}]}],"locationIds":["aws-us-east-1"],"deviceId":1,"frequency":5,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true}}}`
	liveInputBrowserCheckV2Data          = BrowserCheckV2Input{}
	liveUpdateBrowserCheckV2Body         = `{"test":{"name":"a-browser-beep-test", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://splunk.com","action":"go_to_url","options":{"url":"https://splunk.com"}},{"name":"click","type":"click_element","selectors":[{"type":"id","value":"clicky"}],"waitForNav":true,"waitForNavTimeout":2000},{"name":"fill in fieldz","type":"enter_value","selectors":[{"type":"id","value":"beep"}],"value":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"accept---Alert","type":"accept_alert"},{"name":"Select-Val-Index","type":"select_option","selectors":[{"type":"id","value":"selectionz"}],"optionSelectorType":"index","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-val-text","type":"select_option","selectors":[{"type":"id","value":"textzz"}],"optionSelectorType":"text","optionSelector":"sdad","waitForNav":false,"waitForNavTimeout":50},{"name":"Select-Val-Val","type":"select_option","selectors":[{"type":"id","value":"valz"}],"optionSelectorType":"value","optionSelector":"{{env.beep-var}}","waitForNav":false,"waitForNavTimeout":50},{"name":"Run JS","type":"run_javascript","value":"beeeeeeep","waitForNav":true,"waitForNavTimeout":2000},{"name":"Save as text","type":"store_variable_from_element","selectors":[{"type":"link","value":"beepval"}],"variableName":"{{env.terraform-test-foo-301}}"},{"name":"Save JS return Val","type":"store_variable_from_javascript","value":"sdasds","variableName":"{{env.terraform-test-foo-301}}","waitForNav":true,"waitForNavTimeout":2000}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","locationIds":["aws-us-east-1"],"deviceId":1,"frequency":15,"schedulingStrategy":"round_robin","active":true,"advancedSettings":{"verifyCertificates":true,"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"cookies":[{"key":"super","value":"dooper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}]}}}`
	updateBrowserCheckV2Data             = BrowserCheckV2Input{}
	liveCreatePortCheckV2Body            = `{"test":{"name":"a - port 443 check","type":"port", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"url":"","port":443,"protocol":"tcp","host":"www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true}}`
	liveInputPortCheckV2Data             = PortCheckV2Input{}
	liveUpdatePortCheckV2Body            = `{"test":{"name":"a2 - port 443 check","type":"port", "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"url":"","port":448,"protocol":"tcp","host":"www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true}}`
	updatePortCheckV2Data                = PortCheckV2Input{}
	liveCreateApiV2Body                  = `{"test":{"active":true,"deviceId":1,"frequency":5, "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"locationIds":["aws-us-east-1"],"name":"a-maximual-API-boop-test","schedulingStrategy":"round_robin","requests":[{"configuration":{"name":"Get-Test","requestMethod":"GET","url":"https://api.us1.signalfx.com/v2/synthetics/v2/tests/api/489","headers":{"X-SF-TOKEN":"jinglebellsbatmanshells","beep":"boop"},"body":null},"setup":[{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"sd","variable":"extractsetupvar"},{"name":"JavaScript run","type":"javascript","code":"asdasd","variable":"jsvarsetup"},{"name":"Save response body","type":"save","value":"{{response.body}}","variable":"savesetupvar"}],"validations":[{"name":"JavaScript run","type":"javascript","code":"codetorun","variable":"jscodevar"},{"name":"Save response body","type":"save","value":"{{response.body}}","variable":"saverespvar"},{"name":"Assert response code equals 200","type":"assert_numeric","actual":"{{response.code}}","expected":"200","comparator":"equals"},{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"js.extractor","variable":"extractjvar"}]}]}}`
	createMinimalApiV2Body               = `{"test":{"active":true,"deviceId":1,"frequency":5, "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"locationIds":["aws-us-east-1"],"name":"a-minimal-API-boop-test","schedulingStrategy":"round_robin","requests":[{"configuration":{"name":"apishortGet-Test","requestMethod":"GET","url":"https://api.us1.signalfx.com/v2/synthetics/v2/tests/api/489"}}]}}`
	inputApiCheckV2Data                  = ApiCheckV2Input{}
	liveUpdateApiCheckV2Body             = `{"test":{"active":true,"deviceId":1,"frequency":5, "automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}],"locationIds":["aws-us-east-1"],"name":"a-API-boop-test","schedulingStrategy":"round_robin","requests":[{"configuration":{"name":"Get-Test","requestMethod": "GET","url":"https://api.us1.signalfx.com/v2/synthetics/v2/tests/api/4892","headers":{"X-SF-TOKEN":"jinglebellsbatmanshells", "beep":"boop"},"body":null},"setup":[{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"$.requests","variable":"custom-varz"}],"validations":[{"name":"Assert response code equals 200","type":"assert_numeric","actual":"{{response.code}}","expected":"200","comparator":"equals"}]}]}}`
	updateApiCheckV2Data                 = ApiCheckV2Input{}
	liveInputDowntimeConfigurationV2Data = DowntimeConfigurationV2Input{}
	updateDowntimeConfigurationV2Data    = DowntimeConfigurationV2Input{}
)

// Every test below creates whatever fixture it needs and deletes it before
// returning, so the suite can be re-run repeatedly against the same org
// without accumulating orphaned resources or colliding on fixed names/ids.

const (
	livePrereqBeepVarName      = "beep-var"
	livePrereqTerraformVarName = "terraform-test-foo-301"
)

// TestMain provisions the two variables referenced by name
// (via {{env.beep-var}} and {{env.terraform-test-foo-301}}) in the http/browser
// check bodies used below, since the Synthetics API requires an
// authentication.password variable reference to resolve to an existing
// variable. Both are torn down after the suite runs.
func TestMain(m *testing.M) {
	c := NewClient(token, realm)

	// A prior run killed before m.Run() returned (e.g. a cancelled CI job) skips
	// both t.Cleanup and the deletes below, leaving these fixed-name variables
	// behind. Reclaim any leftovers before creating fresh ones so that run isn't
	// stuck failing every subsequent run with a duplicate-name error.
	reclaimPrerequisiteVariable(c, livePrereqBeepVarName)
	reclaimPrerequisiteVariable(c, livePrereqTerraformVarName)

	beepVar, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: livePrereqBeepVarName, Value: "bar", Secret: false}})
	if err != nil {
		log.Fatalf("failed to create prerequisite variable %q: %v", livePrereqBeepVarName, err)
	}

	terraformVar, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: livePrereqTerraformVarName, Value: "bar", Secret: false}})
	if err != nil {
		if _, delErr := c.DeleteVariableV2(beepVar.Variable.ID); delErr != nil {
			log.Printf("failed to clean up prerequisite variable %q: %v", livePrereqBeepVarName, delErr)
		}
		log.Fatalf("failed to create prerequisite variable %q: %v", livePrereqTerraformVarName, err)
	}

	code := m.Run()

	if _, err := c.DeleteVariableV2(beepVar.Variable.ID); err != nil {
		log.Printf("failed to clean up prerequisite variable %q: %v", livePrereqBeepVarName, err)
	}
	if _, err := c.DeleteVariableV2(terraformVar.Variable.ID); err != nil {
		log.Printf("failed to clean up prerequisite variable %q: %v", livePrereqTerraformVarName, err)
	}

	os.Exit(code)
}

// reclaimPrerequisiteVariable deletes any existing variable with the given name so a
// stale leftover from an interrupted prior run doesn't cause CreateVariableV2 below to
// fail with a duplicate-name error.
func reclaimPrerequisiteVariable(c *Client, name string) {
	existing, _, err := c.GetVariablesV2()
	if err != nil {
		log.Printf("failed to list variables while checking for leftover %q: %v", name, err)
		return
	}
	for _, v := range existing.Variable {
		if v.Name != name {
			continue
		}
		if _, err := c.DeleteVariableV2(v.ID); err != nil {
			log.Printf("failed to reclaim leftover prerequisite variable %q: %v", name, err)
		}
	}
}

func TestLiveGetChecksV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	err := json.Unmarshal([]byte(liveGetChecksV2Body), &liveInputGetChecksV2)
	if err != nil {
		t.Fatal(err)
	}

	// Make the request with your check settings and print result
	res, _, err := c.GetChecksV2(&liveInputGetChecksV2)
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

	err := json.Unmarshal([]byte(liveCreateVariableV2Body), &liveInputVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(liveInputVariableV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateVariableV2(&liveInputVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	t.Cleanup(func() {
		if _, err := c.DeleteVariableV2(res.Variable.ID); err != nil {
			t.Errorf("failed to clean up variable %d: %v", res.Variable.ID, err)
		}
	})

}

func TestLiveUpdateVariableV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: "a-variable-named-foodz", Value: "bar", Secret: false}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteVariableV2(created.Variable.ID); err != nil {
			t.Errorf("failed to clean up variable %d: %v", created.Variable.ID, err)
		}
	})

	err = json.Unmarshal([]byte(liveUpdateVariableV2Body), &updateVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(updateVariableV2Data)

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateVariableV2(created.Variable.ID, &updateVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

}

func TestLiveGetVariableV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: "a-variable-named-foodz", Value: "bar", Secret: false}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteVariableV2(created.Variable.ID); err != nil {
			t.Errorf("failed to clean up variable %d: %v", created.Variable.ID, err)
		}
	})

	// Make the request with your check settings and print result
	res, _, err := c.GetVariableV2(created.Variable.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

}

func TestLiveDeleteVariableV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: "a-variable-named-foodz", Value: "bar", Secret: false}})
	if err != nil {
		t.Fatal(err)
	}

	// Make the request with your check settings and print result
	res, err := c.DeleteVariableV2(created.Variable.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

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

func TestLiveHttpCheckCreateUpdateDeleteV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)
	var err error

	checkId, err := CreateHttpCheckV2(liveCreateHttpCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteHttpCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up http check %d: %v", checkId, err)
		}
	})

	err = UpdateHttpCheckV2(checkId, liveUpdateHttpCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = GetHttpCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

	err2 := UpdateHttpCheckV2(checkId, liveCreateHttpCheckV2Body, c)
	if err2 != nil {
		t.Fatal(err2)
	}

	err = GetHttpCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func CreateHttpCheckV2(test string, c *Client) (int, error) {

	err := json.Unmarshal([]byte(test), &liveInputHttpCheckV2Data)
	if err != nil {
		return 0, err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateHttpCheckV2(&liveInputHttpCheckV2Data)
	if err != nil {
		return 0, err
	}
	fmt.Printf("\nReq was: \n%v\n", reqDetail)
	JsonPrint(res)
	liveInputHttpCheckV2Data = HttpCheckV2Input{}

	return res.Test.ID, nil
}

func UpdateHttpCheckV2(checkId int, test string, c *Client) error {

	err := json.Unmarshal([]byte(test), &liveInputHttpCheckV2Data)
	if err != nil {
		return err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateHttpCheckV2(checkId, &liveInputHttpCheckV2Data)
	if err != nil {
		return err
	}
	fmt.Printf("\nReq was: \n%v\n", reqDetail)
	JsonPrint(res)
	liveInputHttpCheckV2Data = HttpCheckV2Input{}

	return nil
}

func GetHttpCheckV2(checkId int, c *Client) error {

	// Make the request with your check settings and print result
	res, _, err := c.GetHttpCheckV2(checkId)
	if err != nil {
		return err
	}

	JsonPrint(res)

	return nil
}

func DeleteHttpCheckV2(checkId int, c *Client) error {
	// Make the request with your check settings and print result
	res, err := c.DeleteHttpCheckV2(checkId)
	if err != nil {
		return err
	}

	JsonPrint(res)

	return nil

}

func TestLiveCreateHttpCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateHttpCheckV2(liveCreateHttpCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Http Check ID: %d", checkId)

	t.Cleanup(func() {
		if err := DeleteHttpCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up http check %d: %v", checkId, err)
		}
	})

}

func TestLiveUpdateHttpCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateHttpCheckV2(liveCreateHttpCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteHttpCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up http check %d: %v", checkId, err)
		}
	})

	err = UpdateHttpCheckV2(checkId, liveUpdateHttpCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveGetHttpCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateHttpCheckV2(liveCreateHttpCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteHttpCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up http check %d: %v", checkId, err)
		}
	})

	err = GetHttpCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteHttpCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateHttpCheckV2(liveCreateHttpCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = DeleteHttpCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveBrowserCheckCreateUpdateAndDeleteV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)
	var err error

	checkId, err := CreateBrowserCheckV2(createMaximalBrowserCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteBrowserCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up browser check %d: %v", checkId, err)
		}
	})

	err = UpdateBrowserCheckV2(checkId, createMinimalBrowserCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = GetBrowserCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

	err2 := UpdateBrowserCheckV2(checkId, createMaximalBrowserCheckV2Body, c)
	if err2 != nil {
		t.Fatal(err2)
	}

	err = GetBrowserCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func CreateBrowserCheckV2(test string, c *Client) (int, error) {

	err := json.Unmarshal([]byte(test), &liveInputBrowserCheckV2Data)
	if err != nil {
		return 0, err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateBrowserCheckV2(&liveInputBrowserCheckV2Data)
	if err != nil {
		return 0, err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return res.Test.ID, nil
}

func UpdateBrowserCheckV2(checkId int, test string, c *Client) error {

	err := json.Unmarshal([]byte(test), &updateBrowserCheckV2Data)
	if err != nil {
		return err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateBrowserCheckV2(checkId, &updateBrowserCheckV2Data)
	if err != nil {
		return err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return nil
}

func GetBrowserCheckV2(checkId int, c *Client) error {
	// Make the request with your check settings and print result
	res, _, err := c.GetBrowserCheckV2(checkId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func DeleteBrowserCheckV2(checkId int, c *Client) error {
	// Make the request with your check settings and print result
	res, err := c.DeleteBrowserCheckV2(checkId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func TestLiveCreateBrowserCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateBrowserCheckV2(createMaximalBrowserCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Browser Check ID: %d", checkId)

	t.Cleanup(func() {
		if err := DeleteBrowserCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up browser check %d: %v", checkId, err)
		}
	})
}

func TestLiveGetBrowserCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateBrowserCheckV2(createMaximalBrowserCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteBrowserCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up browser check %d: %v", checkId, err)
		}
	})

	err = GetBrowserCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveUpdateBrowserCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateBrowserCheckV2(createMaximalBrowserCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteBrowserCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up browser check %d: %v", checkId, err)
		}
	})

	err = UpdateBrowserCheckV2(checkId, createMinimalBrowserCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteBrowserCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateBrowserCheckV2(createMaximalBrowserCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = DeleteBrowserCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func CreateApiCheckV2(test string, c *Client) (int, error) {

	err := json.Unmarshal([]byte(test), &inputApiCheckV2Data)
	if err != nil {
		return 0, err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateApiCheckV2(&inputApiCheckV2Data)
	if err != nil {
		return 0, err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return res.Test.ID, nil
}

func UpdateApiCheckV2(checkId int, test string, c *Client) error {

	err := json.Unmarshal([]byte(test), &updateApiCheckV2Data)
	if err != nil {
		return err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateApiCheckV2(checkId, &updateApiCheckV2Data)
	if err != nil {
		return err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return nil
}

func GetApiCheckV2(checkId int, c *Client) error {
	// Make the request with your check settings and print result
	res, _, err := c.GetApiCheckV2(checkId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func DeleteApiCheckV2(checkId int, c *Client) error {
	// Make the request with your check settings and print result
	res, err := c.DeleteApiCheckV2(checkId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func TestLiveApiCheckCreateUpdateAndDeleteV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)
	var err error

	checkId, err := CreateApiCheckV2(liveCreateApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteApiCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up api check %d: %v", checkId, err)
		}
	})

	err = UpdateApiCheckV2(checkId, createMinimalApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = GetApiCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

	err2 := UpdateApiCheckV2(checkId, liveCreateApiV2Body, c)
	if err2 != nil {
		t.Fatal(err2)
	}

	err = GetApiCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreateApiCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateApiCheckV2(liveCreateApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Api Check ID: %d", checkId)

	t.Cleanup(func() {
		if err := DeleteApiCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up api check %d: %v", checkId, err)
		}
	})
}

func TestLiveGetApiCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateApiCheckV2(liveCreateApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteApiCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up api check %d: %v", checkId, err)
		}
	})

	err = GetApiCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveUpdateApiCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateApiCheckV2(liveCreateApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteApiCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up api check %d: %v", checkId, err)
		}
	})

	err = UpdateApiCheckV2(checkId, liveUpdateApiCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeleteApiCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateApiCheckV2(liveCreateApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = DeleteApiCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func CreatePortCheckV2(test string, c *Client) (int, error) {

	err := json.Unmarshal([]byte(test), &liveInputPortCheckV2Data)
	if err != nil {
		return 0, err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreatePortCheckV2(&liveInputPortCheckV2Data)
	if err != nil {
		return 0, err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return res.Test.ID, nil
}

func UpdatePortCheckV2(checkId int, test string, c *Client) error {

	err := json.Unmarshal([]byte(test), &updatePortCheckV2Data)
	if err != nil {
		return err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdatePortCheckV2(checkId, &updatePortCheckV2Data)
	if err != nil {
		return err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return nil
}

func GetPortCheckV2(checkId int, c *Client) error {
	// Make the request with your check settings and print result
	res, _, err := c.GetPortCheckV2(checkId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func DeletePortCheckV2(checkId int, c *Client) error {
	// Make the request with your check settings and print result
	res, err := c.DeletePortCheckV2(checkId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func TestLivePortCheckCreateUpdateAndDeleteV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)
	var err error

	checkId, err := CreatePortCheckV2(liveCreatePortCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeletePortCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up port check %d: %v", checkId, err)
		}
	})

	err = UpdatePortCheckV2(checkId, liveUpdatePortCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = GetPortCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

	err2 := UpdatePortCheckV2(checkId, liveCreatePortCheckV2Body, c)
	if err2 != nil {
		t.Fatal(err2)
	}

	err = GetPortCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreatePortCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreatePortCheckV2(liveCreatePortCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Port Check ID: %d", checkId)

	t.Cleanup(func() {
		if err := DeletePortCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up port check %d: %v", checkId, err)
		}
	})
}

func TestLiveGetPortCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreatePortCheckV2(liveCreatePortCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeletePortCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up port check %d: %v", checkId, err)
		}
	})

	err = GetPortCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveUpdatePortCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreatePortCheckV2(liveCreatePortCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeletePortCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up port check %d: %v", checkId, err)
		}
	})

	err = UpdatePortCheckV2(checkId, liveUpdatePortCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveDeletePortCheckV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreatePortCheckV2(liveCreatePortCheckV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = DeletePortCheckV2(checkId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveCreateLocationV2(t *testing.T) {

	err := json.Unmarshal([]byte(liveCreateLocationV2Body), &liveInputLocationV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(liveInputLocationV2Data)

	// Make the request with your location settings and print result
	res, reqDetail, err := c.CreateLocationV2(&liveInputLocationV2Data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	t.Cleanup(func() {
		if _, err := c.DeleteLocationV2(res.Location.ID); err != nil {
			t.Errorf("failed to clean up location %q: %v", res.Location.ID, err)
		}
	})

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

	created, _, err := c.CreateLocationV2(&LocationV2Input{Location: Location{ID: "private-data-center-go-test-delete", Label: "Data Center place", Default: false}})
	if err != nil {
		t.Fatal(err)
	}

	// Make the request with your location settings and print result
	res, err := c.DeleteLocationV2(created.Location.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

}

func TestLiveDowntimeConfigurationCreateUpdateAndDeleteV2(t *testing.T) {

	//Create your client with the token
	c := NewClient(token, realm)
	var err error

	checkId, err := CreateApiCheckV2(liveCreateApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteApiCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up api check %d: %v", checkId, err)
		}
	})

	//There are restrictions on startTime and endTime for a downtime_configuration so we set the startTime to 10 days
	//in the future and the endTime to be 1 hour after the startTime
	tenDaysFromNow := time.Now().AddDate(0, 0, 10)
	year, month, day := tenDaysFromNow.Date()
	startTime := fmt.Sprintf("%d-%02d-%02dT20:00:00.000Z", year, int(month), day)
	endTime := fmt.Sprintf("%d-%02d-%02dT21:00:00.000Z", year, int(month), day)

	createDowntimeConfigurationV2Body := fmt.Sprintf("{\"downtimeConfiguration\":{\"name\":\"dc test\",\"description\":\"My super awesome test downtimeConfiguration\",\"rule\":\"augment_data\",\"testIds\":[%d],\"startTime\":\"%s\",\"endTime\":\"%s\"}}", checkId, startTime, endTime)

	downtimeConfigId, err := CreateDowntimeConfigurationV2(createDowntimeConfigurationV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteDowntimeConfigurationV2(downtimeConfigId, c); err != nil {
			t.Errorf("failed to clean up downtime configuration %d: %v", downtimeConfigId, err)
		}
	})

	err = GetDowntimeConfigurationV2(downtimeConfigId, c)
	if err != nil {
		t.Fatal(err)
	}

	updateDowntimeConfigurationV2Body := fmt.Sprintf("{\"downtimeConfiguration\":{\"name\":\"dc test\",\"description\":\"My super awesome test downtimeConfiguration\",\"rule\":\"pause_tests\",\"testIds\":[%d],\"startTime\":\"%s\",\"endTime\":\"%s\"}}", checkId, startTime, endTime)

	err = UpdateDowntimeConfigurationV2(downtimeConfigId, updateDowntimeConfigurationV2Body, c)
	if err != nil {
		t.Fatal(err)
	}

	err = GetDowntimeConfigurationV2(downtimeConfigId, c)
	if err != nil {
		t.Fatal(err)
	}

}

func CreateDowntimeConfigurationV2(downtimeConfiguration string, c *Client) (int, error) {

	err := json.Unmarshal([]byte(downtimeConfiguration), &liveInputDowntimeConfigurationV2Data)
	if err != nil {
		return 0, err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.CreateDowntimeConfigurationV2(&liveInputDowntimeConfigurationV2Data)
	if err != nil {
		return 0, err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return res.DowntimeConfiguration.ID, nil
}

func UpdateDowntimeConfigurationV2(downtimeConfigId int, downtimeConfiguration string, c *Client) error {

	err := json.Unmarshal([]byte(downtimeConfiguration), &updateDowntimeConfigurationV2Data)
	if err != nil {
		return err
	}

	// Make the request with your check settings and print result
	res, reqDetail, err := c.UpdateDowntimeConfigurationV2(downtimeConfigId, &updateDowntimeConfigurationV2Data)
	if err != nil {
		return err
	}
	fmt.Println(reqDetail)
	JsonPrint(res)

	return nil
}

func GetDowntimeConfigurationV2(downtimeConfigId int, c *Client) error {
	// Make the request with your check settings and print result
	res, _, err := c.GetDowntimeConfigurationV2(downtimeConfigId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func DeleteDowntimeConfigurationV2(downtimeConfigId int, c *Client) error {
	// Make the request with your check settings and print result
	res, err := c.DeleteDowntimeConfigurationV2(downtimeConfigId)
	if err != nil {
		return err
	}
	JsonPrint(res)

	return nil
}

func liveCreateSslCheckV2Input(name string) *SslCheckV2Input {
	input := &SslCheckV2Input{}
	input.Test.Name = name
	input.Test.LocationIds = []string{"aws-us-east-1"}
	input.Test.Frequency = 10
	input.Test.SchedulingStrategy = "round_robin"
	input.Test.Active = true
	input.Test.Automaticretries = 1
	input.Test.Host = "www.splunk.com"
	input.Test.Port = 443
	input.Test.AllowSelfSigned = false
	input.Test.AllowUntrustedRoot = false
	return input
}

func TestLiveSslCheckCreateUpdateAndDeleteV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateSslCheckV2(liveCreateSslCheckV2Input("a-maximal-ssl-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteSslCheckV2(created.Test.ID); err != nil {
			t.Errorf("failed to clean up ssl check %d: %v", created.Test.ID, err)
		}
	})
	JsonPrint(created)

	newFrequency := 30
	newActive := false
	update := &SslCheckV2UpdateInput{}
	update.Test.Frequency = &newFrequency
	update.Test.Active = &newActive

	updated, _, err := c.UpdateSslCheckV2(created.Test.ID, update)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(updated)

	res, _, err := c.GetSslCheckV2(created.Test.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)
}

func TestLiveCreateSslCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	res, _, err := c.CreateSslCheckV2(liveCreateSslCheckV2Input("a-minimal-ssl-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteSslCheckV2(res.Test.ID); err != nil {
			t.Errorf("failed to clean up ssl check %d: %v", res.Test.ID, err)
		}
	})
	JsonPrint(res)
}

func TestLiveGetSslCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateSslCheckV2(liveCreateSslCheckV2Input("a-gettable-ssl-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteSslCheckV2(created.Test.ID); err != nil {
			t.Errorf("failed to clean up ssl check %d: %v", created.Test.ID, err)
		}
	})

	res, _, err := c.GetSslCheckV2(created.Test.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)
}

func TestLiveUpdateSslCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateSslCheckV2(liveCreateSslCheckV2Input("an-updatable-ssl-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteSslCheckV2(created.Test.ID); err != nil {
			t.Errorf("failed to clean up ssl check %d: %v", created.Test.ID, err)
		}
	})

	newFrequency := 30
	newActive := false
	update := &SslCheckV2UpdateInput{}
	update.Test.Frequency = &newFrequency
	update.Test.Active = &newActive

	res, _, err := c.UpdateSslCheckV2(created.Test.ID, update)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

	confirmed, _, err := c.GetSslCheckV2(created.Test.ID)
	if err != nil {
		t.Fatal(err)
	}
	if confirmed.Test.Frequency != newFrequency {
		t.Errorf("frequency = %d, want %d", confirmed.Test.Frequency, newFrequency)
	}
}

func TestLiveDeleteSslCheckV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateSslCheckV2(liveCreateSslCheckV2Input("a-deletable-ssl-integration-test"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.DeleteSslCheckV2(created.Test.ID); err != nil {
		t.Fatal(err)
	}
}

func liveCreateTotpVariableV2Input(name string) *TotpVariableV2Input {
	return &TotpVariableV2Input{Totp: TotpVariableInput{
		Name:       name,
		Secret:     "JBSWY3DPEHPK3PXP",
		Digits:     6,
		Interval:   30,
		HmacDigest: "sha1",
	}}
}

func TestLiveTotpVariableCreateUpdateAndDeleteV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateTotpVariableV2(liveCreateTotpVariableV2Input("a-maximal-totp-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteTotpVariableV2(created.Totp.ID); err != nil {
			t.Errorf("failed to clean up totp variable %d: %v", created.Totp.ID, err)
		}
	})
	JsonPrint(created)

	newDigits := 8
	newInterval := 60
	updated, _, err := c.UpdateTotpVariableV2(created.Totp.ID, &TotpVariableV2UpdateInput{Totp: TotpVariableUpdateInput{
		Digits:   &newDigits,
		Interval: &newInterval,
	}})
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(updated)

	res, _, err := c.GetTotpVariableV2(created.Totp.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

	list, _, err := c.GetTotpVariablesV2()
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, totp := range list.Totps {
		if totp.ID == created.Totp.ID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("created totp variable %d not present in GetTotpVariablesV2 response", created.Totp.ID)
	}
}

func TestLiveCreateTotpVariableV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	res, _, err := c.CreateTotpVariableV2(liveCreateTotpVariableV2Input("a-minimal-totp-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteTotpVariableV2(res.Totp.ID); err != nil {
			t.Errorf("failed to clean up totp variable %d: %v", res.Totp.ID, err)
		}
	})
	JsonPrint(res)
}

func TestLiveGetTotpVariableV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateTotpVariableV2(liveCreateTotpVariableV2Input("a-gettable-totp-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteTotpVariableV2(created.Totp.ID); err != nil {
			t.Errorf("failed to clean up totp variable %d: %v", created.Totp.ID, err)
		}
	})

	res, _, err := c.GetTotpVariableV2(created.Totp.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)
}

func TestLiveUpdateTotpVariableV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateTotpVariableV2(liveCreateTotpVariableV2Input("an-updatable-totp-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteTotpVariableV2(created.Totp.ID); err != nil {
			t.Errorf("failed to clean up totp variable %d: %v", created.Totp.ID, err)
		}
	})

	newDigits := 8
	newInterval := 60
	res, _, err := c.UpdateTotpVariableV2(created.Totp.ID, &TotpVariableV2UpdateInput{Totp: TotpVariableUpdateInput{
		Digits:   &newDigits,
		Interval: &newInterval,
	}})
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

	confirmed, _, err := c.GetTotpVariableV2(created.Totp.ID)
	if err != nil {
		t.Fatal(err)
	}
	if confirmed.Totp.Digits != newDigits {
		t.Errorf("digits = %d, want %d", confirmed.Totp.Digits, newDigits)
	}
	if confirmed.Totp.Interval != newInterval {
		t.Errorf("interval = %d, want %d", confirmed.Totp.Interval, newInterval)
	}
}

func TestLiveDeleteTotpVariableV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateTotpVariableV2(liveCreateTotpVariableV2Input("a-deletable-totp-integration-test"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.DeleteTotpVariableV2(created.Totp.ID); err != nil {
		t.Fatal(err)
	}
}

const liveTestCaCertContent = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJekNDQWd1Z0F3SUJBZ0lVZWtuOHdoNzhCRG9CaHkxZ21xWXBPVjhnWGJjd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0lURWZNQjBHQTFVRUF3d1djM2x1ZEdndGFXNTBaV2R5WVhScGIyNHRkR1Z6ZERBZUZ3MHlOakE0TVRBeApOVFU0TXpsYUZ3MHpOakE0TURjeE5UVTRNemxhTUNFeEh6QWRCZ05WQkFNTUZuTjViblJvTFdsdWRHVm5jbUYwCmFXOXVMWFJsYzNRd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURKODdKRWVNKzUKV0x5UGF1ZnBmSzVJbWZJMFJhSndybnozZnEzd3c3d2VnSlR4bkt2WkJZNUJaOHhyVjVWcDl6WS8vUE5TM29GZwpLbUJUQnVQZEVMVHU0VkZKZDlTc0lYajhtUnNoNS81aWJUekxaNk1GNVI3NjIzZkJpQ1lNMWZOVXhRMVQvZTlIClV1cC9xb2ZBZ21nMDJEZzJUWE1RMU1IVi9HUGJvaTBjWVR0TnlZR1hVbWNDYkhCZllkNGN1UmI4a0FWbE1UZW8KTjlmU2FEakxqbzNnazhIK1VITlpaRU1WS0dVczkzUE1MV2JzV2NmTE01TWptU21FakxrT3BMbXkwRXRYWC9MOApGYVordU9YTmtLRU5XcCtlRVJJb256aVlPcXN3bitCdVAxbmpMUktTV3AvK1ZtL0s2V1ZGTDFPOHFlOEN6Y3J2CmpWWXNraFVHR21pbkFnTUJBQUdqVXpCUk1CMEdBMVVkRGdRV0JCVDMrMDhpTkdkajFLTEdrdlY5TVE5bXkvd0gKWURBZkJnTlZIU01FR0RBV2dCVDMrMDhpTkdkajFLTEdrdlY5TVE5bXkvd0hZREFQQmdOVkhSTUJBZjhFQlRBRApBUUgvTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFCOGxLK0hVays0Wk5ZUG5ZWE94QzhBK1VTL2tvU0U5cmlECjlYb29Oc3diQ3pMR2RWTkFKb1pybHAzRnZBZVZXTXN1bkNrdE9YaTlUSEJnN0c0cmJ3b09PaE1nTU11Q3Z4QlQKOHByZmNOYm5xU3lSemNGWTIyNHBuYU85ci9YajNFTkhpZmg0QzBKZ2xVTk5wWjgwdTFUS0ZKaGl3OUlicUFCdQpXM0pEeEY2MGk3R2hmdnFmUmZBWkt3cFNOYTFBaTVvTFBYSVVlaytFeEFaeC8wd3htT0xOdGxhK2RyT0UvWUpYCjNsZGJ5cjBydFZkZEVTcy9tR1k3Y0tMQmhLVkRWTUpwTHkzSi9qTEIxUm5uOHZtRmF6b1JyTjYvUXhNajVGTSsKdDhINDRnaVRwbC90VTdDelpETTdNbHByN01qMU80WFdqSWlRcW5HVXhPNElxL1d0dWZIQwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="

func liveCreateCaCertificateV2Input(name string) *CaCertificateV2Input {
	return &CaCertificateV2Input{CaCert: CaCertificateInput{
		Name:          name,
		Content:       liveTestCaCertContent,
		FileExtension: "pem",
		Filename:      "ca_cert_file.pem",
	}}
}

func TestLiveCaCertificateCreateUpdateAndDeleteV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateCaCertificateV2(liveCreateCaCertificateV2Input("a-maximal-cacert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteCaCertificateV2(created.CaCert.ID); err != nil {
			t.Errorf("failed to clean up ca certificate %d: %v", created.CaCert.ID, err)
		}
	})
	JsonPrint(created)

	newDescription := "updated by integration test"
	updated, _, err := c.UpdateCaCertificateV2(created.CaCert.ID, &CaCertificateV2UpdateInput{CaCert: CaCertificateUpdateInput{
		Description: &newDescription,
	}})
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(updated)

	res, _, err := c.GetCaCertificateV2(created.CaCert.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

	list, _, err := c.GetCaCertificatesV2()
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, cert := range list.CaCerts {
		if cert.ID == created.CaCert.ID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("created ca certificate %d not present in GetCaCertificatesV2 response", created.CaCert.ID)
	}
}

func TestLiveCreateCaCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	res, _, err := c.CreateCaCertificateV2(liveCreateCaCertificateV2Input("a-minimal-cacert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteCaCertificateV2(res.CaCert.ID); err != nil {
			t.Errorf("failed to clean up ca certificate %d: %v", res.CaCert.ID, err)
		}
	})
	JsonPrint(res)
}

func TestLiveGetCaCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateCaCertificateV2(liveCreateCaCertificateV2Input("a-gettable-cacert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteCaCertificateV2(created.CaCert.ID); err != nil {
			t.Errorf("failed to clean up ca certificate %d: %v", created.CaCert.ID, err)
		}
	})

	res, _, err := c.GetCaCertificateV2(created.CaCert.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)
}

func TestLiveUpdateCaCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateCaCertificateV2(liveCreateCaCertificateV2Input("an-updatable-cacert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteCaCertificateV2(created.CaCert.ID); err != nil {
			t.Errorf("failed to clean up ca certificate %d: %v", created.CaCert.ID, err)
		}
	})

	newDescription := "updated by integration test"
	res, _, err := c.UpdateCaCertificateV2(created.CaCert.ID, &CaCertificateV2UpdateInput{CaCert: CaCertificateUpdateInput{
		Description: &newDescription,
	}})
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

	confirmed, _, err := c.GetCaCertificateV2(created.CaCert.ID)
	if err != nil {
		t.Fatal(err)
	}
	if confirmed.CaCert.Description != newDescription {
		t.Errorf("description = %q, want %q", confirmed.CaCert.Description, newDescription)
	}
}

func TestLiveDeleteCaCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateCaCertificateV2(liveCreateCaCertificateV2Input("a-deletable-cacert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.DeleteCaCertificateV2(created.CaCert.ID); err != nil {
		t.Fatal(err)
	}
}

const liveTestClientCertPrivateKeyContent = "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2QUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktZd2dnU2lBZ0VBQW9JQkFRREo4N0pFZU0rNVdMeVAKYXVmcGZLNUltZkkwUmFKd3JuejNmcTN3dzd3ZWdKVHhuS3ZaQlk1Qlo4eHJWNVZwOXpZLy9QTlMzb0ZnS21CVApCdVBkRUxUdTRWRkpkOVNzSVhqOG1Sc2g1LzVpYlR6TFo2TUY1Ujc2MjNmQmlDWU0xZk5VeFExVC9lOUhVdXAvCnFvZkFnbWcwMkRnMlRYTVExTUhWL0dQYm9pMGNZVHROeVlHWFVtY0NiSEJmWWQ0Y3VSYjhrQVZsTVRlb045ZlMKYURqTGpvM2drOEgrVUhOWlpFTVZLR1VzOTNQTUxXYnNXY2ZMTTVNam1TbUVqTGtPcExteTBFdFhYL0w4RmFaKwp1T1hOa0tFTldwK2VFUklvbnppWU9xc3duK0J1UDFuakxSS1NXcC8rVm0vSzZXVkZMMU84cWU4Q3pjcnZqVllzCmtoVUdHbWluQWdNQkFBRUNnZ0VBSXI0SVpJd3VIREkrV2lQbm1yem0xTG1iTjgvazlxS21BQVBzazVkd3hRU1UKMnczN2FGM3l6NkMyUTU4eEpxWXZVSW5KS0cvNzdObk5jV3NsaHpIcEZwRnZwUVoyOFZmZTB3SFo3NWJVSmdXcAo2RW8vZXZPa1JUNjlWdTkvc0VTY1ZIQ0Q3dmVvRXVxYVNmVkIzbVh3M0dwNEhTdHN5Sy81V3NGTlFvc2ZYSnExClZtdGFqMWxkRXR2UmhLanY5ZkdYcnRZSFpVR2c2VzVqV0h0a0F5VmZhSkdSbU9HcTA2UVd4aWg0YW9lVHVhQ3QKSUxwci9lUkdQbHpiSVIxWXlhblhGMzVkdGVRTE9mQWR2YXBCWFJlbzA0VWZSWW1hbnVVYSt3Um5ENnBUTEhIKwpNaFNoQkZhQzR4OERxOGJJZnFWOGR5VElyU3pnUUtQY0t3akM3dGZqQVFLQmdRRG1KQklhYnh0L0pacXEvc2gzCm9MUmg5aERFcEV1ejVtYnAyaVNhcXZEYkZQRm9Ic2JVQ2daTHZUL01seS9QYmY0QlBYdjluRWkyRGphNTFFVHAKeWNMZDNnU2JkVS9pQzgwV1BOQzErZ0hwbDBYS0RTaEtGRlRHT2orZ1VNc2ExMkl6aUFRK0R1Ymh5Y1c2Y0QrbwpOeFc2YUloT0JjZHBjK0t0ZUtWWmxVQXRBUUtCZ1FEZ3BNZDN6cktINXNvemlWaXBtRlRwNEI1em1Pd0JPc3grCjN5c2tBSFNXZUkyVkl5T1djcGoxWlhzMUtwaUlUMU9FVGU3UWVpTUFXcmNoVjRMekxiVXdwRzYvRG1icUU5aEUKTW12bmYrS01wL3F1ZVU4akdNSS9PdXpRdVBNVHFIcml6Y0xkdTZCYStnVVNFV1EzdENaeXM5MnlQdlpYTDN3dgpGMkVMZ3ZRTnB3S0JnSHdPOWJOS01ZaFl2UWR3VUtBc0FSRE5sRHhzVkdLbDBOUSt3M3ljcVRsd0VMSVA1UjVvClNQeUxCOWxCcG9RcXhzSGtZdkpUVE43V3lxbGh3OFJDL3NpYTVlRG5YQ2grTkEvSXVMbGdDNmZmNDc4SFdMQ1cKUlJ5V1NiWWgxMXFnd0U4SEEwSnd4Z1R3djZYQTNJL1JJZVZhZEIrYS9lUGFsRmJ1c2pPWVFRQUJBb0dBZThoUQpZU1AwSEE1L3ZJWWg1TkdiZUlPV1Evd3ZqejNuRU1ISDg3Nk1mNTFONXEvR0hGQnBHRThpNU5qajA3aGlQTFQwCnNzdWFIY2Zld1BDSHA1ZTREMldMNEpyKytseVUvbjhLRmpYUmo4Ky93Z1AySjFDdE9Fb3YwNU1WM2U4b1IzRTUKdnhSejk2MXN2ZGYzY1BwRGRWREhDRURKWEtFOXZIVVZkRkprU0dFQ2dZQmM5ZUNrR2Z0clk3SGk3WktadnFEVwpVd210S2t5aW9zTnFraW4zZUZNRE50UGVMaGh5S3laa1N3WExpcVYyVkZqNG1uYXd4KzFYd2hoanlTTFJualhECmlKdlRQQTFtKzBNRmQrWWhiRjZkVmxQaThTNkMzeDNremlFOGdSZVduVDBNaGx2UG8rVm10Q1E1empRMTFzZUMKWFVoQytiaEdxZjFaRUpua1RsbHUrUT09Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K"

func liveCreateClientCertificateV2Input(name string) *ClientCertificateV2Input {
	return &ClientCertificateV2Input{Certificate: ClientCertificateInput{
		Name:   name,
		Domain: "api.integration-test.example.com",
		PublicKey: ClientCertificateKeyInput{
			Content:       liveTestCaCertContent,
			Filename:      "client.crt",
			FileExtension: "pem",
		},
		PrivateKey: ClientCertificatePrivateKeyInput{
			Content:       liveTestClientCertPrivateKeyContent,
			Filename:      "client.key",
			FileExtension: "pem",
		},
	}}
}

func TestLiveClientCertificateCreateUpdateAndDeleteV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateClientCertificateV2(liveCreateClientCertificateV2Input("a-maximal-clientcert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteClientCertificateV2(created.Certificate.ID); err != nil {
			t.Errorf("failed to clean up client certificate %d: %v", created.Certificate.ID, err)
		}
	})
	JsonPrint(created)

	newDescription := "updated by integration test"
	updated, _, err := c.UpdateClientCertificateV2(created.Certificate.ID, &ClientCertificateV2UpdateInput{Certificate: ClientCertificateUpdateInput{
		Description: &newDescription,
	}})
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(updated)

	res, _, err := c.GetClientCertificateV2(created.Certificate.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

	list, _, err := c.GetClientCertificatesV2()
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, cert := range list.Certificates {
		if cert.ID == created.Certificate.ID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("created client certificate %d not present in GetClientCertificatesV2 response", created.Certificate.ID)
	}
}

func TestLiveCreateClientCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	res, _, err := c.CreateClientCertificateV2(liveCreateClientCertificateV2Input("a-minimal-clientcert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteClientCertificateV2(res.Certificate.ID); err != nil {
			t.Errorf("failed to clean up client certificate %d: %v", res.Certificate.ID, err)
		}
	})
	JsonPrint(res)
}

func TestLiveGetClientCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateClientCertificateV2(liveCreateClientCertificateV2Input("a-gettable-clientcert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteClientCertificateV2(created.Certificate.ID); err != nil {
			t.Errorf("failed to clean up client certificate %d: %v", created.Certificate.ID, err)
		}
	})

	res, _, err := c.GetClientCertificateV2(created.Certificate.ID)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)
}

func TestLiveUpdateClientCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateClientCertificateV2(liveCreateClientCertificateV2Input("an-updatable-clientcert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteClientCertificateV2(created.Certificate.ID); err != nil {
			t.Errorf("failed to clean up client certificate %d: %v", created.Certificate.ID, err)
		}
	})

	newDescription := "updated by integration test"
	res, _, err := c.UpdateClientCertificateV2(created.Certificate.ID, &ClientCertificateV2UpdateInput{Certificate: ClientCertificateUpdateInput{
		Description: &newDescription,
	}})
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)

	confirmed, _, err := c.GetClientCertificateV2(created.Certificate.ID)
	if err != nil {
		t.Fatal(err)
	}
	if confirmed.Certificate.Description != newDescription {
		t.Errorf("description = %q, want %q", confirmed.Certificate.Description, newDescription)
	}
}

func TestLiveDeleteClientCertificateV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateClientCertificateV2(liveCreateClientCertificateV2Input("a-deletable-clientcert-integration-test"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.DeleteClientCertificateV2(created.Certificate.ID); err != nil {
		t.Fatal(err)
	}
}

func TestLiveGetVariablesV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: "a-listable-variable-integration-test", Value: "bar", Secret: false}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteVariableV2(created.Variable.ID); err != nil {
			t.Errorf("failed to clean up variable %d: %v", created.Variable.ID, err)
		}
	})

	res, _, err := c.GetVariablesV2()
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, v := range res.Variable {
		if v.ID == created.Variable.ID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("created variable %d not present in GetVariablesV2 response", created.Variable.ID)
	}
}

func TestLiveGetLocationsV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	res, _, err := c.GetLocationsV2()
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Location) == 0 {
		t.Error("expected at least one location in GetLocationsV2 response, got none")
	}
	JsonPrint(res)
}

func TestLiveGetDowntimeConfigurationsV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	checkId, err := CreateApiCheckV2(createMinimalApiV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteApiCheckV2(checkId, c); err != nil {
			t.Errorf("failed to clean up api check %d: %v", checkId, err)
		}
	})

	tenDaysFromNow := time.Now().AddDate(0, 0, 10)
	year, month, day := tenDaysFromNow.Date()
	startTime := fmt.Sprintf("%d-%02d-%02dT20:00:00.000Z", year, int(month), day)
	endTime := fmt.Sprintf("%d-%02d-%02dT21:00:00.000Z", year, int(month), day)

	createDowntimeConfigurationV2Body := fmt.Sprintf("{\"downtimeConfiguration\":{\"name\":\"dc list test\",\"description\":\"created by integration test\",\"rule\":\"augment_data\",\"testIds\":[%d],\"startTime\":\"%s\",\"endTime\":\"%s\"}}", checkId, startTime, endTime)

	downtimeConfigId, err := CreateDowntimeConfigurationV2(createDowntimeConfigurationV2Body, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := DeleteDowntimeConfigurationV2(downtimeConfigId, c); err != nil {
			t.Errorf("failed to clean up downtime configuration %d: %v", downtimeConfigId, err)
		}
	})

	res, _, err := c.GetDowntimeConfigurationsV2(&GetDowntimeConfigurationsV2Options{})
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, dc := range res.Downtimeconfigurations {
		if dc.ID == downtimeConfigId {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("created downtime configuration %d not present in GetDowntimeConfigurationsV2 response", downtimeConfigId)
	}
}

func TestLiveGetExcludedFileTypesV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	res, _, err := c.GetExcludedFileTypesV2()
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(res)
}

func liveCreateHttpCheckV2WithNullablePortInput(name string) *HttpCheckV2InputWithNullablePort {
	input := &HttpCheckV2InputWithNullablePort{}
	input.Test.Name = name
	input.Test.Type = "http"
	input.Test.URL = "https://www.splunk.com"
	input.Test.LocationIds = []string{"aws-us-east-1"}
	input.Test.Frequency = 10
	input.Test.SchedulingStrategy = "round_robin"
	input.Test.Active = true
	input.Test.RequestMethod = "GET"
	input.Test.Automaticretries = 1
	input.Test.Port = *NewNullableInt(443)
	return input
}

func TestLiveHttpCheckWithNullablePortCreateUpdateAndDeleteV2(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateHttpCheckV2WithNullablePort(liveCreateHttpCheckV2WithNullablePortInput("a-maximal-nullable-port-http-integration-test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteHttpCheckV2(created.Test.ID); err != nil {
			t.Errorf("failed to clean up http check %d: %v", created.Test.ID, err)
		}
	})
	JsonPrint(created)

	update := liveCreateHttpCheckV2WithNullablePortInput("a-maximal-nullable-port-http-integration-test")
	update.Test.Port = *NewNullInt()

	updated, _, err := c.UpdateHttpCheckV2WithNullablePort(created.Test.ID, update)
	if err != nil {
		t.Fatal(err)
	}
	JsonPrint(updated)

	res, _, err := c.GetHttpCheckV2WithNullablePort(created.Test.ID)
	if err != nil {
		t.Fatal(err)
	}
	if res.Test.Port.Value != nil {
		t.Errorf("port = %v, want nil", *res.Test.Port.Value)
	}
	JsonPrint(res)
}

func TestLiveGetHttpCheckV2ReturnsErrorForNonexistentID(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	if _, _, err := c.GetHttpCheckV2(0); err == nil {
		t.Fatal("expected error getting nonexistent http check, got nil")
	}
}

func TestLiveGetVariableV2ReturnsErrorForNonexistentID(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	if _, _, err := c.GetVariableV2(0); err == nil {
		t.Fatal("expected error getting nonexistent variable, got nil")
	}
}

func TestLiveDeleteApiCheckV2ReturnsErrorForNonexistentID(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	if _, err := c.DeleteApiCheckV2(0); err == nil {
		t.Fatal("expected error deleting nonexistent api check, got nil")
	}
}

func TestLiveCreateVariableV2ReturnsErrorOnDuplicateName(t *testing.T) {
	//Create your client with the token
	c := NewClient(token, realm)

	created, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: "a-duplicate-name-integration-test", Value: "bar", Secret: false}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := c.DeleteVariableV2(created.Variable.ID); err != nil {
			t.Errorf("failed to clean up variable %d: %v", created.Variable.ID, err)
		}
	})

	if _, _, err := c.CreateVariableV2(&VariableV2Input{Variable: Variable{Name: "a-duplicate-name-integration-test", Value: "bar", Secret: false}}); err == nil {
		t.Fatal("expected error creating variable with duplicate name, got nil")
	}
}
