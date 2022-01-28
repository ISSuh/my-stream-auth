package storage

import (
	"github.com/ISSuh/my-stream-auth/user"
)

// Storage for user account
type AccountStorage interface {
	InitAccountStorage() error

	CreateAccount(user user.Account) error
	DeleteAccount(user user.Account) error
	// AccessAccount(user User) Account
}

// Storage for active session
type SessionStorage interface {
	InitSessionStorage() error

	AddSession(session user.Session) error
	// DeleteSession(user User) error
	// GetSession(user User) (*Session, error)
}
