package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int64
	UUID      uuid.UUID
	OwnerID   int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
