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
	"net/http"
	"reflect"
	"testing"
)

var (
	getLocationsV2Body  = `{"locations":[{"id":"aws-af-south-1","label":"AWS - Cape Town","default":false,"type":"public","country":"ZA"},{"id":"aws-ap-east-1","label":"AWS - Hong Kong","default":false,"type":"public","country":"HK"},{"id":"aws-ap-northeast-1","label":"AWS - Tokyo","default":false,"type":"public","country":"JP"},{"id":"aws-ap-northeast-2","label":"AWS - Seoul","default":false,"type":"public","country":"KR"},{"id":"aws-ap-northeast-3","label":"AWS - Osaka","default":false,"type":"public","country":"JP"},{"id":"aws-ap-south-1","label":"AWS - Mumbai","default":false,"type":"public","country":"IN"},{"id":"aws-ap-southeast-1","label":"AWS - Singapore","default":false,"type":"public","country":"SG"},{"id":"aws-ap-southeast-2","label":"AWS - Sydney","default":false,"type":"public","country":"AU"},{"id":"aws-ap-southeast-3","label":"AWS - Jakarta","default":false,"type":"public","country":"ID"},{"id":"aws-ca-central-1","label":"AWS - Montreal","default":false,"type":"public","country":"CA"},{"id":"aws-eu-central-1","label":"AWS - Frankfurt","default":false,"type":"public","country":"DE"},{"id":"aws-eu-north-1","label":"AWS - Stockholm","default":false,"type":"public","country":"SE"},{"id":"aws-eu-south-1","label":"AWS - Milan","default":false,"type":"public","country":"IT"},{"id":"aws-eu-west-1","label":"AWS - Dublin","default":false,"type":"public","country":"IE"},{"id":"aws-eu-west-2","label":"AWS - London","default":false,"type":"public","country":"GB"},{"id":"aws-eu-west-3","label":"AWS - Paris","default":false,"type":"public","country":"FR"},{"id":"aws-me-south-1","label":"AWS - Bahrain","default":false,"type":"public","country":"BH"},{"id":"aws-sa-east-1","label":"AWS - São Paulo","default":false,"type":"public","country":"BR"},{"id":"aws-us-east-1","label":"AWS - N. Virginia","default":true,"type":"public","country":"US"},{"id":"aws-us-east-2","label":"AWS - Ohio","default":false,"type":"public","country":"US"},{"id":"aws-us-west-1","label":"AWS - N. California","default":false,"type":"public","country":"US"},{"id":"aws-us-west-2","label":"AWS - Oregon","default":false,"type":"public","country":"US"}],"defaultLocationIds":["aws-us-east-1"]}`
	inputGetLocationsV2 = verifyLocationsV2Input(getLocationsV2Body)
	getLocationV2Body   = `{"location":{"id":"aws-us-east-1","label":"AWS - N. Virginia","default":true,"type":"public","country":"US"},"meta":{"activeTestIds":[489,493,482,492,496,495,490,1856,2100,1568,1650,2559],"pausedTestIds":[]}}`
	inputGetLocationV2  = verifyLocationV2Input(getLocationV2Body)
)

func TestGetLocationsV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/locations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getLocationsV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetLocationsV2()

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.Location[1].ID, inputGetLocationsV2.Location[1].ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Location[1].ID, inputGetLocationsV2.Location[1].ID)
	}

	if !reflect.DeepEqual(resp.Location[1].Label, inputGetLocationsV2.Location[1].Label) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Location[1].Label, inputGetLocationsV2.Location[1].Label)
	}

	if !reflect.DeepEqual(resp.Location[1].Default, inputGetLocationsV2.Location[1].Default) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Location[1].Default, inputGetLocationsV2.Location[1].Default)
	}

	if !reflect.DeepEqual(resp.Location[1].Type, inputGetLocationsV2.Location[1].Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Location[1].Type, inputGetLocationsV2.Location[1].Type)
	}

	if !reflect.DeepEqual(resp.Location[1].Country, inputGetLocationsV2.Location[1].Country) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Location[1].Country, inputGetLocationsV2.Location[1].Country)
	}

	if !reflect.DeepEqual(resp.DefaultLocationIds, inputGetLocationsV2.DefaultLocationIds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DefaultLocationIds, inputGetLocationsV2.DefaultLocationIds)
	}

}

func verifyLocationsV2Input(stringInput string) *LocationsV2Response {
	check := &LocationsV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestGetLocationV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/locations/aws-us-east-1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getLocationV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, _, err := testClient.GetLocationV2("aws-us-east-1")

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.ID, inputGetLocationV2.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ID, inputGetLocationV2.ID)
	}

	if !reflect.DeepEqual(resp.Label, inputGetLocationV2.Label) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Label, inputGetLocationV2.Label)
	}

	if !reflect.DeepEqual(resp.Default, inputGetLocationV2.Default) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Default, inputGetLocationV2.Default)
	}

	if !reflect.DeepEqual(resp.Type, inputGetLocationV2.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Type, inputGetLocationV2.Type)
	}

	if !reflect.DeepEqual(resp.Country, inputGetLocationV2.Country) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Country, inputGetLocationV2.Country)
	}

	if !reflect.DeepEqual(resp.ActiveTestIds, inputGetLocationV2.ActiveTestIds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ActiveTestIds, inputGetLocationV2.ActiveTestIds)
	}

	if !reflect.DeepEqual(resp.PausedTestIds, inputGetLocationV2.PausedTestIds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.PausedTestIds, inputGetLocationV2.PausedTestIds)
	}

}

func verifyLocationV2Input(stringInput string) *LocationV2Response {
	check := &LocationV2Response{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}

func TestGetLocationsV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/locations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.GetLocationsV2()
	if err == nil {
		t.Fatal("expected a parse error, but got none")
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	if resp != nil {
		t.Errorf("expected nil response, got %#v", resp)
	}
}

func TestGetLocationsV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})
	_, _, err := unreachableClient.GetLocationsV2()
	if err == nil {
		t.Fatal("expected a connection error")
	}
}

func TestGetLocationV2ReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/locations/aws-us-east-1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.GetLocationV2("aws-us-east-1")
	if err == nil {
		t.Fatal("expected a parse error, but got none")
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	if resp != nil {
		t.Errorf("expected nil response, got %#v", resp)
	}
}

func TestGetLocationV2ReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})
	_, _, err := unreachableClient.GetLocationV2("aws-us-east-1")
	if err == nil {
		t.Fatal("expected a connection error")
	}
}
