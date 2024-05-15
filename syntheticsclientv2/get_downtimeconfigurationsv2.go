// Copyright 2024 Splunk, Inc.
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
	"bytes"
	"encoding/json"
	"fmt"
)

func parseDowntimeConfigurationV2Response(response string) (*DowntimeConfigurationV2Response, error) {
	// Parse the response and return the check object
	var check DowntimeConfigurationV2Response
	err := json.Unmarshal([]byte(response), &check)
	if err != nil {
		return nil, err
	}

	return &check, err
}

func (c Client) GetDowntimeConfigurationV2(id int) (*DowntimeConfigurationV2Response, *RequestDetails, error) {

	details, err := c.makePublicAPICall("GET",
		fmt.Sprintf("/downtime_configurations/%d", id),
		bytes.NewBufferString("{}"),
		nil)

	if err != nil {
		return nil, details, err
	}

	check, err := parseDowntimeConfigurationV2Response(details.ResponseBody)
	if err != nil {
		return check, details, err
	}

	return check, details, nil
}

func parseDowntimeConfigurationsV2Response(response string) (*DowntimeConfigurationsV2Response, error) {
	// Parse the response and return the check object
	var check DowntimeConfigurationsV2Response
	err := json.Unmarshal([]byte(response), &check)
	if err != nil {
		return nil, err
	}

	return &check, err
}

func (c Client) GetDowntimeConfigurationsV2() (*DowntimeConfigurationsV2Response, *RequestDetails, error) {

	details, err := c.makePublicAPICall("GET",
		"/downtime_configurations",
		bytes.NewBufferString("{}"),
		nil)

	if err != nil {
		return nil, details, err
	}

	check, err := parseDowntimeConfigurationsV2Response(details.ResponseBody)
	if err != nil {
		return check, details, err
	}

	return check, details, nil
}
