package goexercises

// error is an builtin interface in Go of the following signature:
// type error interface {
// 	Error() string
// }

// You can implement this interface by creating a struct with an Error method.
// This is a simple way to create custom errors in Go.

type MyError struct {
	title, message string
}

// this is a private function, do not change the name
func newMyError(title, message string) MyError {
	return MyError{title, message}
}

var ErrUserNotFound = newMyError("user not found", "the user you are looking for does not exist")

func (e MyError) Error() string {
	// TODO: Implement
	// should be formatted as "{Title}: {Message}"
	return ""
}

func IsMyErrorType(err error) bool {
	// TODO: Implement
	// should return true if the error type is a MyError
	// tips: errors.As() can be used to check the error type
	return false
}

func IsErrUserNotFound(err error) bool {
	// TODO: Implement
	// should return true if the error is an ErrUserNotFound
	// tips: errors.Is() can be used to check the error object
	return false
}

// for tests, do not change this function
// in real world, you should not expose this function
func NewMyErrorTest(title, message string) MyError {
	return newMyError(title, message)
}
