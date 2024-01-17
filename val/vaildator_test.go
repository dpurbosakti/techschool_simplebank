package val

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateUsername(t *testing.T) {
	testCases := []struct {
		name   string
		value  string
		expect error
	}{
		// Test valid username
		{"valid username", "john_doe123", nil},

		// Test username with minimum length
		{"username with minimum length", "abc", nil},

		// Test username with maximum length
		{"username with maximum length", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuv", nil},

		// Test username with invalid characters
		{"username with invalid characters", "john.doe", fmt.Errorf("must contain only lowercase letters, digits, or underscore")},
		{"username with invalid characters", "JohnDoe", fmt.Errorf("must contain only lowercase letters, digits, or underscore")},
		{"username with invalid characters", "john_doe!", fmt.Errorf("must contain only lowercase letters, digits, or underscore")},
		{"username with invalid characters", "john_doe@123", fmt.Errorf("must contain only lowercase letters, digits, or underscore")},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateUsername(tc.value)
			require.Equal(t, tc.expect, err)
		})

	}
}

func TestValidateFullName(t *testing.T) {
	// Test cases for ValidateString
	t.Run("valid input", func(t *testing.T) {
		err := ValidateFullName("John Doe")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("value is too short", func(t *testing.T) {
		err := ValidateFullName("Jo")
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})

	t.Run("value is too long", func(t *testing.T) {
		err := ValidateFullName("John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John Doe John")
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})

	// Test case for isValidFullName
	t.Run("contains invalid characters", func(t *testing.T) {
		err := ValidateFullName("John Doe!")
		if err == nil {
			t.Error("expected error, but got nil")
		}
	})
}

// func TestValidateEmail(t *testing.T) {
// 	testCases := []struct {
// 		name   string
// 		email  string
// 		expect error
// 	}{
// 		{
// 			name:   "Valid email address",
// 			email:  "test@example.com",
// 			expect: nil,
// 		},
// 		{
// 			name:   "Email address with missing domain",
// 			email:  "test@",
// 			expect: fmt.Errorf("is not a valid email address"),
// 		},
// 		{
// 			name:   "Email address with invalid format",
// 			email:  "test.example.com",
// 			expect: fmt.Errorf("is not a valid email address"),
// 		},
// 		{
// 			name:   "Email address with special characters",
// 			email:  "!@#$%^&*()_+=-{}[]|\\:\";'<>?,./test@example.com",
// 			expect: nil,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := ValidateEmail(tc.email)
// 			fmt.Println("ERROR: ", err)
// 			if err != tc.expect {
// 				t.Errorf("Expected error: %v, but got: %v", tc.expect, err)
// 			}
// 		})
// 	}
// }
