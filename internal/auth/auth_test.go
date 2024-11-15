package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		inputHeader   http.Header
		expected      string
		errorContains string
	}{
		{
			name: "valid test",
			inputHeader: http.Header{
				"Authorization": {"ApiKey 123456789"},
			},
			expected: "123456789",
		},
		{
			name:          "no header test",
			inputHeader:   http.Header{},
			expected:      "",
			errorContains: "no authorization header included",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			key, err := GetAPIKey(tc.inputHeader)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if key != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: tokens don't match", i, tc.name)
			}
		})
	}
}
