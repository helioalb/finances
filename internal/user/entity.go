package user

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        int64
	UUID      uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
