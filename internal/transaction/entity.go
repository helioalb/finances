package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	UUID        uuid.UUID
	AccountID   int64
	Amount      int
	Type        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
