package tests

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKeySuccess(t *testing.T) {
	correctApiKey := "djfaksf;asjfja;sfj;akj"
	apiKeyHeader := "ApiKey " + correctApiKey
	header := http.Header{
		"Authorization": {apiKeyHeader},
	}
	apiKey, err := auth.GetAPIKey(header)
	if err != nil {
		t.Errorf("test api key failure : %v", err)
	}

	if apiKey != correctApiKey {
		t.Errorf("api keys do not match")
	}
}

func TestGetApiKeyFailure(t *testing.T) {
	correctApiKey := "djfaksf;asjfja;sfj;akj"
	apiKeyHeader := "Bearer " + correctApiKey
	header := http.Header{
		"Authorization": {apiKeyHeader},
	}
	apiKey, err := auth.GetAPIKey(header)

	// TODO: - Purposely broken code. Change back to ==
	if err != nil {
		t.Errorf("not returning malformed api key format")
	}

	if apiKey != "" {
		t.Errorf("should not return the correct api key")
	}
}
