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
	"reflect"
	"testing"
)

var (
	updateHttpCheckV2Body  = `{"test":{"automaticRetries": 1, "customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "name":"morebeeps-test","type":"http","url":"https://www.splunk.com","location_ids":["aws-us-east-1"],"frequency":10,"scheduling_strategy":"round_robin","active":true,"request_method":"GET","body":null,"port":443,"headers":[{"name":"boop","value":"beep"}]}}`
	inputHttpCheckV2Update = HttpCheckV2Input{}
)

func TestUpdateHttpCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateHttpCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateHttpCheckV2Body), &inputHttpCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateHttpCheckV2(10, &inputHttpCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputHttpCheckV2Update.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputHttpCheckV2Update.Test.Name)
	}

	if !reflect.DeepEqual(resp.Test.Customproperties, inputHttpCheckV2Update.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputHttpCheckV2Update.Test.Customproperties)
	}

	if !reflect.DeepEqual(resp.Test.Port, inputHttpCheckV2Update.Test.Port) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, inputHttpCheckV2Update.Test.Port)
	}
}

func TestUpdateHttpCheckV2HandlesEmptyResponseBody(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/11", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	err := json.Unmarshal([]byte(updateHttpCheckV2Body), &inputHttpCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, details, err := testClient.UpdateHttpCheckV2(11, &inputHttpCheckV2Update)

	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response on empty body")
	}
	if details == nil {
		t.Fatal("expected request details")
	}
}

func TestUpdateHttpCheckV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateHttpCheckV2Body), &inputHttpCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, details, err := testClient.UpdateHttpCheckV2(12, &inputHttpCheckV2Update)

	if err == nil {
		t.Fatal("expected an error on malformed JSON response, but got none")
	}
	if resp != nil && resp.Test.Name != "" {
		t.Error("expected empty response struct on parse error")
	}
	if details == nil {
		t.Fatal("expected request details even on parse error")
	}
}

func TestUpdateHttpCheckV2ReturnsErrorWhenRequestFails(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	err := json.Unmarshal([]byte(updateHttpCheckV2Body), &inputHttpCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, details, err := unreachableClient.UpdateHttpCheckV2(13, &inputHttpCheckV2Update)

	if err == nil {
		t.Fatal("expected a connection error, but got none")
	}
	if resp != nil {
		t.Errorf("expected nil response on network error, but got %#v", resp)
	}
	if details == nil {
		t.Fatal("expected request details to be populated")
	}
}

func TestUpdateHttpCheckV2WithNullablePortReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/15", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	inputWithNullablePort := HttpCheckV2InputWithNullablePort{}
	inputWithNullablePort.Test.Name = "test-http"
	resp, details, err := testClient.UpdateHttpCheckV2WithNullablePort(15, &inputWithNullablePort)

	if err == nil {
		t.Fatal("expected an error on malformed JSON response, but got none")
	}
	if resp != nil && resp.Test.Name != "" {
		t.Error("expected empty response struct on parse error")
	}
	if details == nil {
		t.Fatal("expected request details even on parse error")
	}
}

func TestUpdateHttpCheckV2WithNullablePortReturnsErrorWhenRequestFails(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	inputWithNullablePort := HttpCheckV2InputWithNullablePort{}
	inputWithNullablePort.Test.Name = "test-http"
	resp, details, err := unreachableClient.UpdateHttpCheckV2WithNullablePort(16, &inputWithNullablePort)

	if err == nil {
		t.Fatal("expected a connection error, but got none")
	}
	if resp != nil {
		t.Errorf("expected nil response on network error, but got %#v", resp)
	}
	if details == nil {
		t.Fatal("expected request details to be populated")
	}
}

func TestUpdateHttpCheckV2WithNullablePortHandlesEmptyResponseBody(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/http/14", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	inputWithNullablePort := HttpCheckV2InputWithNullablePort{}
	inputWithNullablePort.Test.Name = "test-http"
	resp, details, err := testClient.UpdateHttpCheckV2WithNullablePort(14, &inputWithNullablePort)

	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response on empty body")
	}
	if details == nil {
		t.Fatal("expected request details")
	}
}
