package auth

import (
	"net/http"
	"testing"
	"errors"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		input          http.Header
		expected       string
		expectedErr		 error
	}{
		{
			input: http.Header{
				"Authorization": []string{"ApiKey 123456"},
			},
			expected: "123456",
			expectedErr: nil,
		},
		{
			input: http.Header{},
			expected: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, c := range cases {
		actual, err := GetAPIKey(c.input)
		if c.expectedErr != nil {
			if !errors.Is(err, c.expectedErr) {
				t.Errorf("Expected error %v, but got %v", c.expectedErr, err)
			}
			continue
		}

		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		if actual != c.expected {
			t.Errorf("Expected: %s, but got %s", c.expected, actual)
		}
	}
}
