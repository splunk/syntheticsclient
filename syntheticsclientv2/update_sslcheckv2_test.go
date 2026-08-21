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
	inputSslCheckV2Update = SslCheckV2UpdateInput{}
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

		requestSslCheckV2Data := SslCheckV2UpdateInput{}
		err = json.Unmarshal(requestBody, &requestSslCheckV2Data)
		if err != nil {
			t.Fatal(err)
		}
		if requestSslCheckV2Data.Test.Host == nil || !reflect.DeepEqual(*requestSslCheckV2Data.Test.Host, "example.com") {
			t.Errorf("request body host \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Host, "example.com")
		}
		if requestSslCheckV2Data.Test.Port == nil || !reflect.DeepEqual(*requestSslCheckV2Data.Test.Port, 8443) {
			t.Errorf("request body port \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Port, 8443)
		}
		if requestSslCheckV2Data.Test.ServerName == nil ||
			requestSslCheckV2Data.Test.ServerName.Value == nil ||
			!reflect.DeepEqual(*requestSslCheckV2Data.Test.ServerName.Value, "example.com") {
			t.Errorf("request body server name \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.ServerName, "example.com")
		}
		if requestSslCheckV2Data.Test.AllowSelfSigned == nil || !reflect.DeepEqual(*requestSslCheckV2Data.Test.AllowSelfSigned, false) {
			t.Errorf("request body allow self signed \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.AllowSelfSigned, false)
		}
		if requestSslCheckV2Data.Test.AllowUntrustedRoot == nil || !reflect.DeepEqual(*requestSslCheckV2Data.Test.AllowUntrustedRoot, true) {
			t.Errorf("request body allow untrusted root \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.AllowUntrustedRoot, true)
		}
		if requestSslCheckV2Data.Test.CaCertificateID == nil ||
			requestSslCheckV2Data.Test.CaCertificateID.Value == nil ||
			!reflect.DeepEqual(*requestSslCheckV2Data.Test.CaCertificateID.Value, expectedCaCertificateID) {
			t.Errorf("request body ca certificate id \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.CaCertificateID, expectedCaCertificateID)
		}
		if requestSslCheckV2Data.Test.Validations == nil || !reflect.DeepEqual(*requestSslCheckV2Data.Test.Validations, expectedValidations) {
			t.Errorf("request body validations \n\n%#v want \n\n%#v", requestSslCheckV2Data.Test.Validations, expectedValidations)
		}
		if requestSslCheckV2Data.Test.Customproperties == nil || !reflect.DeepEqual(*requestSslCheckV2Data.Test.Customproperties, expectedCustomProperties) {
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

	if !reflect.DeepEqual(resp.Test.Name, "ssl-check-updated") {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, "ssl-check-updated")
	}
	if !reflect.DeepEqual(resp.Test.Host, "example.com") {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Host, "example.com")
	}
	if !reflect.DeepEqual(resp.Test.Port, 8443) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, 8443)
	}
	if resp.Test.CaCertificateID == nil || !reflect.DeepEqual(*resp.Test.CaCertificateID, expectedCaCertificateID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.CaCertificateID, expectedCaCertificateID)
	}
	if !reflect.DeepEqual(resp.Test.Validations, expectedValidations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Validations, expectedValidations)
	}
}

func TestUpdateSslCheckV2AllowsPartialPayload(t *testing.T) {
	setup()
	defer teardown()

	active := false
	inputSslCheckV2Update := SslCheckV2UpdateInput{}
	inputSslCheckV2Update.Test.Active = &active

	testMux.HandleFunc("/tests/ssl/1654", func(w http.ResponseWriter, r *http.Request) {
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

		if len(requestTestFields) != 1 {
			t.Fatalf("request body test field count %d want 1: %s", len(requestTestFields), requestBody)
		}
		rawActive, ok := requestTestFields["active"]
		if !ok {
			t.Fatal("request body missing test.active")
		}
		if string(rawActive) != "false" {
			t.Fatalf("request body test.active %s want false", rawActive)
		}

		w.WriteHeader(http.StatusOK)
	})

	resp, _, err := testClient.UpdateSslCheckV2(1654, &inputSslCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}

func TestUpdateSslCheckV2AllowsEmptyValidations(t *testing.T) {
	setup()
	defer teardown()

	inputSslCheckV2Data := SslCheckV2UpdateInput{}
	validations := make([]Validations, 0)
	inputSslCheckV2Data.Test.Validations = &validations

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

	_, _, err := testClient.UpdateSslCheckV2(1651, &inputSslCheckV2Data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSslCheckV2PreservesNilServerName(t *testing.T) {
	setup()
	defer teardown()

	inputSslCheckV2Data := SslCheckV2UpdateInput{}
	inputSslCheckV2Data.Test.ServerName = NewNullString()

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

	_, _, err := testClient.UpdateSslCheckV2(1653, &inputSslCheckV2Data)
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

func TestUpdateSslCheckV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/ssl/1650", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateSslCheckV2Body), &inputSslCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateSslCheckV2(1650, &inputSslCheckV2Update)

	if err == nil {
		t.Fatal("expected error on malformed JSON response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on error, got %#v", resp)
	}
	if !strings.Contains(err.Error(), "invalid character") {
		t.Errorf("expected JSON unmarshal error, got: %v", err)
	}
}

func TestUpdateSslCheckV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	err := json.Unmarshal([]byte(updateSslCheckV2Body), &inputSslCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = unreachableClient.UpdateSslCheckV2(1650, &inputSslCheckV2Update)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
