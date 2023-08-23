# syntheticsclient
A Splunk Synthetics (Formerly Rigor) client for golang.

## Installation
`go get https://github.com/splunk/syntheticsclient.git`

## Important Note

V2 client is used to make API calls and CRUD operations to the Splunk Observability Synthetics endpoints (E.G. [API Tests](https://dev.splunk.com/observability/reference/api/synthetics_api_tests/))

**Deprecated** V1 Client is used to make the API calls for the [Splunk Synthetics (Formerly Rigor) public API](https://monitoring-api.rigor.com/). 

## Example Usages
```go
package main

import (
	"fmt"
	"os"
	"encoding/json"
	sc2 "github.com/splunk/syntheticsclient/v2/syntheticsclientv2"
)

func main() {
	//Expects a token is available from the API_ACCESS_TOKEN environment variable
	//Expects a realm (e.g. us1) is available from REALM environment variable
	token := os.Getenv("API_ACCESS_TOKEN")
	realm := os.Getenv("REALM")

	//Create your client with the token
	c := sc2.NewClient(token, realm)

	//Take your ugly (but valid) JSON string as bytes and unmarshal into a CreateHttpCheckV2 struct
	jsonData := []byte(`{"test":{"name":"http-test","type":"http","url":"https://www.splunk.com","locationIds":["aws-us-east-1"],"frequency":10,"schedulingStrategy":"round_robin","active":true,"requestMethod":"GET","body":null,"headers":[{"name":"boop","value":"beep"}]}}`)
	var httpCheckDetail sc2.HttpCheckV2Input
	err := json.Unmarshal(jsonData, &httpCheckDetail)
	if err != nil {
		fmt.Println(err)
	}

	//Use your converted JSON to make the request and print
	res, _, err := c.CreateHttpCheckV2(&httpCheckDetail)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
```

## API Documentation
API Docs are [available here](https://dev.splunk.com/observability/reference)

## Additional Information
This client is largely a copypasta mutation of the [go-victor](https://github.com/victorops/go-victorops) client for Splunk On-Call (formerly known as VictorOps).