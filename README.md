# syntheticsclient
A Splunk Synthetics (Formerly Rigor) client for golang.

## Installation
`go get https://github.com/greatestusername-splunk/syntheticsclient.git`

## Important Note

This client is used to make the API calls that are mentioned here [Splunk Synthetics (Formerly Rigor) public API](https://monitoring-api.rigor.com/). However, some features are not implemented or publicly available yet. 

## Example Usages
```go
package syntheticsclient

import (
	"fmt"

	"https://github.com/greatestusername-splunk/syntheticsclient.git"
)

func main() {
	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	token := os.Getenv("API_ACCESS_TOKEN")

	//Create your client with the token
	c := NewClient(token)

	// Initialize your check settings
	o := CreateHttpCheck{
	Name:               "test test",
	Frequency:          5,
	URL:                "https://www.google.com"}

	// Make the request with your check settings and print result
  res, _, err := c.CreateHttpCheck(&o)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}
}
```

Another possibility for initializing your interface is to unmarshal a valid JSON string into the needed struct
```go
package syntheticsclient

import (
	"fmt"
	"encoding/json"

	"https://github.com/greatestusername-splunk/syntheticsclient.git"
)

func main() {

	token := os.Getenv("API_ACCESS_TOKEN")

	c := NewClient(token)

	//Take your ugly (but valid) JSON string as bytes and unmarshal into a CreateBrowserCheck struct
	jsonData := []byte(`{"name":"test","type":"real_browser","frequency":10,"round_robin":true,"auto_retry":false,"enabled":true,"integrations":[],"http_request_headers":{"User-Agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"},"notifications":{"sms":false,"call":false,"email":true,"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false,"notify_who":[{"sms":false,"call":false,"email":true,"type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/1"},"id":1}],"notification_windows":[],"escalations":[]},"url":"https://www.google.com/","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36","auto_update_user_agent":true,"viewport":{"width":1366,"height":768},"enforce_ssl_validation":true,"browser":{"type":"chrome"},"dns_overrides":{},"wait_for_full_metrics":true,"tags":[],"blackout_periods":[{"timezone":"Eastern Time (US & Canada)","start_time":"2000-01-01T07:00:00.000Z","end_time":"2000-01-01T14:00:00.000Z","repeat_type":"daily","duration_in_minutes":420,"is_repeat":true,"created_at":"2020-11-11T21:54:32.000Z","updated_at":"2021-03-30T19:02:54.000Z"}],"steps":[],"javascript_files":[],"threshold_monitors":[],"excluded_files":[],"cookies":[],"connection":{"download_bandwidth":20000,"upload_bandwidth":5000,"latency":28,"packet_loss":0}}`)
	var browserCheckDetail CreateBrowserCheck
	err := json.Unmarshal(jsonData, &browserCheckDetail)
	if err != nil {
		fmt.Println(err)
	}

	//Use your converted JSON to make the request and print
	res, _, err := c.CreateBrowserCheck(&browserCheckDetail)
	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}
}
```

## Additional Information
This client is largely a copypasta mutation of the [go-victor](https://github.com/victorops/go-victorops) client for Splunk On-Call (formerly known as VictorOps).