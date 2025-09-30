package goexercises

import "time"

type UserSession interface {
	Username() string
	IsAdmin() bool
	ExpiresAt() time.Time
}

type UserSessionImpl struct {
	username  string
	isAdmin   bool
	createdAt time.Time
}

const TokenExpirationTime = 24 * time.Hour

func NewUserSession(username string, isAdmin bool) UserSession {
	// TODO: Implement
	return nil
}

// For tests, do not change this function
func NewUserSessionTest(username string, isAdmin bool, createdAt time.Time) UserSession {
	return &UserSessionImpl{
		username:  username,
		isAdmin:   isAdmin,
		createdAt: createdAt,
	}
}

func (u UserSessionImpl) Username() string {
	// TODO: Implement
	return ""
}

func (u UserSessionImpl) IsAdmin() bool {
	// TODO: Implement
	//
	return false
}

func (u UserSessionImpl) ExpiresAt() time.Time {
	// TODO: Implement
	// tips: TokenExpirationTime + createdAt
	return time.Time{}
}
