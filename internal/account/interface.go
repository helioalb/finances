package account

import "context"

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Entity, error)
}
