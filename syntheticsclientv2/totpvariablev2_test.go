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
	createTotpVariableV2Body         = `{"totp":{"description":"TOTP for terraform","digits":6,"hmacDigest":"SHA1","interval":30,"name":"terraform-totp","secret":"create-totp-secret"}}`
	createTotpVariableV2ResponseBody = `{"totp":{"id":101,"name":"terraform-totp","description":"TOTP for terraform","secret":"<REDACTED>","digits":6,"interval":30,"hmacDigest":"SHA1","createdAt":"2026-06-21T10:15:30.001Z","createdBy":"creator","updatedAt":"2026-06-21T10:16:30.001Z","updatedBy":"updater"}}`
	inputTotpVariableV2Data          = TotpVariableV2Input{}
	outputTotpVariableV2Data         = verifyTotpVariableV2Input(createTotpVariableV2ResponseBody)

	getTotpVariableV2Body  = `{"totp":{"id":102,"name":"login-totp","description":"Login MFA","secret":"<REDACTED>","digits":8,"interval":60,"hmacDigest":"SHA256","createdAt":"2026-06-20T08:15:30.001Z","createdBy":"creator-id","updatedAt":"2026-06-21T09:16:30.001Z","updatedBy":"updater-id"}}`
	inputGetTotpVariableV2 = verifyTotpVariableV2Input(getTotpVariableV2Body)

	getTotpVariablesV2Body  = `{"totps":[{"id":201,"name":"first-totp","description":"First MFA","secret":"<REDACTED>","digits":6,"interval":30,"hmacDigest":"SHA1","createdAt":"2026-06-20T08:15:30.001Z","createdBy":"creator-one","updatedAt":"2026-06-21T09:16:30.001Z","updatedBy":"updater-one"},{"id":202,"name":"second-totp","description":"Second MFA","secret":"<REDACTED>","digits":8,"interval":60,"hmacDigest":"SHA512","createdAt":"2026-06-20T10:15:30.001Z","createdBy":"creator-two","updatedAt":"2026-06-21T11:16:30.001Z","updatedBy":"updater-two"}]}`
	inputGetTotpVariablesV2 = verifyTotpVariablesV2Input(getTotpVariablesV2Body)

	updateTotpVariableV2ResponseBody = `{"totp":{"id":103,"name":"updated-totp","description":"Updated TOTP","secret":"<REDACTED>","digits":8,"interval":45,"hmacDigest":"SHA512","createdAt":"2026-06-20T08:15:30.001Z","createdBy":"creator-id","updatedAt":"2026-06-21T09:16:30.001Z","updatedBy":"updater-id"}}`
	outputUpdateTotpVariableV2Data   = verifyTotpVariableV2Input(updateTotpVariableV2ResponseBody)
)

func TestCreateTotpVariableV2(t *testing.T) {
	setup()
	defer teardown()

	err := json.Unmarshal([]byte(createTotpVariableV2Body), &inputTotpVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/totps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, requestTotpFields := readTotpRequestFields(t, r)
		assertTotpRequestFields(t, requestTotpFields, map[string]string{
			"description": `"TOTP for terraform"`,
			"digits":      `6`,
			"hmacDigest":  `"SHA1"`,
			"interval":    `30`,
			"name":        `"terraform-totp"`,
			"secret":      `"create-totp-secret"`,
		})
		assertTotpResponseOnlyFieldsAbsent(t, requestTotpFields)

		_, err := w.Write([]byte(createTotpVariableV2ResponseBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.CreateTotpVariableV2(&inputTotpVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Totp, outputTotpVariableV2Data.Totp) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Totp, outputTotpVariableV2Data.Totp)
	}
	assertDoesNotContain(t, details.RequestBody, "create-totp-secret")
	assertContains(t, details.RequestBody, `"secret":"[REDACTED]"`, `"hmacDigest":"SHA1"`)
	assertRawRequestDoesNotExposeSensitiveDetails(t, details, "apiKey", "create-totp-secret")
}

func TestGetTotpVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/totps/102", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getTotpVariableV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetTotpVariableV2(102)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Totp, inputGetTotpVariableV2.Totp) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Totp, inputGetTotpVariableV2.Totp)
	}
	if resp.Totp.Secret != "<REDACTED>" {
		t.Errorf("returned secret \n\n%#v want \n\n%#v", resp.Totp.Secret, "<REDACTED>")
	}
}

func TestGetTotpVariablesV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/totps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getTotpVariablesV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetTotpVariablesV2()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Totps, inputGetTotpVariablesV2.Totps) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Totps, inputGetTotpVariablesV2.Totps)
	}
}

func TestUpdateTotpVariableV2(t *testing.T) {
	setup()
	defer teardown()

	description := "Updated TOTP"
	digits := 8
	hmacDigest := "SHA512"
	interval := 45
	secret := "update-totp-secret"
	updateInput := TotpVariableV2UpdateInput{
		Totp: TotpVariableUpdateInput{
			Description: &description,
			Digits:      &digits,
			HmacDigest:  &hmacDigest,
			Interval:    &interval,
			Secret:      &secret,
		},
	}

	testMux.HandleFunc("/totps/103", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, requestTotpFields := readTotpRequestFields(t, r)
		assertTotpRequestFields(t, requestTotpFields, map[string]string{
			"description": `"Updated TOTP"`,
			"digits":      `8`,
			"hmacDigest":  `"SHA512"`,
			"interval":    `45`,
			"secret":      `"update-totp-secret"`,
		})
		assertTotpUpdateForbiddenFieldsAbsent(t, requestTotpFields)

		_, err := w.Write([]byte(updateTotpVariableV2ResponseBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.UpdateTotpVariableV2(103, &updateInput)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Totp, outputUpdateTotpVariableV2Data.Totp) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Totp, outputUpdateTotpVariableV2Data.Totp)
	}
	assertDoesNotContain(t, details.RequestBody, "update-totp-secret")
	assertContains(t, details.RequestBody, `"secret":"[REDACTED]"`, `"hmacDigest":"SHA512"`)
	assertRawRequestDoesNotExposeSensitiveDetails(t, details, "apiKey", "update-totp-secret")
}

func TestUpdateTotpVariableV2OmitsNilSecret(t *testing.T) {
	setup()
	defer teardown()

	description := ""
	digits := 8
	hmacDigest := "SHA256"
	interval := 60
	updateInput := TotpVariableV2UpdateInput{
		Totp: TotpVariableUpdateInput{
			Description: &description,
			Digits:      &digits,
			HmacDigest:  &hmacDigest,
			Interval:    &interval,
		},
	}

	testMux.HandleFunc("/totps/104", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, requestTotpFields := readTotpRequestFields(t, r)
		assertTotpRequestFields(t, requestTotpFields, map[string]string{
			"description": `""`,
			"digits":      `8`,
			"hmacDigest":  `"SHA256"`,
			"interval":    `60`,
		})
		assertTotpUpdateForbiddenFieldsAbsent(t, requestTotpFields)
		if _, ok := requestTotpFields["secret"]; ok {
			t.Fatal("request body should omit totp.secret when Secret is nil")
		}

		w.WriteHeader(http.StatusOK)
	})

	resp, _, err := testClient.UpdateTotpVariableV2(104, &updateInput)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for blank successful update body")
	}
}

func TestDeleteTotpVariableV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/totps/105", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.DeleteTotpVariableV2(105)
	if err != nil {
		t.Fatal(err)
	}
	if resp != http.StatusNoContent {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, http.StatusNoContent)
	}
}

func TestCreateTotpVariableV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	err := json.Unmarshal([]byte(createTotpVariableV2Body), &inputTotpVariableV2Data)
	if err != nil {
		t.Fatal(err)
	}

	testMux.HandleFunc("/totps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.CreateTotpVariableV2(&inputTotpVariableV2Data)
	if err == nil {
		t.Fatal("expected error on malformed response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on parse error, got %#v", resp)
	}
}

func TestCreateTotpVariableV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	inputData := TotpVariableV2Input{
		Totp: TotpVariableInput{
			Name:       "test-totp",
			Secret:     "test-secret",
			Digits:     6,
			Interval:   30,
			HmacDigest: "SHA1",
		},
	}

	_, _, err := unreachableClient.CreateTotpVariableV2(&inputData)
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestGetTotpVariableV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/totps/102", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{invalid json}"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetTotpVariableV2(102)
	if err == nil {
		t.Fatal("expected error on malformed response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on parse error, got %#v", resp)
	}
}

func TestGetTotpVariableV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	_, _, err := unreachableClient.GetTotpVariableV2(102)
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestGetTotpVariablesV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/totps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{invalid json array}"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetTotpVariablesV2()
	if err == nil {
		t.Fatal("expected error on malformed response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on parse error, got %#v", resp)
	}
}

func TestGetTotpVariablesV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	_, _, err := unreachableClient.GetTotpVariablesV2()
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestUpdateTotpVariableV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	description := "Updated TOTP"
	digits := 8
	hmacDigest := "SHA512"
	interval := 45
	secret := "update-totp-secret"
	updateInput := TotpVariableV2UpdateInput{
		Totp: TotpVariableUpdateInput{
			Description: &description,
			Digits:      &digits,
			HmacDigest:  &hmacDigest,
			Interval:    &interval,
			Secret:      &secret,
		},
	}

	testMux.HandleFunc("/totps/103", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte("{malformed response}"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.UpdateTotpVariableV2(103, &updateInput)
	if err == nil {
		t.Fatal("expected error on malformed response, got nil")
	}
	if resp != nil {
		t.Errorf("expected nil response on parse error, got %#v", resp)
	}
}

func TestUpdateTotpVariableV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	description := "Updated TOTP"
	digits := 8
	hmacDigest := "SHA512"
	interval := 45
	secret := "update-totp-secret"
	updateInput := TotpVariableV2UpdateInput{
		Totp: TotpVariableUpdateInput{
			Description: &description,
			Digits:      &digits,
			HmacDigest:  &hmacDigest,
			Interval:    &interval,
			Secret:      &secret,
		},
	}

	_, _, err := unreachableClient.UpdateTotpVariableV2(103, &updateInput)
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestUpdateTotpVariableV2ReturnsNonNilResponseOnEmptyBody(t *testing.T) {
	setup()
	defer teardown()

	description := "Updated TOTP"
	digits := 8
	hmacDigest := "SHA512"
	interval := 45
	updateInput := TotpVariableV2UpdateInput{
		Totp: TotpVariableUpdateInput{
			Description: &description,
			Digits:      &digits,
			HmacDigest:  &hmacDigest,
			Interval:    &interval,
		},
	}

	testMux.HandleFunc("/totps/106", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(""))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.UpdateTotpVariableV2(106, &updateInput)
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non-nil response for empty body with 2xx status")
	}
}

func TestDeleteTotpVariableV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})

	_, err := unreachableClient.DeleteTotpVariableV2(105)
	if err == nil {
		t.Fatal("expected connection error, got nil")
	}
}

func TestDeleteTotpVariableV2ReturnsErrorOnNon2xxStatus(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/totps/107", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusMultipleChoices)
	})

	resp, err := testClient.DeleteTotpVariableV2(107)
	if err == nil {
		t.Fatal("expected error on non-2xx status code, got nil")
	}
	if resp != http.StatusMultipleChoices {
		t.Errorf("returned status \n\n%#v want \n\n%#v", resp, http.StatusMultipleChoices)
	}
	if !strings.Contains(err.Error(), "Response code") {
		t.Errorf("expected error to contain 'Response code', got %s", err.Error())
	}
}

func verifyTotpVariableV2Input(stringInput string) *TotpVariableV2Response {
	check := &TotpVariableV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func verifyTotpVariablesV2Input(stringInput string) *TotpVariablesV2Response {
	check := &TotpVariablesV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func readTotpRequestFields(t *testing.T, r *http.Request) ([]byte, map[string]json.RawMessage) {
	t.Helper()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	var requestEnvelope map[string]json.RawMessage
	err = json.Unmarshal(requestBody, &requestEnvelope)
	if err != nil {
		t.Fatal(err)
	}
	if len(requestEnvelope) != 1 {
		t.Fatalf("request body envelope keys \n\n%#v want only totp", requestEnvelope)
	}
	rawTotp, ok := requestEnvelope["totp"]
	if !ok {
		t.Fatal("request body missing totp envelope")
	}

	var requestTotpFields map[string]json.RawMessage
	err = json.Unmarshal(rawTotp, &requestTotpFields)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := requestTotpFields["hmac_digest"]; ok {
		t.Fatal("request body should use hmacDigest, not hmac_digest")
	}
	if _, ok := requestTotpFields["hmacdigest"]; ok {
		t.Fatal("request body should use hmacDigest, not hmacdigest")
	}

	return requestBody, requestTotpFields
}

func assertTotpRequestFields(t *testing.T, requestTotpFields map[string]json.RawMessage, expectedFields map[string]string) {
	t.Helper()

	if len(requestTotpFields) != len(expectedFields) {
		t.Errorf("request body totp field count %d want %d", len(requestTotpFields), len(expectedFields))
	}
	for field, expected := range expectedFields {
		rawValue, ok := requestTotpFields[field]
		if !ok {
			t.Errorf("request body missing totp.%s", field)
			continue
		}
		if string(rawValue) != expected {
			t.Errorf("request body totp.%s %s want %s", field, rawValue, expected)
		}
	}
}

func assertTotpResponseOnlyFieldsAbsent(t *testing.T, requestTotpFields map[string]json.RawMessage) {
	t.Helper()

	for _, field := range []string{"id", "createdAt", "createdBy", "updatedAt", "updatedBy"} {
		if _, ok := requestTotpFields[field]; ok {
			t.Errorf("request body should not include totp.%s", field)
		}
	}
}

func assertTotpUpdateForbiddenFieldsAbsent(t *testing.T, requestTotpFields map[string]json.RawMessage) {
	t.Helper()

	assertTotpResponseOnlyFieldsAbsent(t, requestTotpFields)
	if _, ok := requestTotpFields["name"]; ok {
		t.Error("request body should not include immutable totp.name")
	}
}

func assertRawRequestDoesNotExposeSensitiveDetails(t *testing.T, details *RequestDetails, apiKey string, secrets ...string) {
	t.Helper()

	if details.RawRequest == nil {
		return
	}

	if details.RawRequest.Header.Get("X-SF-TOKEN") == apiKey {
		t.Errorf("raw request exposes real X-SF-TOKEN header")
	}

	if details.RawRequest.GetBody == nil {
		return
	}

	body, err := details.RawRequest.GetBody()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = body.Close() }()

	rawBody, err := io.ReadAll(body)
	if err != nil {
		t.Fatal(err)
	}
	assertDoesNotContain(t, string(rawBody), secrets...)
}

func assertContains(t *testing.T, text string, values ...string) {
	t.Helper()

	for _, value := range values {
		if !strings.Contains(text, value) {
			t.Errorf("expected text to contain %q\ntext:\n%s", value, text)
		}
	}
}

func assertDoesNotContain(t *testing.T, text string, values ...string) {
	t.Helper()

	for _, value := range values {
		if strings.Contains(text, value) {
			t.Errorf("expected text to omit %q\ntext:\n%s", value, text)
		}
	}
}
