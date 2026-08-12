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
	"net/http"
	"reflect"
	"testing"
)

var getChromeFlagsBody = `{"chromeFlags":[{"name":"disable-gpu"},{"name":"lang","value":"en-US"}]}`

func TestGetChromeFlags(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/chrome_flags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if r.URL.Path != "/chrome_flags" {
			t.Fatalf("path = %q, want /chrome_flags", r.URL.Path)
		}
		_, err := w.Write([]byte(getChromeFlagsBody))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.GetChromeFlags()
	if err != nil {
		t.Fatal(err)
	}
	if details.StatusCode != http.StatusOK {
		t.Fatalf("status code = %d, want %d", details.StatusCode, http.StatusOK)
	}

	langValue := "en-US"
	expected := []ChromeFlag{
		{Name: "disable-gpu"},
		{Name: "lang", Value: &langValue},
	}
	if !reflect.DeepEqual(resp.ChromeFlags, expected) {
		t.Fatalf("ChromeFlags = %#v, want %#v", resp.ChromeFlags, expected)
	}
}

func TestParseChromeFlagsResponse(t *testing.T) {
	resp, err := parseChromeFlagsResponse(getChromeFlagsBody)
	if err != nil {
		t.Fatal(err)
	}

	langValue := "en-US"
	expected := []ChromeFlag{
		{Name: "disable-gpu"},
		{Name: "lang", Value: &langValue},
	}
	if !reflect.DeepEqual(resp.ChromeFlags, expected) {
		t.Fatalf("ChromeFlags = %#v, want %#v", resp.ChromeFlags, expected)
	}
}

func TestGetChromeFlagsReturnsErrorOnMalformedResponse(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/chrome_flags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, err := w.Write([]byte("{not valid json"))
		if err != nil {
			t.Fatal(err)
		}
	})

	resp, details, err := testClient.GetChromeFlags()
	if err == nil {
		t.Fatal("expected a parse error, but got none")
	}
	if details == nil {
		t.Fatal("expected request details")
	}
	if resp != nil {
		t.Errorf("expected nil response, got %#v", resp)
	}
}

func TestGetChromeFlagsReturnsErrorOnNetworkFailure(t *testing.T) {
	unreachableClient := NewConfigurableClient("apiKey", "realm", ClientArgs{publicBaseUrl: "http://127.0.0.1:1"})
	_, _, err := unreachableClient.GetChromeFlags()
	if err == nil {
		t.Fatal("expected a connection error")
	}
}

func TestParseChromeFlagsResponseReturnsErrorOnMalformedJSON(t *testing.T) {
	resp, err := parseChromeFlagsResponse("{not valid json")
	if err == nil {
		t.Fatal("expected a parse error, but got none")
	}
	if resp != nil {
		t.Errorf("expected nil response, got %#v", resp)
	}
}
