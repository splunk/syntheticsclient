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
	"fmt"
	"net/http"
	"testing"
)

var (
	deleteVariableV2RespBody = ``
)

func TestDeleteVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/variables/19", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		_, err := w.Write([]byte(deleteVariableV2RespBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, err := testClient.DeleteVariableV2(19)
	if err != nil {
		fmt.Println(resp)
		t.Fatal(err)
	}
	fmt.Println(resp)
}
