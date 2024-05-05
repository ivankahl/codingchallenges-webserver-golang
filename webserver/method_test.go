package webserver

import (
	"testing"
)

func TestMethodFromStringValid(t *testing.T) {
	var tests = []struct {
		input    string
		expected Method
	}{
		{"GET", MethodGet},
		{"HEAD", MethodHead},
		{"POST", MethodPost},
		{"PUT", MethodPut},
		{"DELETE", MethodDelete},
		{"CONNECT", MethodConnect},
		{"OPTIONS", MethodOptions},
		{"TRACE", MethodTrace},
		{"PATCH", MethodPatch},
		{"*", MethodAny},
	}

	for _, test := range tests {
		method, err := methodFromString(test.input)
		if err != nil {
			t.Errorf("methodFromString(%q) unexpected error: %v", test.input, err)
		}
		if method != test.expected {
			t.Errorf("methodFromString(%q) expected %q, got %q", test.input, test.expected, method)
		}
	}
}

func TestMethodFromStringInvalid(t *testing.T) {
	var tests = []struct {
		input         string
		expectedError error
	}{
		{"", ErrInvalidMethod},
		{"INVALID", ErrInvalidMethod},
	}

	for _, test := range tests {
		method, err := methodFromString(test.input)
		if err == nil {
			t.Errorf("methodFromString(%q) expected error, got nil", test.input)
		}
		if err != test.expectedError {
			t.Errorf("methodFromString(%q) expected error %v, got %v", test.input, test.expectedError, err)
		}
		if method != "" {
			t.Errorf("methodFromString(%q) expected empty method on error, got %q", test.input, method)
		}
	}
}
