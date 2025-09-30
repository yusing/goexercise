package goexercises

import (
	. "goexercises"
	"testing"
)

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{"valid user", User{Name: "John", Email: "john@example.com", Age: 25}, false},
		{"empty name", User{Name: "", Email: "john@example.com", Age: 25}, true},
		{"invalid email", User{Name: "John", Email: "notanemail", Age: 25}, true},
		{"negative age", User{Name: "John", Email: "john@example.com", Age: -1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserUpdateEmail(t *testing.T) {
	user := User{Name: "John", Email: "old@example.com", Age: 25}

	err := user.UpdateEmail("new@example.com")
	if err != nil {
		t.Errorf("UpdateEmail() error = %v", err)
	}
	if user.Email != "new@example.com" {
		t.Errorf("UpdateEmail() email = %v, want new@example.com", user.Email)
	}

	err = user.UpdateEmail("invalidemail")
	if err == nil {
		t.Error("UpdateEmail() with invalid email should return error")
	}
	if user.Email != "new@example.com" {
		t.Error("UpdateEmail() should not change email on error")
	}
}
