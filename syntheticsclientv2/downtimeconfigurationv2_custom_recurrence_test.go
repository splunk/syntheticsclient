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
	"testing"
)

// SYN-7042: the downtime API names the custom recurrence fields
// recurrence.repeats.unit and recurrence.repeats.value. Sending repeatUnit
// and repeatValue instead got them dropped by strong parameters, so creating a
// custom recurring downtime failed with 422 "can't be blank".
func TestCreateDowntimeConfigurationV2SendsCustomRecurrenceAsUnitAndValue(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fields := readDowntimeConfigurationV2RepeatsFields(t, r)

		if rawUnit, ok := fields["unit"]; !ok {
			t.Error("request body missing recurrence.repeats.unit")
		} else if string(rawUnit) != `"days"` {
			t.Errorf("request body recurrence.repeats.unit = %s, want \"days\"", rawUnit)
		}

		if rawValue, ok := fields["value"]; !ok {
			t.Error("request body missing recurrence.repeats.value")
		} else if string(rawValue) != "3" {
			t.Errorf("request body recurrence.repeats.value = %s, want 3", rawValue)
		}

		if _, ok := fields["repeatUnit"]; ok {
			t.Error("request body sent recurrence.repeats.repeatUnit, which the API does not accept")
		}
		if _, ok := fields["repeatValue"]; ok {
			t.Error("request body sent recurrence.repeats.repeatValue, which the API does not accept")
		}

		if _, err := w.Write([]byte(`{"downtimeConfiguration":{"id":1,"name":"dc test","rule":"augment_data","recurrence":{"repeats":{"type":"custom","unit":"days","value":3},"end":{"type":"after","value":"10"}}}}`)); err != nil {
			t.Errorf("writing mock response: %v", err)
		}
	})

	repeatUnit := "days"
	repeatValue := 3
	input := DowntimeConfigurationV2Input{}
	input.Name = "dc test"
	input.Rule = "augment_data"
	input.Recurrence = &Recurrence{
		Repeats: Repeats{
			Type:  "custom",
			Unit:  &repeatUnit,
			Value: &repeatValue,
		},
		End: &End{Type: "after", Value: "10"},
	}

	resp, _, err := testClient.CreateDowntimeConfigurationV2(&input)
	if err != nil {
		t.Fatal(err)
	}

	repeats := resp.Recurrence.Repeats
	if repeats.Unit == nil || *repeats.Unit != "days" {
		t.Errorf("response repeats.Unit = %#v, want days", repeats.Unit)
	}
	if repeats.Value == nil || *repeats.Value != 3 {
		t.Errorf("response repeats.Value = %#v, want 3", repeats.Value)
	}
}

// The API only returns unit/value for a custom recurrence, so a non-custom
// recurrence must not send them either.
func TestCreateDowntimeConfigurationV2OmitsUnitAndValueForNonCustomRecurrence(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fields := readDowntimeConfigurationV2RepeatsFields(t, r)

		if rawUnit, ok := fields["unit"]; ok {
			t.Errorf("request body recurrence.repeats.unit = %s, want field omitted", rawUnit)
		}
		if rawValue, ok := fields["value"]; ok {
			t.Errorf("request body recurrence.repeats.value = %s, want field omitted", rawValue)
		}

		if _, err := w.Write([]byte(`{"downtimeConfiguration":{"id":2,"name":"dc test","rule":"augment_data","recurrence":{"repeats":{"type":"daily"}}}}`)); err != nil {
			t.Errorf("writing mock response: %v", err)
		}
	})

	input := DowntimeConfigurationV2Input{}
	input.Name = "dc test"
	input.Rule = "augment_data"
	input.Recurrence = &Recurrence{Repeats: Repeats{Type: "daily"}}

	resp, _, err := testClient.CreateDowntimeConfigurationV2(&input)
	if err != nil {
		t.Fatal(err)
	}

	repeats := resp.Recurrence.Repeats
	if repeats.Unit != nil {
		t.Errorf("response repeats.Unit = %#v, want nil", repeats.Unit)
	}
	if repeats.Value != nil {
		t.Errorf("response repeats.Value = %#v, want nil", repeats.Value)
	}
}

// Runs on the test server's handler goroutine, where t.Fatal is not allowed, so
// failures are reported with t.Errorf and a nil map is returned; the caller's
// lookups then report the individual fields as missing.
func readDowntimeConfigurationV2RepeatsFields(t *testing.T, r *http.Request) map[string]json.RawMessage {
	t.Helper()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("reading request body: %v", err)
		return nil
	}

	var envelope struct {
		DowntimeConfiguration struct {
			Recurrence struct {
				Repeats map[string]json.RawMessage `json:"repeats"`
			} `json:"recurrence"`
		} `json:"downtimeConfiguration"`
	}
	if err := json.Unmarshal(requestBody, &envelope); err != nil {
		t.Errorf("request body JSON parse failed: %v", err)
		return nil
	}

	return envelope.DowntimeConfiguration.Recurrence.Repeats
}
