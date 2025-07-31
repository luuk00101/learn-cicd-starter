package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		authHeader    string
		expectedKey   string
		expectedError error
	}{
		"valid API key":                   {authHeader: "ApiKey test-key-123", expectedKey: "test-key-123", expectedError: nil},
		"no authorization header":         {authHeader: "", expectedKey: "", expectedError: auth.ErrNoAuthHeaderIncluded},
		"malformed header - wrong prefix": {authHeader: "Bearer test-key-123", expectedKey: "", expectedError: auth.ErrMalformedAuthHeader},
		"malformed header - no key":       {authHeader: "ApiKey", expectedKey: "", expectedError: auth.ErrMalformedAuthHeader},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			headers.Set("Authorization", tt.authHeader)

			key, err := auth.GetAPIKey(headers)

			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() key = %v, want %v", key, tt.expectedKey)
			}

			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("GetAPIKey() error = nil, want %v", tt.expectedError)
				} else if err.Error() != tt.expectedError.Error() {
					t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expectedError)
				}
			}
		})
	}
}
