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
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

var (
	getBrowserBody  = `{"ID":206537,"name":"TestTheDepths","type":"real_browser","frequency":30,"created_at":"2021-07-13T17:07:50.000Z","updated_at":"2021-07-13T17:07:50.000Z","links":{"self":"https://monitoring-api.rigor.com/v2/checks/206537","self_html":"https://monitoring.rigor.com/checks/real-browsers/206537","metrics":"https://monitoring-api.rigor.com/v2/checks/206537/metrics"},"status":{},"notifications":{"email":true,"notify_who":[{"email":true,"custom_email":"","type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/18100"},"id":18100}],"notify_after_failure_count":2,"notify_on_location_failure":true,"notification_windows":[{"start_time":"2000-01-01T09:30:00.000Z","end_time":"2000-01-01T11:30:00.000Z","duration_in_minutes":120,"time_zone":"Eastern Time (US \u0026 Canada)"}],"escalations":[{"sms":true,"email":true,"after_minutes":5,"notify_who":[{"custom_email":"","type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/18100"},"id":18100}],"notification_window":{}}]},"response_time_monitor_milliseconds":9999,"http_request_headers":{"User-Agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"},"round_robin":true,"auto_retry":true,"enabled":true,"integrations":[{"id":71,"name":"Send data to SIM"}],"url":"https://www.google.com","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36","auto_update_user_agent":true,"viewport":{"height":768,"width":1366},"enforce_ssl_validation":true,"browser":{"label":"Google Chrome","code":"chrome"},"dns_overrides":{},"wait_for_full_metrics":true,"tags":[{"id":8449,"name":"myTag"}],"blackout_periods":[{"start_date":"2021-07-20","end_date":"2021-07-20","timezone":"Eastern Time (US \u0026 Canada)","start_time":"2000-01-01T17:00:00.000Z","end_time":"2000-01-01T17:30:00.000Z","duration_in_minutes":30,"created_at":"2021-07-13T17:07:51.000Z","updated_at":"2021-07-13T17:07:51.000Z"}],"locations":[{"id":6,"name":"N. Virginia, United States","world_region":"NA","region_code":"na-us-virginia"}],"steps":[{"item_method":"click_element","how":"id","what":"clickything","updated_at":"2021-07-13T17:07:51.000Z","created_at":"2021-07-13T17:07:51.000Z","name":"click"},{"item_method":"click_element","how":"id","what":"secondclickything","updated_at":"2021-07-13T17:07:51.000Z","created_at":"2021-07-13T17:07:51.000Z","name":"click again","position":1}],"javascript_files":[{"id":322,"name":"helloworld","created_at":"2021-07-08T21:58:54.000Z","updated_at":"2021-07-08T21:58:54.000Z","links":{"self":"https://production-javascript-files.s3.amazonaws.com/account-1965/9c617000c26501391a970242ac110004/helloworld.js?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=AKIAJZXH5XPVBC6HGICA%2F20210713%2Fus-east-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20210713T170810Z\u0026X-Amz-Expires=1800\u0026X-Amz-SignedHeaders=host\u0026X-Amz-Signature=a8d82709f775b9acfac8bf8a4d73984a9663d3d7b0f9ffe0261a07de0ada3c32"}}],"threshold_monitors":[{"matcher":"www.google.com","metric_name":"first_byte_time_ms","comparison_type":"less_than","value":9999,"created_at":"2021-07-13T17:07:51.000Z","updated_at":"2021-07-13T17:07:51.000Z"}],"excluded_files":[{"exclusion_type":"preset","preset_name":"chartbeat","url":"static\\.chartbeat\\.com","created_at":"2021-07-13T17:07:51.000Z","updated_at":"2021-07-13T17:07:51.000Z"}],"cookies":[{"key":"cookie","value":"startswithc","domain":"nodomain.com","path":"/path/to/path"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`
	inputGetBrowser = verifyInput(string(getBrowserBody))
)

func TestGetBrowserCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/206537", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte(getBrowserBody))
		if err != nil {
			t.Errorf("returned error: %#v", err)
		}
	})

	resp, _, err := testClient.GetCheck(206537)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(resp.ID, inputGetBrowser.ID) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ID, inputGetBrowser.ID)
	}

	if !reflect.DeepEqual(resp.Name, inputGetBrowser.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Name, inputGetBrowser.Name)
	}

	if !reflect.DeepEqual(resp.Type, inputGetBrowser.Type) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Type, inputGetBrowser.Type)
	}

	if !reflect.DeepEqual(resp.Frequency, inputGetBrowser.Frequency) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Frequency, inputGetBrowser.Frequency)
	}

	if !reflect.DeepEqual(resp.Paused, inputGetBrowser.Paused) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Paused, inputGetBrowser.Paused)
	}

	if !reflect.DeepEqual(resp.Muted, inputGetBrowser.Muted) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Muted, inputGetBrowser.Muted)
	}

	if !reflect.DeepEqual(resp.CreatedAt, inputGetBrowser.CreatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.CreatedAt, inputGetBrowser.CreatedAt)
	}

	if !reflect.DeepEqual(resp.UpdatedAt, inputGetBrowser.UpdatedAt) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.UpdatedAt, inputGetBrowser.UpdatedAt)
	}

	if !reflect.DeepEqual(resp.Links, inputGetBrowser.Links) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Links, inputGetBrowser.Links)
	}

	if !reflect.DeepEqual(resp.Status, inputGetBrowser.Status) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Status, inputGetBrowser.Status)
	}

	if !reflect.DeepEqual(resp.Notifications, inputGetBrowser.Notifications) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Notifications, inputGetBrowser.Notifications)
	}

	if !reflect.DeepEqual(resp.ResponseTimeMonitorMilliseconds, inputGetBrowser.ResponseTimeMonitorMilliseconds) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ResponseTimeMonitorMilliseconds, inputGetBrowser.ResponseTimeMonitorMilliseconds)
	}

	if !reflect.DeepEqual(resp.HTTPRequestHeaders, inputGetBrowser.HTTPRequestHeaders) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPRequestHeaders, inputGetBrowser.HTTPRequestHeaders)
	}

	if !reflect.DeepEqual(resp.HTTPRequestBody, inputGetBrowser.HTTPRequestBody) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPRequestBody, inputGetBrowser.HTTPRequestBody)
	}

	if !reflect.DeepEqual(resp.HTTPMethod, inputGetBrowser.HTTPMethod) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.HTTPMethod, inputGetBrowser.HTTPMethod)
	}

	if !reflect.DeepEqual(resp.RoundRobin, inputGetBrowser.RoundRobin) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.RoundRobin, inputGetBrowser.RoundRobin)
	}

	if !reflect.DeepEqual(resp.AutoRetry, inputGetBrowser.AutoRetry) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.AutoRetry, inputGetBrowser.AutoRetry)
	}

	if !reflect.DeepEqual(resp.Enabled, inputGetBrowser.Enabled) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Enabled, inputGetBrowser.Enabled)
	}

	if !reflect.DeepEqual(resp.Integrations, inputGetBrowser.Integrations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Integrations, inputGetBrowser.Integrations)
	}

	if !reflect.DeepEqual(resp.URL, inputGetBrowser.URL) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.URL, inputGetBrowser.URL)
	}

	if !reflect.DeepEqual(resp.UserAgent, inputGetBrowser.UserAgent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.UserAgent, inputGetBrowser.UserAgent)
	}

	if !reflect.DeepEqual(resp.AutoUpdateUserAgent, inputGetBrowser.AutoUpdateUserAgent) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.AutoUpdateUserAgent, inputGetBrowser.AutoUpdateUserAgent)
	}

	if !reflect.DeepEqual(resp.Viewport, inputGetBrowser.Viewport) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Viewport, inputGetBrowser.Viewport)
	}

	if !reflect.DeepEqual(resp.EnforceSslValidation, inputGetBrowser.EnforceSslValidation) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.EnforceSslValidation, inputGetBrowser.EnforceSslValidation)
	}

	if !reflect.DeepEqual(resp.Browser, inputGetBrowser.Browser) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Browser, inputGetBrowser.Browser)
	}

	if !reflect.DeepEqual(resp.DNSOverrides, inputGetBrowser.DNSOverrides) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.DNSOverrides, inputGetBrowser.DNSOverrides)
	}

	if !reflect.DeepEqual(resp.WaitForFullMetrics, inputGetBrowser.WaitForFullMetrics) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.WaitForFullMetrics, inputGetBrowser.WaitForFullMetrics)
	}

	if !reflect.DeepEqual(resp.Tags, inputGetBrowser.Tags) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Tags, inputGetBrowser.Tags)
	}

	if !reflect.DeepEqual(resp.BlackoutPeriods, inputGetBrowser.BlackoutPeriods) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.BlackoutPeriods, inputGetBrowser.BlackoutPeriods)
	}

	if !reflect.DeepEqual(resp.Locations, inputGetBrowser.Locations) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Locations, inputGetBrowser.Locations)
	}

	if !reflect.DeepEqual(resp.Steps, inputGetBrowser.Steps) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Steps, inputGetBrowser.Steps)
	}

	if !reflect.DeepEqual(resp.JavascriptFiles, inputGetBrowser.JavascriptFiles) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.JavascriptFiles, inputGetBrowser.JavascriptFiles)
	}

	if !reflect.DeepEqual(resp.ThresholdMonitors, inputGetBrowser.ThresholdMonitors) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ThresholdMonitors, inputGetBrowser.ThresholdMonitors)
	}

	if !reflect.DeepEqual(resp.ExcludedFiles, inputGetBrowser.ExcludedFiles) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.ExcludedFiles, inputGetBrowser.ExcludedFiles)
	}

	if !reflect.DeepEqual(resp.Cookies, inputGetBrowser.Cookies) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Cookies, inputGetBrowser.Cookies)
	}

	if !reflect.DeepEqual(resp.Connection, inputGetBrowser.Connection) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Connection, inputGetBrowser.Connection)
	}

	if !reflect.DeepEqual(resp.SuccessCriteria, inputGetBrowser.SuccessCriteria) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.SuccessCriteria, inputGetBrowser.SuccessCriteria)
	}

}

func verifyInput(stringInput string) *GetCheck {
	check := &GetCheck{}
	err := json.Unmarshal([]byte(stringInput), check)
	if err != nil {
		panic(err)
	}
	return check
}
