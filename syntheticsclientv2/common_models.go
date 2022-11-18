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
	"time"
)

// Common and shared struct models used for more complex requests
type Links struct {
	Self     string `json:"self,omitempty"`
	SelfHTML string `json:"self_html,omitempty"`
	Metrics  string `json:"metrics,omitempty"`
	LastRun  string `json:"last_run,omitempty"`
}

type Tags []struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Status struct {
	LastCode           int    `json:"last_code,omitempty"`
	LastMessage        string `json:"last_message,omitempty"`
	LastResponseTime   int    `json:"last_response_time,omitempty"`
	LastRunAt          string `json:"last_run_at,omitempty"`
	LastFailureAt      string `json:"last_failure_at,omitempty"`
	LastAlertAt        string `json:"last_alert_at,omitempty"`
	HasFailure         bool   `json:"has_failure,omitempty"`
	HasLocationFailure bool   `json:"has_location_failure,omitempty"`
}

type NotifyWho struct {
	Sms             bool   `json:"sms,omitempty"`
	Call            bool   `json:"call,omitempty"`
	Email           bool   `json:"email,omitempty"`
	CustomUserEmail string `json:"custom_email"`
	Type            string `json:"type,omitempty"`
	Links           Links  `json:"links,omitempty"`
	ID              int    `json:"id,omitempty"`
}

type NotificationWindows []struct {
	StartTime         string `json:"start_time,omitempty"`
	EndTime           string `json:"end_time,omitempty"`
	DurationInMinutes int    `json:"duration_in_minutes,omitempty"`
	TimeZone          string `json:"time_zone,omitempty"`
}

type NotificationWindow struct {
	StartTime         string `json:"start_time,omitempty"`
	EndTime           string `json:"end_time,omitempty"`
	DurationInMinutes int    `json:"duration_in_minutes,omitempty"`
	TimeZone          string `json:"time_zone,omitempty"`
}

type Escalations struct {
	Sms                bool               `json:"sms,omitempty"`
	Email              bool               `json:"email,omitempty"`
	Call               bool               `json:"call,omitempty"`
	AfterMinutes       int                `json:"after_minutes,omitempty"`
	NotifyWho          []NotifyWho        `json:"notify_who,omitempty"`
	IsRepeat           bool               `json:"is_repeat,omitempty"`
	NotificationWindow NotificationWindow `json:"notification_window,omitempty"`
}

type Notifications struct {
	Sms                     bool                `json:"sms,omitempty"`
	Email                   bool                `json:"email,omitempty"`
	Call                    bool                `json:"call,omitempty"`
	NotifyWho               []NotifyWho         `json:"notify_who,omitempty"`
	NotifyAfterFailureCount int                 `json:"notify_after_failure_count,omitempty"`
	NotifyOnLocationFailure bool                `json:"notify_on_location_failure,omitempty"`
	NotificationWindows     NotificationWindows `json:"notification_windows,omitempty"`
	Escalations             []Escalations       `json:"escalations,omitempty"`
	Muted                   bool                `json:"muted,omitempty"`
}

type SuccessCriteria struct {
	ActionType       string `json:"action_type,omitempty"`
	ComparisonString string `json:"comparison_string,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

type Connection struct {
	UploadBandwidth   int     `json:"upload_bandwidth,omitempty"`
	DownloadBandwidth int     `json:"download_bandwidth,omitempty"`
	Latency           int     `json:"latency,omitempty"`
	PacketLoss        float64 `json:"packet_loss,omitempty"`
}

type Locations []struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	WorldRegion string `json:"world_region,omitempty"`
	RegionCode  string `json:"region_code,omitempty"`
}

type Integrations []struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type HTTPRequestHeaders struct {
	UserAgent string `json:"User-Agent,omitempty"`
}

type Browser struct {
	Label string `json:"label,omitempty"`
	Code  string `json:"code,omitempty"`
}

type Steps struct {
	ItemMethod   string `json:"item_method,omitempty"`
	Value        string `json:"value,omitempty"`
	How          string `json:"how,omitempty"`
	What         string `json:"what,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	VariableName string `json:"variable_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Position     int    `json:"position,omitempty"`
}

type Cookies struct {
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
	Domain string `json:"domain,omitempty"`
	Path   string `json:"path,omitempty"`
}

type JavascriptFiles struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Links     Links  `json:"links,omitempty"`
}

type ExcludedFiles struct {
	ExclusionType string `json:"exclusion_type,omitempty"`
	PresetName    string `json:"preset_name,omitempty"`
	URL           string `json:"url,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

type BlackoutPeriods []struct {
	StartDate         string `json:"start_date,omitempty"`
	EndDate           string `json:"end_date,omitempty"`
	Timezone          string `json:"timezone,omitempty"`
	StartTime         string `json:"start_time,omitempty"`
	EndTime           string `json:"end_time,omitempty"`
	RepeatType        string `json:"repeat_type,omitempty"`
	DurationInMinutes int    `json:"duration_in_minutes,omitempty"`
	IsRepeat          bool   `json:"is_repeat,omitempty"`
	MonthlyRepeatType string `json:"monthly_repeat_type,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
}

type Viewport struct {
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}

type ThresholdMonitors struct {
	Matcher        string `json:"matcher,omitempty"`
	MetricName     string `json:"metric_name,omitempty"`
	ComparisonType string `json:"comparison_type,omitempty"`
	Value          int    `json:"value,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
}

type DNSOverrides struct {
	OriginalDomainCom string `json:"original.domain.com,omitempty"`
	OriginalHostCom   string `json:"original.host.com,omitempty"`
}

type Networkconnection struct {
	Description       string `json:"description,omitempty"`
	Downloadbandwidth int    `json:"downloadBandwidth,omitempty"`
	Latency           int    `json:"latency,omitempty"`
	Packetloss        int    `json:"packetLoss,omitempty"`
	Uploadbandwidth   int    `json:"uploadBandwidth,omitempty"`
}

type Advancedsettings struct {
	Authentication					`json:"authentication"`
	Cookiesv2  							`json:"cookies"`
	BrowserHeaders					`json:"headers,omitempty"`
	Verifycertificates bool `json:"verifyCertificates,omitempty"`
} 

type Authentication struct {
	Password string `json:"password"`
	Username string `json:"username"`
} 

type Cookiesv2 []struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
}

type BrowserHeaders []struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Domain string `json:"domain"`
}

type Transactions []struct {
	Name  string 	`json:"name"`
	Stepsv2				`json:"steps"`
}

type BusinessTransactions []struct {
	Name  											string `json:"name"`
	BusinessTransactionStepsV2 				 `json:"steps"`
}

type Stepsv2 []struct {
	Name         string `json:"name"`
	Selector     string `json:"selector,omitempty"`
	Selectortype string `json:"selectorType,omitempty"`
	Type         string `json:"type,omitempty"`
	Waitfornav   bool   `json:"waitForNav,omitempty"`
}

type BusinessTransactionStepsV2 []struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	URL          string `json:"url,omitempty"`
	Action       string `json:"action,omitempty"`
	WaitForNav   bool   `json:"wait_for_nav"`
	SelectorType string `json:"selector_type,omitempty"`
	Selector     string `json:"selector,omitempty"`
}

type Device struct {
	ID                int    `json:"id,omitempty"`
	Label             string `json:"label,omitempty"`
	Networkconnection `json:"networkConnection,omitempty"`
	Viewportheight 		int `json:"viewportHeight,omitempty"`
	Viewportwidth  		int `json:"viewportWidth,omitempty"`
}

type Requests []struct {
	Configuration `json:"configuration,omitempty"`
	Setup 				`json:"setup,omitempty"`
	Validations 	`json:"validations,omitempty"`
}

type Configuration struct {
	Body    			string `json:"body"`
	Headers 			`json:"headers,omitempty"`
	Name          string `json:"name,omitempty"`
	Requestmethod string `json:"requestMethod,omitempty"`
	URL           string `json:"url,omitempty"`
}

type Headers map[string]interface{}

type Setup []struct {
	Extractor string `json:"extractor,omitempty"`
	Name      string `json:"name,omitempty"`
	Source    string `json:"source,omitempty"`
	Type      string `json:"type,omitempty"`
	Variable  string `json:"variable,omitempty"`
}

type Validations []struct {
	Actual     string `json:"actual,omitempty"`
	Comparator string `json:"comparator,omitempty"`
	Expected   string `json:"expected,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
} 

type Checks []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Frequency int    `json:"frequency"`
	Paused    bool   `json:"paused"`
	Muted     bool   `json:"muted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Links     Links  `json:"links"`
	Status    Status `json:"status"`
	Tags      Tags   `json:"tags"`
}

type Tests        []struct {
	Active             bool      `json:"active"`
	Createdat          time.Time `json:"createdAt"`
	Frequency          int       `json:"frequency"`
	ID                 int       `json:"id"`
	Locationids        []string  `json:"locationIds"`
	Name               string    `json:"name"`
	Schedulingstrategy string    `json:"schedulingStrategy"`
	Type               string    `json:"type"`
	Updatedat          time.Time `json:"updatedAt"`
}

type GetChecksV2Options struct {
	TestType    	string 	`json:"testType"`
	PerPage 	int    	`json:"perPage"`
	Page    	int 	  `json:"page"`
	Search   	string  `json:"search"`
	OrderBy		string	`json:"orderBy"`
}

type Errors []struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type HttpHeaders []struct {
	Name       	string `json:"name,omitempty"`
	Value 			string `json:"value,omitempty"`
}

type Variable struct {
	Createdat   time.Time `json:"createdAt,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Secret      bool      `json:"secret"`
	Updatedat   time.Time `json:"updatedAt,omitempty"`
	Value       string    `json:"value"`
}

type DeleteCheck struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Errors  Errors `json:"errors"`
}

type VariableV2Response struct {
	Variable `json:"variable"`
}

type VariableV2Input struct {
	Variable `json:"variable"`
}

type ChecksV2Response struct {
	Nextpagelink int `json:"nextPageLink"`
	Perpage      int `json:"perPage"`
	Tests        `json:"tests"`
	Totalcount int `json:"totalCount"`
}

type HttpCheckV2Response struct {
	Test struct {
		ID                 int         `json:"id"`
		Name               string      `json:"name"`
		Active             bool        `json:"active"`
		Frequency          int         `json:"frequency"`
		SchedulingStrategy string      `json:"scheduling_strategy"`
		CreatedAt          time.Time   `json:"created_at,omitempty"`
		UpdatedAt          time.Time   `json:"updated_at,omitempty"`
		LocationIds        []string    `json:"location_ids"`
		Type               string      `json:"type"`
		URL                string      `json:"url"`
		RequestMethod      string      `json:"request_method"`
		Body               string 		`json:"body,omitempty"`
		HttpHeaders        						 `json:"headers,omitempty"`
} `json:"test"`
}

type HttpCheckV2Input struct {
	Test struct {
		Name               string      `json:"name"`
		Type               string      `json:"type"`
		URL                string      `json:"url"`
		LocationIds        []string    `json:"location_ids"`
		Frequency          int         `json:"frequency"`
		SchedulingStrategy string      `json:"scheduling_strategy"`
		Active             bool        `json:"active"`
		RequestMethod      string      `json:"request_method"`
		Body               string	 			`json:"body,omitempty"`
		HttpHeaders             			`json:"headers,omitempty"`
	} `json:"test"`
}

type ApiCheckV2Input struct {
	Test struct {
		Active      bool     			`json:"active"`
		Deviceid    int      			`json:"device_id"`
		Frequency   int      			`json:"frequency"`
		Locationids []string 			`json:"location_ids"`
		Name        string   			`json:"name"`
		Requests    							`json:"requests"`
		Schedulingstrategy string `json:"schedulingStrategy"`
	} `json:"test"`
}

type ApiCheckV2Response struct {
	Test struct {
		Active    bool      					`json:"active,omitempty"`
		Createdat time.Time 					`json:"createdAt,omitempty"`
		Device    										`json:"device,omitempty"`
		Frequency   int      					`json:"frequency,omitempty"`
		ID          int      					`json:"id,omitempty"`
		Locationids []string 					`json:"location_ids,omitempty"`
		Name        string   					`json:"name,omitempty"`
		Requests    									`json:"requests,omitempty"`
		Schedulingstrategy string    	`json:"schedulingStrategy,omitempty"`
		Type               string    	`json:"type,omitempty"`
		Updatedat          time.Time 	`json:"updatedAt,omitempty"`
	} 
}

type BrowserCheckV2Input struct {
	Test struct {
		Name                 	string 		`json:"name"`
		BusinessTransactions  					`json:"business_transactions"`
		Urlprotocol       	 	string   	`json:"urlProtocol"`
		Starturl           		string   	`json:"startUrl"`
		LocationIds        		[]string 	`json:"location_ids"`
		DeviceID           		int      	`json:"device_id"`
		Frequency          		int      	`json:"frequency"`
		Schedulingstrategy 		string   	`json:"scheduling_strategy"`
		Active             		bool     	`json:"active"`
		Advancedsettings   							`json:"advanced_settings"`
	} `json:"test"`
}

type BrowserCheckV2Response struct {
	Test struct {
		Active           		bool 			`json:"active"`
		Advancedsettings 							`json:"advanced_settings"`
		BusinessTransactions  				`json:"business_transactions"`
		Createdat 					time.Time `json:"createdAt"`
		Device												`json:"device"`
		Frequency          	int      	`json:"frequency"`
		ID                 	int      	`json:"id"`
		Locationids        	[]string 	`json:"location_ids"`
		Name               	string		`json:"name"`
		Schedulingstrategy	string		`json:"scheduling_strategy"`
		Transactions									`json:"transactions"`
		Type      					string    `json:"type"`
		Updatedat 					time.Time `json:"updatedAt"`
	} `json:"test"`
}

type HttpCheckInput struct {
	ID                 int                `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Type               string             `json:"type,omitempty"`
	Frequency          int                `json:"frequency,omitempty"`
	Paused             bool               `json:"paused,omitempty"`
	Muted              bool               `json:"muted,omitempty"`
	CreatedAt          string             `json:"created_at,omitempty"`
	UpdatedAt          string             `json:"updated_at,omitempty"`
	Links              Links              `json:"links,omitempty"`
	Tags               []string           `json:"tags"`
	Status             Status             `json:"status,omitempty"`
	RoundRobin         bool               `json:"round_robin,omitempty"`
	AutoRetry          bool               `json:"auto_retry,omitempty"`
	Enabled            bool               `json:"enabled,omitempty"`
	BlackoutPeriods    BlackoutPeriods    `json:"blackout_periods,omitempty"`
	Locations          []int              `json:"locations,omitempty"`
	Integrations       []int              `json:"integrations,omitempty"`
	HTTPRequestHeaders HTTPRequestHeaders `json:"http_request_headers,omitempty"`
	HTTPRequestBody    string             `json:"http_request_body,omitempty"`
	Notifications      Notifications      `json:"notifications,omitempty"`
	URL                string             `json:"url,omitempty"`
	HTTPMethod         string             `json:"http_method,omitempty"`
	SuccessCriteria    []SuccessCriteria  `json:"success_criteria,omitempty"`
	Connection         Connection         `json:"connection,omitempty"`
}

type HttpCheckResponse struct {
	ID                 int                `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	Type               string             `json:"type,omitempty"`
	Frequency          int                `json:"frequency,omitempty"`
	Paused             bool               `json:"paused,omitempty"`
	Muted              bool               `json:"muted,omitempty"`
	CreatedAt          string             `json:"created_at,omitempty"`
	UpdatedAt          string             `json:"updated_at,omitempty"`
	Links              Links              `json:"links,omitempty"`
	Tags               Tags               `json:"tags,omitempty"`
	Status             Status             `json:"status,omitempty"`
	RoundRobin         bool               `json:"round_robin,omitempty"`
	AutoRetry          bool               `json:"auto_retry,omitempty"`
	Enabled            bool               `json:"enabled,omitempty"`
	BlackoutPeriods    BlackoutPeriods    `json:"blackout_periods,omitempty"`
	Locations          Locations          `json:"locations,omitempty"`
	Integrations       Integrations       `json:"integrations,omitempty"`
	HTTPRequestHeaders HTTPRequestHeaders `json:"http_request_headers,omitempty"`
	HTTPRequestBody    string             `json:"http_request_body,omitempty"`
	Notifications      Notifications      `json:"notifications,omitempty"`
	URL                string             `json:"url,omitempty"`
	HTTPMethod         string             `json:"http_method,omitempty"`
	SuccessCriteria    []SuccessCriteria  `json:"success_criteria,omitempty"`
	Connection         Connection         `json:"connection,omitempty"`
}

type BrowserCheckInput struct {
	ID                   int                 `json:"id,omitempty"`
	Name                 string              `json:"name,omitempty"`
	Type                 string              `json:"type,omitempty"`
	Frequency            int                 `json:"frequency,omitempty"`
	Paused               bool                `json:"paused,omitempty"`
	Muted                bool                `json:"muted,omitempty"`
	CreatedAt            string              `json:"created_at,omitempty"`
	UpdatedAt            string              `json:"updated_at,omitempty"`
	Links                Links               `json:"links,omitempty"`
	Tags                 []string            `json:"tags"`
	Status               Status              `json:"status,omitempty"`
	RoundRobin           bool                `json:"round_robin,omitempty"`
	AutoRetry            bool                `json:"auto_retry,omitempty"`
	Enabled              bool                `json:"enabled,omitempty"`
	BlackoutPeriods      BlackoutPeriods     `json:"blackout_periods,omitempty"`
	Locations            []int               `json:"locations,omitempty"`
	Integrations         []int               `json:"integrations,omitempty"`
	HTTPRequestHeaders   HTTPRequestHeaders  `json:"http_request_headers,omitempty"`
	HTTPRequestBody      string              `json:"http_request_body,omitempty"`
	HTTPMethod           string              `json:"http_method,omitempty"`
	Notifications        Notifications       `json:"notifications,omitempty"`
	URL                  string              `json:"url,omitempty"`
	UserAgent            string              `json:"user_agent,omitempty"`
	AutoUpdateUserAgent  bool                `json:"auto_update_user_agent,omitempty"`
	Browser              Browser             `json:"browser,omitempty"`
	Steps                []Steps             `json:"steps,omitempty"`
	Cookies              []Cookies           `json:"cookies,omitempty"`
	JavascriptFiles      []JavascriptFiles   `json:"javascript_files,omitempty"`
	ExcludedFiles        []ExcludedFiles     `json:"excluded_files,omitempty"`
	Viewport             Viewport            `json:"viewport,omitempty"`
	EnforceSslValidation bool                `json:"enforce_ssl_validation,omitempty"`
	ThresholdMonitors    []ThresholdMonitors `json:"threshold_monitors,omitempty"`
	DNSOverrides         DNSOverrides        `json:"dns_overrides,omitempty"`
	Connection           Connection          `json:"connection,omitempty"`
	WaitForFullMetrics   bool                `json:"wait_for_full_metrics,omitempty"`
}

type BrowserCheckResponse struct {
	ID                   int                 `json:"id,omitempty"`
	Name                 string              `json:"name,omitempty"`
	Type                 string              `json:"type,omitempty"`
	Frequency            int                 `json:"frequency,omitempty"`
	Paused               bool                `json:"paused,omitempty"`
	Muted                bool                `json:"muted,omitempty"`
	CreatedAt            string              `json:"created_at,omitempty"`
	UpdatedAt            string              `json:"updated_at,omitempty"`
	Links                Links               `json:"links,omitempty"`
	Tags                 Tags                `json:"tags,omitempty"`
	Status               Status              `json:"status,omitempty"`
	RoundRobin           bool                `json:"round_robin,omitempty"`
	AutoRetry            bool                `json:"auto_retry,omitempty"`
	Enabled              bool                `json:"enabled,omitempty"`
	BlackoutPeriods      BlackoutPeriods     `json:"blackout_periods,omitempty"`
	Locations            Locations           `json:"locations,omitempty"`
	Integrations         Integrations        `json:"integrations,omitempty"`
	HTTPRequestHeaders   HTTPRequestHeaders  `json:"http_request_headers,omitempty"`
	HTTPRequestBody      string              `json:"http_request_body,omitempty"`
	HTTPMethod           string              `json:"http_method,omitempty"`
	Notifications        Notifications       `json:"notifications,omitempty"`
	URL                  string              `json:"url,omitempty"`
	UserAgent            string              `json:"user_agent,omitempty"`
	AutoUpdateUserAgent  bool                `json:"auto_update_user_agent,omitempty"`
	Browser              Browser             `json:"browser,omitempty"`
	Steps                []Steps             `json:"steps,omitempty"`
	Cookies              []Cookies           `json:"cookies,omitempty"`
	JavascriptFiles      []JavascriptFiles   `json:"javascript_files,omitempty"`
	ExcludedFiles        []ExcludedFiles     `json:"excluded_files,omitempty"`
	Viewport             Viewport            `json:"viewport,omitempty"`
	EnforceSslValidation bool                `json:"enforce_ssl_validation,omitempty"`
	ThresholdMonitors    []ThresholdMonitors `json:"threshold_monitors,omitempty"`
	DNSOverrides         DNSOverrides        `json:"dns_overrides,omitempty"`
	Connection           Connection          `json:"connection,omitempty"`
	WaitForFullMetrics   bool                `json:"wait_for_full_metrics,omitempty"`
}
