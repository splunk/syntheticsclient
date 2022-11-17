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
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	testClient *Client

	// server is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testClient = NewConfigurableClient("apiKey", "realm", ClientArgs{
		publicBaseUrl: testServer.URL,
	})
	log.Printf("Client instantiated: %s", testClient.publicBaseURL)
}

func teardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestConfigurableClient(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	args := ClientArgs{
		timeoutSeconds: 30,
		publicBaseUrl:  testServer.URL,
	}

	testConfigurableClient := NewConfigurableClient("snakedonut", "realm-o", args)
	log.Printf("Client instantiated: %s", testServer.URL)
	if testConfigurableClient.GetHTTPClient() == nil {
		t.Errorf("http client is nil")
	}
	if testConfigurableClient.apiKey != "snakedonut" {
		t.Errorf("apiKey was not correctly passed")
	}
	if testConfigurableClient.apiKey != "realm-o" {
		t.Errorf("apiKey was not correctly passed")
	}
}

func TestConfigurableClientTimeout(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testMux.HandleFunc("/v2/checks/12", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	})

	testConfigurableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{
		timeoutSeconds: 1,
		publicBaseUrl:  testServer.URL,
	})
	log.Printf("Client instantiated: %s", testServer.URL)
	_, _, err := testConfigurableClient.GetCheck(12)
	if !strings.Contains(err.Error(), "context deadline exceeded (Client.Timeout exceeded while awaiting headers)") {
		t.Errorf("expected to see timeout error, but saw: %s", err.Error())
	}
}
