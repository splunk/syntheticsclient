.PHONY: default all build clean test fmtcheck test-cover

FILES=./...

default: test

all: clean build test

build: fmtcheck
	go build -tags=unit_tests ./...

clean:
	@echo "==> Cleaning out old builds "
	go clean
	rm -rf coverage.txt test-results.json


fmt:
	@echo "==> Fixing source code with gofmt "
	gofmt -s -w .

lint:
	@echo "==> Checking source code against linters "
	@GOGC=30 golangci-lint run ./syntheticsclientv2/...

fmtcheck: fmt lint

test: fmtcheck
	@echo "==> Running all tests"
	go test $(FILES) -v -tags=unit_tests -timeout=30s -parallel=4 -cover

test-cover: clean fmtcheck
	@echo "==> Running all tests with coverage and JSON output"
	go test $(FILES) -tags=unit_tests -timeout=30s -parallel=8 -json -cover -covermode=atomic -coverprofile coverage.txt > test-results.json
