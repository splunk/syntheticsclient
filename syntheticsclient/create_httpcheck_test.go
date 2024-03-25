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

package syntheticsclient

import (
	"net/http"
	"reflect"
	"testing"
)

var (
	createHttpBody  = `{"id":19,"name":"New Homepage","type":"http","frequency":5,"paused":false,"muted":false,"created_at":"2021-06-14T14:12:31.494Z","updated_at":"2021-06-14T14:12:31.494Z","links":{"self":"https://monitoring-api.rigor.com/v2/checks/191588","self_html":"https://monitoring.rigor.com/checks/http/191588","metrics":"https://monitoring-api.rigor.com/v2/checks/191588/metrics","last_run":""},"tags":[{"id":9,"name":"test"}],"status":{"last_code":0,"last_message":"","last_response_time":0,"last_run_at":"0001-01-01T00:00:00Z","last_failure_at":"0001-01-01T00:00:00Z","last_alert_at":"0001-01-01T00:00:00Z","has_failure":false,"has_location_failure":false},"round_robin":true,"auto_retry":false,"enabled":true,"blackout_periods":[],"locations":[{"id":1,"name":"N. Virginia, United States","world_region":"NA","region_code":"na-us-virginia"}],"integrations":[],"http_request_headers":{"User-Agent":"Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"},"notifications":{"sms":false,"email":true,"call":false,"notify_who":[{"sms":false,"email":true,"call":false,"links":{}}],"notify_after_failure_count":2,"notify_on_location_failure":true,"notification_windows":[],"escalations":[],"muted":false},"url":"https://www.google.com","http_method":"GET","success_criteria":[{"action_type":"presence_of_text","comparison_string":"About","created_at":"2021-06-14T14:12:31.526Z","updated_at":"2021-06-14T14:12:31.526Z"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`
	inputCreateHttp = verifyInput(string(createHttpBody))
)

func TestCreateHttpCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/http", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(createHttpBody))
		if err != nil {
			t.Errorf("returned error: %#v", err)
		}
	})

	resp, _, err := testClient.CreateHttpCheck(&HttpCheckInput{
		Name:       "New Homepage",
		Type:       "http",
		Frequency:  5,
		Paused:     false,
		Muted:      false,
		Tags:       []string{"test"},
		RoundRobin: true,
		AutoRetry:  false,
		Enabled:    true,
		Locations:  []int{1},
		HTTPRequestHeaders: HTTPRequestHeaders{
			UserAgent: "Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)",
		},
		URL:             "https://www.google.com",
		HTTPMethod:      "GET",
		SuccessCriteria: []SuccessCriteria{},
		Connection: Connection{
			UploadBandwidth:   5000,
			DownloadBandwidth: 20000,
			Latency:           28,
			PacketLoss:        0,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.ID, inputCreateHttp.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ID, inputCreateHttp.ID)
	}

	if !reflect.DeepEqual(resp.Name, inputCreateHttp.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Name, inputCreateHttp.Name)
	}

	if !reflect.DeepEqual(resp.Type, inputCreateHttp.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Type, inputCreateHttp.Type)
	}

	if !reflect.DeepEqual(resp.Frequency, inputCreateHttp.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Frequency, inputCreateHttp.Frequency)
	}

	if !reflect.DeepEqual(resp.Paused, inputCreateHttp.Paused) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Paused, inputCreateHttp.Paused)
	}

	if !reflect.DeepEqual(resp.Muted, inputCreateHttp.Muted) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Muted, inputCreateHttp.Muted)
	}

	if !reflect.DeepEqual(resp.CreatedAt, inputCreateHttp.CreatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.CreatedAt, inputCreateHttp.CreatedAt)
	}

	if !reflect.DeepEqual(resp.UpdatedAt, inputCreateHttp.UpdatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.UpdatedAt, inputCreateHttp.UpdatedAt)
	}

	if !reflect.DeepEqual(resp.Links, inputCreateHttp.Links) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Links, inputCreateHttp.Links)
	}

	if !reflect.DeepEqual(resp.Status, inputCreateHttp.Status) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Status, inputCreateHttp.Status)
	}

	if !reflect.DeepEqual(resp.Notifications, inputCreateHttp.Notifications) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Notifications, inputCreateHttp.Notifications)
	}

	if !reflect.DeepEqual(resp.HTTPRequestHeaders, inputCreateHttp.HTTPRequestHeaders) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPRequestHeaders, inputCreateHttp.HTTPRequestHeaders)
	}

	if !reflect.DeepEqual(resp.HTTPRequestBody, inputCreateHttp.HTTPRequestBody) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPRequestBody, inputCreateHttp.HTTPRequestBody)
	}

	if !reflect.DeepEqual(resp.HTTPMethod, inputCreateHttp.HTTPMethod) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPMethod, inputCreateHttp.HTTPMethod)
	}

	if !reflect.DeepEqual(resp.RoundRobin, inputCreateHttp.RoundRobin) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.RoundRobin, inputCreateHttp.RoundRobin)
	}

	if !reflect.DeepEqual(resp.AutoRetry, inputCreateHttp.AutoRetry) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.AutoRetry, inputCreateHttp.AutoRetry)
	}

	if !reflect.DeepEqual(resp.Enabled, inputCreateHttp.Enabled) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Enabled, inputCreateHttp.Enabled)
	}

	if !reflect.DeepEqual(resp.Integrations, inputCreateHttp.Integrations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Integrations, inputCreateHttp.Integrations)
	}

	if !reflect.DeepEqual(resp.URL, inputCreateHttp.URL) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.URL, inputCreateHttp.URL)
	}

	if !reflect.DeepEqual(resp.Tags, inputCreateHttp.Tags) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Tags, inputCreateHttp.Tags)
	}

	if !reflect.DeepEqual(resp.BlackoutPeriods, inputCreateHttp.BlackoutPeriods) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.BlackoutPeriods, inputCreateHttp.BlackoutPeriods)
	}

	if !reflect.DeepEqual(resp.Locations, inputCreateHttp.Locations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Locations, inputCreateHttp.Locations)
	}

	if !reflect.DeepEqual(resp.Connection, inputCreateHttp.Connection) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Connection, inputCreateHttp.Connection)
	}

	if !reflect.DeepEqual(resp.SuccessCriteria, inputCreateHttp.SuccessCriteria) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.SuccessCriteria, inputCreateHttp.SuccessCriteria)
	}

}
