package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		header         http.Header
		expected_token string
		wantErr        bool
	}{
		{
			name:           "Valid token",
			header:         http.Header{"Authorization": []string{"ApiKey token"}},
			expected_token: "token",
			wantErr:        false,
		},
		{
			name:           "No authorization header",
			header:         http.Header{},
			expected_token: "",
			wantErr:        true,
		},
		{
			name:           "Field does not start with Bearer",
			header:         http.Header{"Authorization": []string{"token testi1"}},
			expected_token: "",
			wantErr:        true,
		},
		{
			name:           "No space",
			header:         http.Header{"Authorization": []string{"ApiKeytoken"}},
			expected_token: "",
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokenString, err := GetAPIKey(test.header)
			if (err != nil) != test.wantErr {
				t.Errorf("TestGetAPIKey() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if tokenString != test.expected_token {
				t.Errorf("TestGetAPIKey() tokenString = %v, want %v", tokenString, test.expected_token)
			}
		})
	}
}
