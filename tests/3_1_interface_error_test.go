package goexercises

import (
	"errors"
	"fmt"
	. "goexercises"
	"testing"
)

func TestMyError_Error(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		message string
		want    string
	}{
		{
			name:    "basic error",
			title:   "validation error",
			message: "invalid input",
			want:    "validation error: invalid input",
		},
		{
			name:    "user not found",
			title:   "user not found",
			message: "the user you are looking for does not exist",
			want:    "user not found: the user you are looking for does not exist",
		},
		{
			name:    "empty message",
			title:   "empty",
			message: "",
			want:    "empty: ",
		},
		{
			name:    "empty title",
			title:   "",
			message: "some message",
			want:    ": some message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewMyErrorTest(tt.title, tt.message)
			if got := err.Error(); got != tt.want {
				t.Errorf("MyError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrUserNotFound(t *testing.T) {
	// Test that ErrUserNotFound is properly formatted
	want := "user not found: the user you are looking for does not exist"
	if got := ErrUserNotFound.Error(); got != want {
		t.Errorf("ErrUserNotFound.Error() = %v, want %v", got, want)
	}
}

func TestIsMyErrorType(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "direct MyError",
			err:  ErrUserNotFound,
			want: true,
		},
		{
			name: "wrapped MyError with fmt.Errorf",
			err:  fmt.Errorf("failed to get user: %w", ErrUserNotFound),
			want: true,
		},
		{
			name: "double wrapped MyError",
			err:  fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", ErrUserNotFound)),
			want: true,
		},
		{
			name: "standard error",
			err:  errors.New("standard error"),
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "wrapped standard error",
			err:  fmt.Errorf("wrapped: %w", errors.New("standard")),
			want: false,
		},
		{
			name: "custom MyError instance",
			err:  NewMyErrorTest("custom", "error"),
			want: true,
		},
		{
			name: "wrapped custom MyError instance",
			err:  fmt.Errorf("context: %w", NewMyErrorTest("custom", "error")),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMyErrorType(tt.err); got != tt.want {
				t.Errorf("IsMyErrorType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsErrUserNotFound(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "exact ErrUserNotFound",
			err:  ErrUserNotFound,
			want: true,
		},
		{
			name: "wrapped ErrUserNotFound",
			err:  fmt.Errorf("database error: %w", ErrUserNotFound),
			want: true,
		},
		{
			name: "double wrapped ErrUserNotFound",
			err:  fmt.Errorf("service error: %w", fmt.Errorf("repository error: %w", ErrUserNotFound)),
			want: true,
		},
		{
			name: "different MyError",
			err:  NewMyErrorTest("different", "error"),
			want: false,
		},
		{
			name: "wrapped different MyError",
			err:  fmt.Errorf("wrapped: %w", NewMyErrorTest("different", "error")),
			want: false,
		},
		{
			name: "standard error",
			err:  errors.New("user not found"),
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "MyError with same message but different instance",
			err:  NewMyErrorTest("user not found", "the user you are looking for does not exist"),
			want: false,
		},
		{
			name: "wrapped MyError with same message but different instance",
			err:  fmt.Errorf("context: %w", NewMyErrorTest("user not found", "the user you are looking for does not exist")),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsErrUserNotFound(tt.err); got != tt.want {
				t.Errorf("IsErrUserNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestErrorWrappingChain tests complex error wrapping scenarios
func TestErrorWrappingChain(t *testing.T) {
	// Create a chain: outer -> middle -> inner (ErrUserNotFound)
	inner := ErrUserNotFound
	middle := fmt.Errorf("middleware error: %w", inner)
	outer := fmt.Errorf("handler error: %w", middle)

	// Should be able to find MyError type at any level
	if !IsMyErrorType(outer) {
		t.Error("IsMyErrorType() should find MyError in deeply wrapped chain")
	}

	// Should be able to find specific error instance at any level
	if !IsErrUserNotFound(outer) {
		t.Error("IsErrUserNotFound() should find ErrUserNotFound in deeply wrapped chain")
	}

	// Should work with errors.Unwrap
	if unwrapped := errors.Unwrap(outer); unwrapped != middle {
		t.Error("errors.Unwrap() should return the middle error")
	}

	if unwrapped := errors.Unwrap(middle); unwrapped != inner {
		t.Error("errors.Unwrap() should return the inner error")
	}
}

// TestMultipleErrorWrapping tests scenarios with multiple wrapped errors
func TestMultipleErrorWrapping(t *testing.T) {
	err1 := NewMyErrorTest("error1", "first error")
	err2 := NewMyErrorTest("error2", "second error")

	wrapped1 := fmt.Errorf("context1: %w", err1)
	wrapped2 := fmt.Errorf("context2: %w", err2)

	// Both should be identified as MyError type
	if !IsMyErrorType(wrapped1) {
		t.Error("IsMyErrorType() should identify wrapped1 as MyError")
	}
	if !IsMyErrorType(wrapped2) {
		t.Error("IsMyErrorType() should identify wrapped2 as MyError")
	}

	// Neither should be ErrUserNotFound
	if IsErrUserNotFound(wrapped1) {
		t.Error("IsErrUserNotFound() should not match wrapped1")
	}
	if IsErrUserNotFound(wrapped2) {
		t.Error("IsErrUserNotFound() should not match wrapped2")
	}
}

// TestErrorInterface tests that MyError properly implements error interface
func TestErrorInterface(t *testing.T) {
	var err error = ErrUserNotFound

	if err.Error() == "" {
		t.Error("MyError.Error() should not return empty string")
	}

	// Should be usable as error
	errMsg := err.Error()
	if errMsg != "user not found: the user you are looking for does not exist" {
		t.Errorf("error message = %v, want 'user not found: the user you are looking for does not exist'", errMsg)
	}
}

// TestErrorsAsExtraction tests extracting the actual MyError from wrapped errors
func TestErrorsAsExtraction(t *testing.T) {
	original := NewMyErrorTest("test", "test message")
	wrapped := fmt.Errorf("context: %w", original)

	var target MyError
	if !errors.As(wrapped, &target) {
		t.Fatal("errors.As() should successfully extract MyError from wrapped error")
	}

	if target.Error() != original.Error() {
		t.Errorf("extracted error = %v, want %v", target.Error(), original.Error())
	}
}

// TestErrorsIsComparison tests comparing error instances
func TestErrorsIsComparison(t *testing.T) {
	err := fmt.Errorf("wrapped: %w", ErrUserNotFound)

	if !errors.Is(err, ErrUserNotFound) {
		t.Error("errors.Is() should match ErrUserNotFound in wrapped error")
	}

	differentErr := NewMyErrorTest("different", "error")
	if errors.Is(err, differentErr) {
		t.Error("errors.Is() should not match different MyError instance")
	}
}
