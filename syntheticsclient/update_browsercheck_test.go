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
	updateBrowserCheckBody = &BrowserCheckInput{
		Name:                "testingcase",
		Type:                "real_browser",
		Frequency:           10,
		Paused:              false,
		Muted:               false,
		RoundRobin:          false,
		AutoRetry:           false,
		Enabled:             true,
		Locations:           []int{6},
		URL:                 "https://www.google.com/",
		UserAgent:           "",
		AutoUpdateUserAgent: false,
		Viewport: Viewport{
			Height: 1366,
			Width:  768,
		},
		EnforceSslValidation: true,
		Connection: Connection{
			UploadBandwidth:   20000,
			DownloadBandwidth: 5000,
			Latency:           28,
			PacketLoss:        0,
		},
		WaitForFullMetrics: true,
	}

	updateBrowserCheckResponse = `{"id":10,"name":"testingcase","type":"real_browser","frequency":10,"paused":false,"muted":false,"created_at":"2021-06-16T17:44:52.000Z","updated_at":"2021-06-16T17:45:34.958Z","links":{"self":"https://monitoring-api.rigor.com/v2/checks/191819","self_html":"https://monitoring.rigor.com/checks/real-browsers/191819","metrics":"https://monitoring-api.rigor.com/v2/checks/191819/metrics","last_run":null,"runs":"https://monitoring-api.rigor.com/v2/checks/real_browsers/191819/runs","share_link":"https://monitoring.rigor.com/share/6bc22dea2eb662aba42d48fa40699c796b46b3bc923f0ff30cfd0f907b77f322*MTk2NTsyOzE5MTgxOQ=="},"status":{"last_code":null,"last_message":null,"last_response_time":null,"last_run_at":null,"last_failure_at":null,"last_alert_at":null,"has_failure":null,"has_location_failure":null},"notifications":{"sms":false,"call":false,"email":true,"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false,"notify_who":[{"sms":false,"call":false,"email":true,"custom_user_email":null,"type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/18100"},"id":18100}],"notification_windows":[],"escalations":[]},"response_time_monitor_milliseconds":null,"http_request_headers":{"User-Agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"},"round_robin":true,"auto_retry":true,"enabled":true,"integrations":[],"url":"https://www.google.com/","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36","auto_update_user_agent":true,"viewport":{"width":1366,"height":768},"enforce_ssl_validation":true,"browser":{"label":"Google Chrome","code":"chrome"},"dns_overrides":{},"wait_for_full_metrics":true,"tags":[],"blackout_periods":[],"locations":[{"id":6,"world_region":"NA","region_code":"na-us-virginia","name":"N. Virginia, United States"}],"steps":[],"javascript_files":[],"threshold_monitors":[],"excluded_files":[],"cookies":[],"connection":{"download_bandwidth":20000,"upload_bandwidth":5000,"latency":28,"packet_loss":0}}`
)

func TestUpdateBrowserCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/real_browsers/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateBrowserCheckResponse))
		if err != nil {
			t.Errorf("returned error: %#v", err)
		}
	})

	resp, _, err := testClient.UpdateBrowserCheck(10, updateBrowserCheckBody)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Name != "testingcase" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Name, "testingcase")
	}
}
