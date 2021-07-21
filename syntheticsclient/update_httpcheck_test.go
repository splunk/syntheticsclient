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
	"testing"
)

var (
	updateBody = &HttpCheckInput{Name: "http test check"}

	updateResponse = `{"id":19,"name":"http test check","type":"http","frequency":5,"paused":false,"muted":false,"created_at":"2021-06-11T15:29:01Z","updated_at":"2021-06-11T18:19:05Z","links":{"self":"https://monitoring-api.rigor.com/v2/checks/191484","self_html":"https://monitoring.rigor.com/checks/http/191484","metrics":"https://monitoring-api.rigor.com/v2/checks/191484/metrics","last_run":"https://monitoring.rigor.com/checks/191484/runs/9338120804"},"tags":[{"id":9,"name":"test"}],"status":{"last_code":200,"last_message":"OK","last_response_time":295,"last_run_at":"2021-06-11T18:19:04Z","last_failure_at":"0001-01-01T00:00:00Z","last_alert_at":"0001-01-01T00:00:00Z","has_failure":false,"has_location_failure":false},"round_robin":true,"auto_retry":false,"enabled":true,"blackout_periods":[],"locations":[{"id":1,"name":"N. Virginia, United States","world_region":"NA","region_code":"na-us-virginia"}],"integrations":[],"http_request_headers":{"User-Agent":"Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"},"notifications":{"sms":false,"email":true,"call":false,"notify_who":[{"sms":false,"email":true,"call":false,"links":{}}],"notify_after_failure_count":2,"notify_on_location_failure":true,"notification_windows":[],"escalations":[],"muted":false},"url":"https://www.google.com","http_method":"GET","success_criteria":[{"action_type":"presence_of_text","comparison_string":"About","created_at":"2021-06-11T18:19:07.81Z","updated_at":"2021-06-11T18:19:07.81Z"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`
)

func TestUpdateHttpCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/http/19", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.Write([]byte(updateResponse))
	})

	resp, _, err := testClient.UpdateHttpCheck(19, updateBody)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Name != "http test check" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Name, "http test check")
	}
}
