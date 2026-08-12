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
	"net/http"
	"reflect"
	"strings"
	"testing"
)

const getSslCheckV2LastRunID = "77ff08f0-3865-42d8-9fff-3efb7c25e176"

var (
	getSslCheckV2Body  = `{"test":{"id":1647,"type":"ssl","name":"ssl-check","frequency":5,"schedulingStrategy":"round_robin","active":true,"locationIds":["aws-us-east-1"],"createdAt":"2022-09-14T14:35:37.801Z","updatedAt":"2022-09-14T14:35:38.099Z","createdBy":"abc1234","updatedBy":"abc1234","customProperties":[{"key":"env","value":"prod"}],"automaticRetries":1,"lastRunStatus":"success","lastRunAt":"2024-03-07T00:47:43.741Z","lastRunCoreMetricsPublishedAt":"2024-03-07T00:47:43.741Z","lastRunLocationId":"aws-us-east-1","lastRunId":"` + getSslCheckV2LastRunID + `","host":"www.splunk.com","port":443,"serverName":"www.splunk.com","allowSelfSigned":true,"allowUntrustedRoot":false,"caCertificateId":42,"validations":[{"name":"Certificate expires later","type":"assert_numeric","actual":"{{certificate.days_until_expiration}}","expected":"30","comparator":"is_greater_than"}]}}`
	inputGetSslCheckV2 = verifySslCheckV2Input(string(getSslCheckV2Body))
)

func TestGetSslCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/ssl/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getSslCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetSslCheckV2(1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Test.ID, inputGetSslCheckV2.Test.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ID, inputGetSslCheckV2.Test.ID)
	}
	if resp.Test.LastRunId != getSslCheckV2LastRunID {
		t.Errorf("returned last run ID %q, want %q", resp.Test.LastRunId, getSslCheckV2LastRunID)
	}
	if !reflect.DeepEqual(resp.Test.Name, inputGetSslCheckV2.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputGetSslCheckV2.Test.Name)
	}
	if !reflect.DeepEqual(resp.Test.Type, inputGetSslCheckV2.Test.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Type, inputGetSslCheckV2.Test.Type)
	}
	if !reflect.DeepEqual(resp.Test.Host, inputGetSslCheckV2.Test.Host) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Host, inputGetSslCheckV2.Test.Host)
	}
	if !reflect.DeepEqual(resp.Test.Port, inputGetSslCheckV2.Test.Port) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Port, inputGetSslCheckV2.Test.Port)
	}
	if !reflect.DeepEqual(resp.Test.ServerName, inputGetSslCheckV2.Test.ServerName) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.ServerName, inputGetSslCheckV2.Test.ServerName)
	}
	if !reflect.DeepEqual(resp.Test.AllowSelfSigned, inputGetSslCheckV2.Test.AllowSelfSigned) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.AllowSelfSigned, inputGetSslCheckV2.Test.AllowSelfSigned)
	}
	if !reflect.DeepEqual(resp.Test.AllowUntrustedRoot, inputGetSslCheckV2.Test.AllowUntrustedRoot) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.AllowUntrustedRoot, inputGetSslCheckV2.Test.AllowUntrustedRoot)
	}
	if !reflect.DeepEqual(resp.Test.CaCertificateID, inputGetSslCheckV2.Test.CaCertificateID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.CaCertificateID, inputGetSslCheckV2.Test.CaCertificateID)
	}
	if !reflect.DeepEqual(resp.Test.Validations, inputGetSslCheckV2.Test.Validations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Validations, inputGetSslCheckV2.Test.Validations)
	}
}

func TestGetSslCheckV2PreservesNilServerName(t *testing.T) {
	setup()
	defer teardown()

	responseBody := strings.Replace(getSslCheckV2Body, `"serverName":"www.splunk.com"`, `"serverName":null`, 1)

	testMux.HandleFunc("/tests/ssl/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(responseBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetSslCheckV2(1)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Test.ServerName != nil {
		t.Fatalf("returned server name \n\n%#v want nil", resp.Test.ServerName)
	}
}

func verifySslCheckV2Input(stringInput string) *SslCheckV2Response {
	check := &SslCheckV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestGetSslCheckV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/ssl/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetSslCheckV2(1)

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

func TestGetSslCheckV2ReturnsErrorWhenRequestFails(t *testing.T) {
	// Use an unreachable address to trigger a connection error
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	_, _, err := unreachableClient.GetSslCheckV2(1)

	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}
