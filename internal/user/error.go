package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailInUse   = errors.New("email already in use")
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
)
