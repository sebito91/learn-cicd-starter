package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyWithEmptyHeader(t *testing.T) {
	// initialize empty header
	header := http.Header{}

	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected an error due to empty header but didn't get one")
	}
}

func TestGetAPIKeyWithBadAuthorizationMarker(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "GarbageKey 12345")

	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected an error due to incorrect Authorization header but didn't get one")
	}
}

func TestGetAPIKeyWithValidAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "APIKey 12345")

	key, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("did not expect error with valid API Key, got: %s", err)
	}

	if key != "12345" {
		t.Fatalf("incorrect APIKey retrieved, expected: 12345, got: %s", key)
	}
}
