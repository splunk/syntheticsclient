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
// recurrence.repeats.unit and recurrence.repeats.value. Sending customFrequency
// and customValue instead got them dropped by strong parameters, so creating a
// custom recurring downtime failed with 422 "can't be blank".
func TestCreateDowntimeConfigurationV2SendsCustomRecurrenceAsUnitAndValue(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/downtime_configurations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fields := readDowntimeConfigurationV2RepeatsFields(t, r)

		rawUnit, ok := fields["unit"]
		if !ok {
			t.Fatal("request body missing recurrence.repeats.unit")
		}
		if string(rawUnit) != `"days"` {
			t.Fatalf("request body recurrence.repeats.unit = %s, want \"days\"", rawUnit)
		}

		rawValue, ok := fields["value"]
		if !ok {
			t.Fatal("request body missing recurrence.repeats.value")
		}
		if string(rawValue) != "3" {
			t.Fatalf("request body recurrence.repeats.value = %s, want 3", rawValue)
		}

		if _, ok := fields["customFrequency"]; ok {
			t.Error("request body sent recurrence.repeats.customFrequency, which the API does not accept")
		}
		if _, ok := fields["customValue"]; ok {
			t.Error("request body sent recurrence.repeats.customValue, which the API does not accept")
		}

		_, err := w.Write([]byte(`{"downtimeConfiguration":{"id":1,"name":"dc test","rule":"augment_data","recurrence":{"repeats":{"type":"custom","unit":"days","value":3},"end":{"type":"after","value":"10"}}}}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	customFrequency := "days"
	customValue := 3
	input := DowntimeConfigurationV2Input{}
	input.Name = "dc test"
	input.Rule = "augment_data"
	input.Recurrence = &Recurrence{
		Repeats: Repeats{
			Type:            "custom",
			Customfrequency: &customFrequency,
			Customvalue:     &customValue,
		},
		End: &End{Type: "after", Value: "10"},
	}

	resp, _, err := testClient.CreateDowntimeConfigurationV2(&input)
	if err != nil {
		t.Fatal(err)
	}

	repeats := resp.Recurrence.Repeats
	if repeats.Customfrequency == nil || *repeats.Customfrequency != "days" {
		t.Errorf("response repeats.Customfrequency = %#v, want days", repeats.Customfrequency)
	}
	if repeats.Customvalue == nil || *repeats.Customvalue != 3 {
		t.Errorf("response repeats.Customvalue = %#v, want 3", repeats.Customvalue)
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

		_, err := w.Write([]byte(`{"downtimeConfiguration":{"id":2,"name":"dc test","rule":"augment_data","recurrence":{"repeats":{"type":"daily"}}}}`))
		if err != nil {
			t.Fatal(err)
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
	if repeats.Customfrequency != nil {
		t.Errorf("response repeats.Customfrequency = %#v, want nil", repeats.Customfrequency)
	}
	if repeats.Customvalue != nil {
		t.Errorf("response repeats.Customvalue = %#v, want nil", repeats.Customvalue)
	}
}

func readDowntimeConfigurationV2RepeatsFields(t *testing.T, r *http.Request) map[string]json.RawMessage {
	t.Helper()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	var envelope struct {
		DowntimeConfiguration struct {
			Recurrence struct {
				Repeats map[string]json.RawMessage `json:"repeats"`
			} `json:"recurrence"`
		} `json:"downtimeConfiguration"`
	}
	if err := json.Unmarshal(requestBody, &envelope); err != nil {
		t.Fatalf("request body JSON parse failed: %v", err)
	}

	return envelope.DowntimeConfiguration.Recurrence.Repeats
}
