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
)


func parseCreateApiCheckV2Response(response string) (*ApiCheckV2Response, error) {

	var createApiCheckV2 ApiCheckV2Response
	JSONResponse := []byte(response)
	err := json.Unmarshal(JSONResponse, &createApiCheckV2)
	if err != nil {
		return nil, err
	}

	return &createApiCheckV2, err
}

func (c Client) CreateApiCheckV2(ApiCheckV2Details *ApiCheckV2Input) (*ApiCheckV2Response, *RequestDetails, error) {

	body, err := json.Marshal(ApiCheckV2Details)
	if err != nil {
		return nil, nil, err
	}

	details, err := c.makePublicAPICall("POST", "/tests/api", bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, details, err
	}

	newApiCheckV2, err := parseCreateApiCheckV2Response(details.ResponseBody)
	if err != nil {
		return newApiCheckV2, details, err
	}

	return newApiCheckV2, details, nil
}
