.PHONY: default all build clean test test-race test-cover test-integration fmt fmtcheck vet lint govulncheck actionlint

FILES=./...

default: test

all: clean build test

build: fmtcheck
	go build $(FILES)

clean:
	@echo "==> Cleaning out old builds "
	go clean
	rm -rf coverage.txt test-results.json main.breakdown integration.jsonl

fmt:
	@echo "==> Fixing source code with gofmt "
	gofmt -s -w .

fmtcheck:
	@echo "==> Checking source code formatting "
	@unformatted=$$(gofmt -s -l .); \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not gofmt'd:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

vet:
	@echo "==> Running go vet "
	go vet $(FILES)

lint:
	@echo "==> Checking source code against linters "
	@GOGC=30 golangci-lint run $(FILES)

govulncheck:
	@echo "==> Checking for known vulnerabilities "
	govulncheck $(FILES)

actionlint:
	@echo "==> Checking GitHub Actions workflows "
	actionlint

test: fmtcheck vet
	@echo "==> Running all tests"
	go test $(FILES) -v -timeout=30s -parallel=4 -cover

test-race: fmtcheck vet
	@echo "==> Running all tests with the race detector"
	go test $(FILES) -race -timeout=60s -parallel=4

test-cover: clean fmtcheck vet
	@echo "==> Running all tests with coverage and JSON output"
	go test $(FILES) -timeout=30s -parallel=8 -json -cover -covermode=atomic -coverprofile coverage.txt > test-results.json

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
