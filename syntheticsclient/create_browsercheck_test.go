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
	createBrowserRespBody = `{"id":10,"name":"testtaste","type":"real_browser","frequency":10,"paused":false,"muted":false,"created_at":"2021-06-15T18:48:23.425Z","updated_at":"2021-06-15T18:48:23.425Z","links":{"self":"https://monitoring-api.rigor.com/v2/checks/191689","self_html":"https://monitoring.rigor.com/checks/real-browsers/191689","metrics":"https://monitoring-api.rigor.com/v2/checks/191689/metrics","last_run":null,"runs":"https://monitoring-api.rigor.com/v2/checks/real_browsers/191689/runs","share_link":"https://monitoring.rigor.com/share/350c659ff38a4781d0c9355252326a1beb8bac01ba4154a98267ac807daf7bbb*MTk2NTsyOzE5MTY4OQ=="},"status":{"last_code":200,"last_message":"","last_response_time":1939,"last_run_at":"2021-06-15T19:06:02.000Z","last_failure_at":null,"last_alert_at":null,"has_failure":false,"has_location_failure":null},"notifications":{"sms":false,"call":false,"email":true,"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false,"notify_who":[{"sms":false,"call":false,"email":true,"custom_user_email":null,"type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/15649"},"id":15649}],"notification_windows":[],"escalations":[]},"response_time_monitor_milliseconds":null,"http_request_headers":{"User-Agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"},"round_robin":true,"auto_retry":false,"enabled":true,"integrations":[],"url":"https://www.google.com/","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36","auto_update_user_agent":true,"viewport":{"width":1366,"height":768},"enforce_ssl_validation":true,"browser":{"label":"Google Chrome","code":"chrome"},"dns_overrides":{},"wait_for_full_metrics":true,"tags":[],"blackout_periods":[],"locations":[{"id":6,"world_region":"NA","region_code":"na-us-virginia","name":"N. Virginia, United States"}],"steps":[],"javascript_files":[],"threshold_monitors":[],"excluded_files":[],"cookies":[],"connection":{"download_bandwidth":20000,"upload_bandwidth":5000,"latency":28,"packet_loss":0}}`
	inputBrowserCreate    = verifyInput(string(createBrowserRespBody))
)

func TestCreateBrowseCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/real_browsers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := w.Write([]byte(createBrowserRespBody))
		if err != nil {
			t.Errorf("returned error: %#v", err)
		}
	})

	resp, _, err := testClient.CreateBrowserCheck(&BrowserCheckInput{
		Name:       "testtaste",
		Type:       "real_browser",
		Frequency:  10,
		Paused:     false,
		Muted:      false,
		RoundRobin: true,
		AutoRetry:  false,
		Enabled:    true,
		Locations:  []int{6},
		HTTPRequestHeaders: HTTPRequestHeaders{
			UserAgent: "Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36",
		},
		Notifications:       Notifications{},
		URL:                 "https://www.google.com/",
		UserAgent:           "",
		AutoUpdateUserAgent: false,
		Browser: Browser{
			Label: "Google Chrome",
			Code:  "chrome",
		},
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
	})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp.ID, inputBrowserCreate.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ID, inputBrowserCreate.ID)
	}

	if !reflect.DeepEqual(resp.Name, inputBrowserCreate.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Name, inputBrowserCreate.Name)
	}

	if !reflect.DeepEqual(resp.Type, inputBrowserCreate.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Type, inputBrowserCreate.Type)
	}

	if !reflect.DeepEqual(resp.Frequency, inputBrowserCreate.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Frequency, inputBrowserCreate.Frequency)
	}

	if !reflect.DeepEqual(resp.Paused, inputBrowserCreate.Paused) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Paused, inputBrowserCreate.Paused)
	}

	if !reflect.DeepEqual(resp.Muted, inputBrowserCreate.Muted) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Muted, inputBrowserCreate.Muted)
	}

	if !reflect.DeepEqual(resp.CreatedAt, inputBrowserCreate.CreatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.CreatedAt, inputBrowserCreate.CreatedAt)
	}

	if !reflect.DeepEqual(resp.UpdatedAt, inputBrowserCreate.UpdatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.UpdatedAt, inputBrowserCreate.UpdatedAt)
	}

	if !reflect.DeepEqual(resp.Links, inputBrowserCreate.Links) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Links, inputBrowserCreate.Links)
	}

	if !reflect.DeepEqual(resp.Status, inputBrowserCreate.Status) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Status, inputBrowserCreate.Status)
	}

	if !reflect.DeepEqual(resp.Notifications, inputBrowserCreate.Notifications) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Notifications, inputBrowserCreate.Notifications)
	}

	if !reflect.DeepEqual(resp.HTTPRequestHeaders, inputBrowserCreate.HTTPRequestHeaders) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPRequestHeaders, inputBrowserCreate.HTTPRequestHeaders)
	}

	if !reflect.DeepEqual(resp.HTTPRequestBody, inputBrowserCreate.HTTPRequestBody) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPRequestBody, inputBrowserCreate.HTTPRequestBody)
	}

	if !reflect.DeepEqual(resp.HTTPMethod, inputBrowserCreate.HTTPMethod) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPMethod, inputBrowserCreate.HTTPMethod)
	}

	if !reflect.DeepEqual(resp.RoundRobin, inputBrowserCreate.RoundRobin) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.RoundRobin, inputBrowserCreate.RoundRobin)
	}

	if !reflect.DeepEqual(resp.AutoRetry, inputBrowserCreate.AutoRetry) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.AutoRetry, inputBrowserCreate.AutoRetry)
	}

	if !reflect.DeepEqual(resp.Enabled, inputBrowserCreate.Enabled) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Enabled, inputBrowserCreate.Enabled)
	}

	if !reflect.DeepEqual(resp.Integrations, inputBrowserCreate.Integrations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Integrations, inputBrowserCreate.Integrations)
	}

	if !reflect.DeepEqual(resp.URL, inputBrowserCreate.URL) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.URL, inputBrowserCreate.URL)
	}

	if !reflect.DeepEqual(resp.UserAgent, inputBrowserCreate.UserAgent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.UserAgent, inputBrowserCreate.UserAgent)
	}

	if !reflect.DeepEqual(resp.AutoUpdateUserAgent, inputBrowserCreate.AutoUpdateUserAgent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.AutoUpdateUserAgent, inputBrowserCreate.AutoUpdateUserAgent)
	}

	if !reflect.DeepEqual(resp.Viewport, inputBrowserCreate.Viewport) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Viewport, inputBrowserCreate.Viewport)
	}

	if !reflect.DeepEqual(resp.EnforceSslValidation, inputBrowserCreate.EnforceSslValidation) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.EnforceSslValidation, inputBrowserCreate.EnforceSslValidation)
	}

	if !reflect.DeepEqual(resp.Browser, inputBrowserCreate.Browser) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Browser, inputBrowserCreate.Browser)
	}

	if !reflect.DeepEqual(resp.DNSOverrides, inputBrowserCreate.DNSOverrides) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DNSOverrides, inputBrowserCreate.DNSOverrides)
	}

	if !reflect.DeepEqual(resp.WaitForFullMetrics, inputBrowserCreate.WaitForFullMetrics) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.WaitForFullMetrics, inputBrowserCreate.WaitForFullMetrics)
	}

	if !reflect.DeepEqual(resp.Tags, inputBrowserCreate.Tags) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Tags, inputBrowserCreate.Tags)
	}

	if !reflect.DeepEqual(resp.BlackoutPeriods, inputBrowserCreate.BlackoutPeriods) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.BlackoutPeriods, inputBrowserCreate.BlackoutPeriods)
	}

	if !reflect.DeepEqual(resp.Locations, inputBrowserCreate.Locations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Locations, inputBrowserCreate.Locations)
	}

	if !reflect.DeepEqual(resp.Steps, inputBrowserCreate.Steps) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Steps, inputBrowserCreate.Steps)
	}

	if !reflect.DeepEqual(resp.JavascriptFiles, inputBrowserCreate.JavascriptFiles) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.JavascriptFiles, inputBrowserCreate.JavascriptFiles)
	}

	if !reflect.DeepEqual(resp.ThresholdMonitors, inputBrowserCreate.ThresholdMonitors) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ThresholdMonitors, inputBrowserCreate.ThresholdMonitors)
	}

	if !reflect.DeepEqual(resp.ExcludedFiles, inputBrowserCreate.ExcludedFiles) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ExcludedFiles, inputBrowserCreate.ExcludedFiles)
	}

	if !reflect.DeepEqual(resp.Cookies, inputBrowserCreate.Cookies) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Cookies, inputBrowserCreate.Cookies)
	}

	if !reflect.DeepEqual(resp.Connection, inputBrowserCreate.Connection) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Connection, inputBrowserCreate.Connection)
	}

}
