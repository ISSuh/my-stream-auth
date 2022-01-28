package user

import (
	"github.com/ISSuh/my-stream-auth/auth"
)

type Session struct {
	User  Account
	Token auth.Token
}
