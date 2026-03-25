package user

import "errors"

var (
	errUserNotFound = errors.New("user not found")
	errEmailInUse   = errors.New("email already in use")
)
