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
	updateSslCheckV2Body  = `{"test":{"name":"ssl-check-updated","frequency":10,"schedulingStrategy":"concurrent","active":false,"locationIds":["aws-us-east-1"],"customProperties":[{"key":"env","value":"stage"}],"automaticRetries":2,"host":"example.com","port":8443,"serverName":"example.com","allowSelfSigned":false,"allowUntrustedRoot":true,"caCertificateId":42,"validations":[{"name":"Certificate expires later","type":"assert_numeric","actual":"{{certificate.days_until_expiration}}","expected":"15","comparator":"is_greater_than"}]}}`
	inputSslCheckV2Update = SslCheckV2Input{}
)

func TestUpdateSslCheckV2(t *testing.T) {
	setup()
	defer teardown()

	expectedCaCertificateID := 42
	expectedValidations := []Validations{
		{
			Name:       "Certificate expires later",
			Type:       "assert_numeric",
			Actual:     "{{certificate.days_until_expiration}}",
			Expected:   "15",
			Comparator: "is_greater_than",
		},
	}
	expectedCustomProperties := []CustomProperties{
		{
			Key:   "env",
			Value: "stage",
		},
	}

	testMux.HandleFunc("/tests/ssl/1650", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
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
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.Host, "example.com") {
			t.Errorf("request body host \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Host, "example.com")
		}
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.Port, 8443) {
			t.Errorf("request body port \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Port, 8443)
		}
		assertStringPtr(t, requestSslCheckV2Data.Test.ServerName, "example.com")
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.AllowSelfSigned, false) {
			t.Errorf("request body allow self signed \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.AllowSelfSigned, false)
		}
		if !reflect.DeepEqual(requestSslCheckV2Data.Test.AllowUntrustedRoot, true) {
			t.Errorf("request body allow untrusted root \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.AllowUntrustedRoot, true)
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

		_, err = w.Write([]byte(updateSslCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateSslCheckV2Body), &inputSslCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateSslCheckV2(1650, &inputSslCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Test.Name, inputSslCheckV2Update.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputSslCheckV2Update.Test.Name)
	}
	if !reflect.DeepEqual(resp.Test.Host, inputSslCheckV2Update.Test.Host) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Host, inputSslCheckV2Update.Test.Host)
	}
	if !reflect.DeepEqual(resp.Test.Port, inputSslCheckV2Update.Test.Port) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, inputSslCheckV2Update.Test.Port)
	}
	if !reflect.DeepEqual(resp.Test.CaCertificateID, inputSslCheckV2Update.Test.CaCertificateID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.CaCertificateID, inputSslCheckV2Update.Test.CaCertificateID)
	}
	if !reflect.DeepEqual(resp.Test.Validations, inputSslCheckV2Update.Test.Validations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Validations, inputSslCheckV2Update.Test.Validations)
	}
}

func TestUpdateSslCheckV2DefaultsNilValidations(t *testing.T) {
	setup()
	defer teardown()

	inputSslCheckV2Data := SslCheckV2Input{}
	err := json.Unmarshal([]byte(updateSslCheckV2Body), &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}
	inputSslCheckV2Data.Test.Validations = nil

	testMux.HandleFunc("/tests/ssl/1651", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
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

		_, err = w.Write([]byte(updateSslCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	_, _, err = testClient.UpdateSslCheckV2(1651, &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSslCheckV2PreservesNilServerName(t *testing.T) {
	setup()
	defer teardown()

	inputBody := strings.Replace(updateSslCheckV2Body, `"serverName":"example.com"`, `"serverName":null`, 1)
	inputSslCheckV2Data := SslCheckV2Input{}
	err := json.Unmarshal([]byte(inputBody), &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/tests/ssl/1653", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
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

		_, err = w.Write([]byte(updateSslCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	_, _, err = testClient.UpdateSslCheckV2(1653, &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSslCheckV2BlankResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/ssl/1652", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
	})

	err := json.Unmarshal([]byte(updateSslCheckV2Body), &inputSslCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateSslCheckV2(1652, &inputSslCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}
