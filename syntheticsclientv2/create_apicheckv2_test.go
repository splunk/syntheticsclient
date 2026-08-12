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
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var (
	createApiV2Body = `{"test":{"automaticRetries": 1,"customProperties": [{"key": "Test_Key", "value": "Test Custom Properties"}], "active":true,"deviceId":1,"frequency":5,"location_ids":["aws-us-east-1"],"name":"boop-test","scheduling_strategy":"round_robin","requests":[{"configuration":{"name":"Get-Test","requestMethod":"GET","url":"https://api.us1.signalfx.com/v2/synthetics/v2/tests/api/489","certificateId":123,"headers":{"X-SF-TOKEN":"jinglebellsbatmanshells", "beep":"plain-header-secret"},"body":null},"setup":[{"name":"Extract from response body","type":"extract_json","source":"{{response.body}}","extractor":"$.requests","variable":"custom-varz"}],"validations":[{"name":"Assert response code equals 200","type":"assert_numeric","actual":"{{response.code}}","expected":"200","comparator":"equals"}]}]}}`
	inputData       = ApiCheckV2Input{}
)

func TestCreateApiCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/tests/api", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(requestBody), `"certificateId":123`) {
			t.Fatalf("request body missing certificateId: %s", requestBody)
		}
		if strings.Contains(string(requestBody), "certificate_id") {
			t.Fatalf("request body contains internal certificate_id field: %s", requestBody)
		}
		_, err = w.Write([]byte(createApiV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createApiV2Body), &inputData)
	if err != nil {
		t.Fatal(err)
	}

	resp, details, err := testClient.CreateApiCheckV2(&inputData)

	if err != nil {
		t.Fatal(err)
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	for _, secret := range []string{"jinglebellsbatmanshells", "plain-header-secret"} {
		if strings.Contains(details.RequestBody, secret) {
			t.Fatalf("expected request details to redact API check header value %q, but saw: %s", secret, details.RequestBody)
		}
	}
	for _, redacted := range []string{`"X-SF-TOKEN":"[REDACTED]"`, `"beep":"[REDACTED]"`} {
		if !strings.Contains(details.RequestBody, redacted) {
			t.Fatalf("expected request details to contain redacted header %q, but saw: %s", redacted, details.RequestBody)
		}
	}

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

	if !reflect.DeepEqual(resp.Test.Customproperties, inputData.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputData.Test.Customproperties)
	}
}

func TestCreateApiCheckV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/tests/api", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createApiV2Body), &inputData)
	if err != nil {
		t.Fatal(err)
	}

	resp, details, err := testClient.CreateApiCheckV2(&inputData)

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

func TestCreateApiCheckV2ReturnsErrorWhenRequestFails(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	err := json.Unmarshal([]byte(createApiV2Body), &inputData)
	if err != nil {
		t.Fatal(err)
	}

	resp, details, err := unreachableClient.CreateApiCheckV2(&inputData)

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
