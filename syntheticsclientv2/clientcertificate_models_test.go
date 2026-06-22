//go:build unit_tests
// +build unit_tests

// Copyright 2026 Splunk, Inc.
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
	"encoding/json"
	"strings"
	"testing"
)

func TestClientCertificateV2InputUsesPublicJSONCasing(t *testing.T) {
	payload := ClientCertificateV2Input{
		Certificate: ClientCertificateInput{
			Name:        "mtls_api_example",
			Description: "mTLS certificate for api.example.com",
			Domain:      "api.example.com",
			PublicKey: ClientCertificateKeyInput{
				Content:       "base64-public",
				Filename:      "client.crt",
				FileExtension: "pem",
			},
			PrivateKey: ClientCertificatePrivateKeyInput{
				Content:       "base64-private",
				Filename:      "client.key",
				FileExtension: "pem",
				Password:      "key-password",
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	jsonBody := string(body)

	for _, field := range []string{`"certificate"`, `"publicKey"`, `"privateKey"`, `"fileExtension"`, `"password"`} {
		if !strings.Contains(jsonBody, field) {
			t.Fatalf("request body missing %s: %s", field, jsonBody)
		}
	}
	for _, field := range []string{"public_key", "private_key", "file_extension", "fileName"} {
		if strings.Contains(jsonBody, field) {
			t.Fatalf("request body contains internal field %s: %s", field, jsonBody)
		}
	}
}

func TestClientCertificateV2ResponseParsesPublicJSONCasing(t *testing.T) {
	body := `{
		"certificate": {
			"id": 123,
			"name": "mtls_api_example",
			"description": "mTLS certificate for api.example.com",
			"domain": "api.example.com",
			"expiresAt": "2027-01-02T03:04:05Z",
			"createdAt": "2026-01-02T03:04:05Z",
			"createdBy": "boris@example.com",
			"updatedAt": "2026-01-03T03:04:05Z",
			"updatedBy": "boris@example.com",
			"publicKey": {
				"id": 501,
				"content": "<REDACTED>",
				"filename": "client.crt",
				"fileExtension": "pem"
			},
			"privateKey": {
				"id": 502,
				"content": "<REDACTED>",
				"filename": "client.key",
				"fileExtension": "pem",
				"password": "<REDACTED>"
			}
		}
	}`

	var response ClientCertificateV2Response
	if err := json.NewDecoder(strings.NewReader(body)).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if response.Certificate.ID != 123 {
		t.Fatalf("certificate id %d want 123", response.Certificate.ID)
	}
	if expiresAt := response.Certificate.ExpiresAt.Format("2006-01-02T15:04:05Z"); expiresAt != "2027-01-02T03:04:05Z" {
		t.Fatalf("expiresAt %s want 2027-01-02T03:04:05Z", expiresAt)
	}
	if response.Certificate.PublicKey.Content != "<REDACTED>" {
		t.Fatalf("public key content %s want <REDACTED>", response.Certificate.PublicKey.Content)
	}
	if response.Certificate.PrivateKey.Password != "<REDACTED>" {
		t.Fatalf("private key password %s want <REDACTED>", response.Certificate.PrivateKey.Password)
	}
}
