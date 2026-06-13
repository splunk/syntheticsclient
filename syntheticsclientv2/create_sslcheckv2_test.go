//go:build unit_tests
// +build unit_tests

// Copyright 2026 Splunk, Inc.
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
	createSslCheckV2Body = `{"test":{"name":"ssl-check","frequency":5,"schedulingStrategy":"round_robin","active":true,"locationIds":["aws-us-east-1"],"customProperties":[{"key":"env","value":"prod"}],"automaticRetries":1,"host":"www.splunk.com","port":443,"serverName":"www.splunk.com","allowSelfSigned":true,"allowUntrustedRoot":false,"caCertificateId":42,"validations":[{"name":"Certificate expires later","type":"assert_numeric","actual":"{{certificate.days_until_expiration}}","expected":"30","comparator":"is_greater_than"}]}}`
	inputSslCheckV2Data  = SslCheckV2Input{}
)

func TestCreateSslCheckV2(t *testing.T) {
	setup()
	defer teardown()

	expectedCaCertificateID := 42
	expectedValidations := []Validations{
		{
			Name:       "Certificate expires later",
			Type:       "assert_numeric",
			Actual:     "{{certificate.days_until_expiration}}",
			Expected:   "30",
			Comparator: "is_greater_than",
		},
	}
	expectedCustomProperties := []CustomProperties{
		{
			Key:   "env",
			Value: "prod",
		},
	}

	testMux.HandleFunc("/tests/ssl", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		var requestEnvelope map[string]json.RawMessage
		err = json.Unmarshal(requestBody, &requestEnvelope)
		if err != nil {
			t.Fatal(err)
		}
		rawTest, ok := requestEnvelope["test"]
		if !ok {
			t.Fatal("request body missing test envelope")
		}

		var requestTestFields map[string]json.RawMessage
		err = json.Unmarshal(rawTest, &requestTestFields)
		if err != nil {
			t.Fatal(err)
		}
		for _, field := range []string{"host", "port", "serverName", "allowSelfSigned", "allowUntrustedRoot", "caCertificateId", "validations", "customProperties"} {
			if _, ok := requestTestFields[field]; !ok {
				t.Errorf("request body missing test.%s", field)
			}
		}

		requestSslCheckV2Data := SslCheckV2Input{}
		err = json.Unmarshal(requestBody, &requestSslCheckV2Data)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.Host, "www.splunk.com") {
			t.Errorf("request body host \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Host, "www.splunk.com")
		}
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.Port, 443) {
			t.Errorf("request body port \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Port, 443)
		}
		assertStringPtr(t, requestSslCheckV2Data.Test.ServerName, "www.splunk.com")
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.AllowSelfSigned, true) {
			t.Errorf("request body allow self signed \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.AllowSelfSigned, true)
		}
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.AllowUntrustedRoot, false) {
			t.Errorf("request body allow untrusted root \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.AllowUntrustedRoot, false)
		}
		if requestSslCheckV2Data.Test.CaCertificateID == nil || !reflect.DeepEqual(*requestSslCheckV2Data.Test.CaCertificateID, expectedCaCertificateID) {
			t.Errorf("request body ca certificate id \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.CaCertificateID, expectedCaCertificateID)
		}
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.Validations, expectedValidations) {
			t.Errorf("request body validations \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Validations, expectedValidations)
		}
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.Customproperties, expectedCustomProperties) {
			t.Errorf("request body custom properties \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Customproperties, expectedCustomProperties)
		}

		_, err = w.Write([]byte(createSslCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(createSslCheckV2Body), &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.CreateSslCheckV2(&inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Test.Name, inputSslCheckV2Data.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputSslCheckV2Data.Test.Name)
	}
	if !reflect.DeepEqual(resp.Test.Host, inputSslCheckV2Data.Test.Host) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Host, inputSslCheckV2Data.Test.Host)
	}
	if !reflect.DeepEqual(resp.Test.Port, inputSslCheckV2Data.Test.Port) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, inputSslCheckV2Data.Test.Port)
	}
	if !reflect.DeepEqual(resp.Test.ServerName, inputSslCheckV2Data.Test.ServerName) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ServerName, inputSslCheckV2Data.Test.ServerName)
	}
	if !reflect.DeepEqual(resp.Test.AllowSelfSigned, inputSslCheckV2Data.Test.AllowSelfSigned) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.AllowSelfSigned, inputSslCheckV2Data.Test.AllowSelfSigned)
	}
	if !reflect.DeepEqual(resp.Test.AllowUntrustedRoot, inputSslCheckV2Data.Test.AllowUntrustedRoot) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.AllowUntrustedRoot, inputSslCheckV2Data.Test.AllowUntrustedRoot)
	}
	if !reflect.DeepEqual(resp.Test.LocationIds, inputSslCheckV2Data.Test.LocationIds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.LocationIds, inputSslCheckV2Data.Test.LocationIds)
	}
	if !reflect.DeepEqual(resp.Test.Customproperties, inputSslCheckV2Data.Test.Customproperties) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Customproperties, inputSslCheckV2Data.Test.Customproperties)
	}
	if !reflect.DeepEqual(resp.Test.Automaticretries, inputSslCheckV2Data.Test.Automaticretries) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Automaticretries, inputSslCheckV2Data.Test.Automaticretries)
	}
	if !reflect.DeepEqual(resp.Test.CaCertificateID, inputSslCheckV2Data.Test.CaCertificateID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.CaCertificateID, inputSslCheckV2Data.Test.CaCertificateID)
	}
	if !reflect.DeepEqual(resp.Test.Validations, inputSslCheckV2Data.Test.Validations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Validations, inputSslCheckV2Data.Test.Validations)
	}
}

func assertStringPtr(t *testing.T, got *string, want string) {
	t.Helper()

	if got == nil {
		t.Fatalf("got nil string pointer want \n\n%#v", want)
	}
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got string pointer value \n\n%#v want \n\n%#v", *got, want)
	}
}

func TestCreateSslCheckV2DefaultsNilValidations(t *testing.T) {
	setup()
	defer teardown()

	inputSslCheckV2Data := SslCheckV2Input{}
	err := json.Unmarshal([]byte(createSslCheckV2Body), &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}
	inputSslCheckV2Data.Test.Validations = nil

	testMux.HandleFunc("/tests/ssl", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		var requestEnvelope map[string]json.RawMessage
		err = json.Unmarshal(requestBody, &requestEnvelope)
		if err != nil {
			t.Fatal(err)
		}
		rawTest, ok := requestEnvelope["test"]
		if !ok {
			t.Fatal("request body missing test envelope")
		}

		var requestTestFields map[string]json.RawMessage
		err = json.Unmarshal(rawTest, &requestTestFields)
		if err != nil {
			t.Fatal(err)
		}
		rawValidations, ok := requestTestFields["validations"]
		if !ok {
			t.Fatal("request body missing test.validations")
		}

		var validations []json.RawMessage
		err = json.Unmarshal(rawValidations, &validations)
		if err != nil {
			t.Fatal(err)
		}
		if validations == nil {
			t.Fatalf("request body validations \n\n%#v want empty slice", string(rawValidations))
		}
		if len(validations) != 0 {
			t.Errorf("request body validations length \n\n%#v want \n\n%#v", len(validations), 0)
		}

		_, err = w.Write([]byte(createSslCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	_, _, err = testClient.CreateSslCheckV2(&inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateSslCheckV2PreservesNilServerName(t *testing.T) {
	setup()
	defer teardown()

	inputBody := strings.Replace(createSslCheckV2Body, `"serverName":"www.splunk.com"`, `"serverName":null`, 1)
	inputSslCheckV2Data := SslCheckV2Input{}
	err := json.Unmarshal([]byte(inputBody), &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/ssl", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		var requestEnvelope map[string]json.RawMessage
		err = json.Unmarshal(requestBody, &requestEnvelope)
		if err != nil {
			t.Fatal(err)
		}
		rawTest, ok := requestEnvelope["test"]
		if !ok {
			t.Fatal("request body missing test envelope")
		}

		var requestTestFields map[string]json.RawMessage
		err = json.Unmarshal(rawTest, &requestTestFields)
		if err != nil {
			t.Fatal(err)
		}
		rawServerName, ok := requestTestFields["serverName"]
		if !ok {
			t.Fatal("request body missing test.serverName")
		}
		if string(rawServerName) != "null" {
			t.Fatalf("request body server name \n\n%#v want JSON null", string(rawServerName))
		}

		_, err = w.Write([]byte(createSslCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	_, _, err = testClient.CreateSslCheckV2(&inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}
}
