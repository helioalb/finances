package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailInUse   = errors.New("email already in use")
)
