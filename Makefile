.PHONY: default all build clean test fmtcheck test-cover test-integration

# syntheticsclient (v1) is deprecated and excluded from builds and coverage;
# only syntheticsclientv2 is built/covered. v1's tests still run via
# TEST_FILES below so `go test` reports them as an explicit SKIP (see
# skipDeprecated in syntheticsclient/synthetics_test.go) rather than the
# package silently vanishing from CI output.
FILES=./syntheticsclientv2/...
TEST_FILES=./syntheticsclient/... ./syntheticsclientv2/...

default: test

all: clean build test

build: fmtcheck
	go build $(FILES)

clean:
	@echo "==> Cleaning out old builds "
	go clean
	rm -rf coverage.txt test-results.json v2.breakdown integration.jsonl


fmt:
	@echo "==> Fixing source code with gofmt "
	gofmt -s -w .

lint:
	@echo "==> Checking source code against linters "
	@GOGC=30 golangci-lint run ./syntheticsclientv2/...

fmtcheck: fmt lint

test: fmtcheck
	@echo "==> Running all tests"
	go test $(TEST_FILES) -v -timeout=30s -parallel=4 -cover -coverpkg=$(FILES)

test-cover: clean fmtcheck
	@echo "==> Running all tests with coverage and JSON output"
	go test $(TEST_FILES) -timeout=30s -parallel=8 -json -cover -covermode=atomic -coverpkg=$(FILES) -coverprofile coverage.txt > test-results.json

# Runs the live integration suite (syntheticsclientv2/integration_test.go) against a real
# Synthetics org. Requires API_ACCESS_TOKEN and REALM in the environment. Not part of the
# default test/test-cover targets: it mutates live state and needs real credentials, so it
# only runs where those are deliberately provided (a developer's shell, or the gated
# integration-test CI job).
test-integration: SHELL:=/bin/bash
test-integration:
	@echo "==> Running live integration tests"
	set -o pipefail; go test -tags=integration -json ./syntheticsclientv2/... -timeout 30m \
		| sed '/X-Sf-Token/d' \
		| tee integration.jsonl \
		| jq -j -r 'if .Action == "output" then .Output else empty end'
