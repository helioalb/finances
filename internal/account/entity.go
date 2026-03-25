package account

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        int64
	UUID      uuid.UUID
	UserID    int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
