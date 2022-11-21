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

func TestLiveGetVariableV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

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

func TestLiveGetPortCheckV2(t *testing.T) {

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

func TestLiveGetHttpCheckV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

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

func TestLiveGetChecksV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	// Make the request with your check settings and print result
	res, _, err := c.GetChecksV2(inputGetChecksV2)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}

func TestLiveGetBrowserCheckV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

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

func TestLiveGetApiCheckV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

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

func TestLiveDeleteVariableV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
	fmt.Println(inputData)

	// Make the request with your check settings and print result
	res, err := c.DeleteVariableV2(254)
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

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
	fmt.Println(inputData)

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

func TestLiveDeleteApiCheckV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
	fmt.Println(inputData)

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

func TestLiveCreateVariableV2(t *testing.T) {

	err := json.Unmarshal([]byte(createVariableV2Body), &inputVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
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

func TestLiveCreateHttpCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createHttpCheckV2Body), &inputHttpCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

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

func TestLiveCreateBrowserCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createBrowserCheckV2Body), &inputBrowserCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
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

func TestLiveCreateApiCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createApiV2Body), &inputData)
	if err != nil {
		t.Fatal(err)
	}

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

func TestLiveCreatePortCheckV2(t *testing.T) {

	err := json.Unmarshal([]byte(createPortCheckV2Body), &inputPortCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

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

func TestLiveDeleteBrowserCheckV2(t *testing.T) {

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
	fmt.Println(inputData)

	// Make the request with your check settings and print result
	res, err := c.DeleteBrowserCheckV2(1115)
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

	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a valid realm (E.G. us0, us1, eu0, etc) environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := NewClient(token, realm)

	fmt.Println(c)
	fmt.Println(inputData)

	// Make the request with your check settings and print result
	res, err := c.DeleteHttpCheckV2(1528)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}

	if err != nil {
		t.Fatal(err)
	}

}
