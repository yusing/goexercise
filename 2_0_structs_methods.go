package goexercises

// User represents a user with validation.
// Implement methods that demonstrate pointer vs value receivers.
type User struct {
	Name  string
	Email string
	Age   int
}

// Validate checks if the user data is valid.
// Should return error if name is empty, email doesn't contain @, or age < 0.
func (u User) Validate() error {
	// TODO: implement
	return nil
}

// UpdateEmail updates the user's email if it's valid (contains @).
// Uses pointer receiver to modify the struct.
// Returns error if email is invalid.
func (u *User) UpdateEmail(newEmail string) error {
	// TODO: implement
	return nil
}
