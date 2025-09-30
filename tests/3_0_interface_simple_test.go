package goexercises

import (
	. "goexercises"
	"testing"
	"time"
)

func TestNewUserSession(t *testing.T) {
	tests := []struct {
		name     string
		username string
		isAdmin  bool
	}{
		{"regular user", "john_doe", false},
		{"admin user", "admin", true},
		{"empty username", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := NewUserSession(tt.username, tt.isAdmin)

			if session == nil {
				t.Fatal("NewUserSession() returned nil")
			}

			if session.Username() != tt.username {
				t.Errorf("Username() = %v, want %v", session.Username(), tt.username)
			}

			if session.IsAdmin() != tt.isAdmin {
				t.Errorf("IsAdmin() = %v, want %v", session.IsAdmin(), tt.isAdmin)
			}

			// Check that ExpiresAt returns a time in the future
			expiresAt := session.ExpiresAt()
			if expiresAt.IsZero() {
				t.Error("ExpiresAt() returned zero time")
			}

			// Should expire after TokenExpirationTime from creation
			now := time.Now()
			expectedExpiry := now.Add(TokenExpirationTime)

			// Allow some tolerance for test execution time (1 second)
			if expiresAt.Before(now) || expiresAt.After(expectedExpiry.Add(time.Second)) {
				t.Errorf("ExpiresAt() = %v, expected around %v", expiresAt, expectedExpiry)
			}
		})
	}
}

func TestUserSessionImpl_Username(t *testing.T) {
	tests := []struct {
		name     string
		username string
		isAdmin  bool
	}{
		{"john_doe", "john_doe", false},
		{"admin", "admin", true},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := NewUserSession(tt.username, tt.isAdmin)

			if got := session.Username(); got != tt.username {
				t.Errorf("Username() = %v, want %v", got, tt.username)
			}
		})
	}
}

func TestUserSessionImpl_IsAdmin(t *testing.T) {
	tests := []struct {
		name    string
		isAdmin bool
	}{
		{"admin user", true},
		{"regular user", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := NewUserSession("test", tt.isAdmin)

			if got := session.IsAdmin(); got != tt.isAdmin {
				t.Errorf("IsAdmin() = %v, want %v", got, tt.isAdmin)
			}
		})
	}
}

func TestUserSessionImpl_ExpiresAt(t *testing.T) {
	session := NewUserSessionTest("test", false, time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))

	expiresAt1 := session.ExpiresAt()
	expiresAt2 := session.ExpiresAt()

	// Should return consistent values across multiple calls
	if !expiresAt1.Equal(expiresAt2) {
		t.Errorf("ExpiresAt() should return consistent values, got %v and %v", expiresAt1.Format(time.DateTime), expiresAt2.Format(time.DateTime))
	}

	// Should return a time in the future
	expectedExpiry := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)
	if !expiresAt1.Equal(expectedExpiry) {
		t.Errorf("ExpiresAt() = %s, should be equal to %s", expiresAt1.Format(time.DateTime), expectedExpiry.Format(time.DateTime))
	}
}
