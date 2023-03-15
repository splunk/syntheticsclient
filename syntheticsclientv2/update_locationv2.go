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
	"bytes"
	"encoding/json"
	"fmt"
)

func parseUpdateLocationV2Response(response string) (*LocationV2Response, error) {
	var updateLocationV2 LocationV2Response
	err := json.Unmarshal([]byte(response), &updateLocationV2)
	if err != nil {
		return nil, err
	}

	return &updateLocationV2, err
}

func (c Client) UpdateLocationV2(id string, LocationV2Details *LocationV2Input) (*LocationV2Response, *RequestDetails, error) {

	body, err := json.Marshal(LocationV2Details)
	if err != nil {
		return nil, nil, err
	}

	requestDetails, err := c.makePublicAPICall("PUT", fmt.Sprintf("/locations/%s", id), bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, requestDetails, err
	}

	updateLocationV2, err := parseUpdateLocationV2Response(requestDetails.ResponseBody)
	if err != nil {
		return updateLocationV2, requestDetails, err
	}

	return updateLocationV2, requestDetails, nil
}
