package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int64
	UUID      uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	// TODO: add email. Should be unique
}

func Create(name string) *User {
	return &User{
		UUID: uuid.New(),
		Name: name,
	}
}
