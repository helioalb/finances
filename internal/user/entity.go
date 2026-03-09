package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int64
	UUID      uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) New() *User {
	return &User{
		UUID:  uuid.New(),
		Name:  u.Name,
		Email: u.Email,
	}
}
