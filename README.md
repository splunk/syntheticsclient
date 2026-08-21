# syntheticsclient

[![Release](https://img.shields.io/github/v/release/splunk/syntheticsclient)](https://github.com/splunk/syntheticsclient/releases)
[![CI Checks](https://img.shields.io/github/actions/workflow/status/splunk/syntheticsclient/ci.yml?branch=v2&label=CI)](https://github.com/splunk/syntheticsclient/actions/workflows/ci.yml?query=branch%3Av2)
[![Build](https://img.shields.io/github/actions/workflow/status/splunk/syntheticsclient/ci.yml?branch=v2&label=build)](https://github.com/splunk/syntheticsclient/actions/workflows/ci.yml?query=branch%3Av2)
[![License](https://img.shields.io/github/license/splunk/syntheticsclient)](https://github.com/splunk/syntheticsclient/blob/v2/LICENSE)

A Splunk Synthetics for Splunk Observability (Formerly Rigor) client for golang.

## Installation
`go get https://github.com/splunk/syntheticsclient.git`

## Development

The supported Go baseline for `main` is the version pinned in [`.go-version`](./.go-version)
(currently Go 1.26.7). Local tooling and CI both resolve their Go toolchain from that file.

Before opening a pull request, run the same checks CI runs:

```shell
make fmtcheck     # gofmt -l, fails on formatting drift
make vet          # go vet ./...
make lint         # golangci-lint v2.12.2, requires golangci-lint on PATH
make test-cover   # go test ./... with coverage
make test-race    # go test ./... -race
make govulncheck  # requires govulncheck on PATH
make actionlint   # lints .github/workflows, requires actionlint on PATH
```

`golangci-lint`, `govulncheck`, and `actionlint` are not vendored; install the versions CI
pins with:

```shell
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2
go install golang.org/x/vuln/cmd/govulncheck@v1.7.0
go install github.com/rhysd/actionlint/cmd/actionlint@v1.7.12
```

> **Note:** `github.com/splunk/syntheticsclient` (V1) is deprecated and excluded from
> build/coverage targets (see the Makefile); its tests still run and report as an explicit
> `SKIP`. It may be removed from this module in a future change (tracked separately).

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

## Request Details

V2 methods return `RequestDetails` for debugging failed or unexpected API calls.
Use `RequestDetails.RequestBody` when you need to inspect the outgoing request.
This field is sanitized before it is returned and redacts API tokens, certificate
content, passwords, and generic secret values.

`RequestDetails.RawRequest` is intentionally not populated by V2 public API calls
because a raw `http.Request` can retain authorization headers or request body
secrets.

## Additional Information
This client is largely a copypasta mutation of the [go-victor](https://github.com/victorops/go-victorops) client for Splunk On-Call (formerly known as VictorOps).

## Contributions
Contributions are welcome and encouraged!

Please see [CONTRIBUTING.md](./CONTRIBUTING.md) for details on contributing to this repository.

Before your contribution can be accepted, you will be asked to sign our
[Splunk Contributor License Agreement (CLA)](https://github.com/splunk/cla-agreement/blob/main/CLA.md).

To agree to the CLA and COC please comment these in **separate individual messages** on your PR:

CLA:
```
I have read the CLA Document and I hereby sign the CLA
```

Code of Conduct:
```
I have read the Code of Conduct and I hereby accept the Terms
