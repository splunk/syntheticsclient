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
	"fmt"
	"net/http"
	"os"
	"testing"
)

var (
	deletePortCheckV2RespBody = ``
)

func TestDeletePortCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/browser/19", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Write([]byte(deletePortCheckV2RespBody))
	})

	resp, err := testClient.DeletePortCheckV2(19)
	if err != nil {
		fmt.Println(resp)
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestLiveDeletePortCheckV2(t *testing.T) {
	setup()
	defer teardown()

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
