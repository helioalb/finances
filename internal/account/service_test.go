package account

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/helioalb/finances/internal/user"
)

type mockRepository struct {
	createFn func(ctx context.Context, account *Entity) (*Entity, error)
}

func (m *mockRepository) Create(ctx context.Context, account *Entity) (*Entity, error) {
	if m.createFn != nil {
		return m.createFn(ctx, account)
	}
	return nil, nil
}

func (m *mockRepository) GetByOwnerUUIDAndName(context.Context, uuid.UUID, string) (*Entity, error) {
	return nil, ErrAccountNotFound
}

type mockUserService struct {
	getByUUIDFn func(ctx context.Context, id uuid.UUID) (*user.Entity, error)
}

func (m *mockUserService) GetByUUID(ctx context.Context, id uuid.UUID) (*user.Entity, error) {
	return m.getByUUIDFn(ctx, id)
}

func TestServiceCreate(t *testing.T) {
	t.Parallel()

	ownerUUID := uuid.New()
	ownerID := int64(42)

	t.Run("creates an account for an existing user", func(t *testing.T) {
		t.Parallel()

		service := NewService(
			&mockRepository{
				createFn: func(_ context.Context, account *Entity) (*Entity, error) {
					if account.UserID != ownerID || account.Name != "Checking" {
						t.Fatalf("unexpected account: %+v", account)
					}
					return account, nil
				},
			},
			&mockUserService{
				getByUUIDFn: func(_ context.Context, id uuid.UUID) (*user.Entity, error) {
					if id != ownerUUID {
						t.Fatalf("unexpected owner UUID: %s", id)
					}
					return &user.Entity{ID: ownerID}, nil
				},
			},
		)

		created, err := service.Create(context.Background(), CreateInput{Name: "Checking", OwnerUUID: ownerUUID})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if created == nil || created.UserID != ownerID {
			t.Fatalf("unexpected created account: %+v", created)
		}
	})

	t.Run("propagates a missing owner", func(t *testing.T) {
		t.Parallel()

		service := NewService(
			&mockRepository{},
			&mockUserService{getByUUIDFn: func(context.Context, uuid.UUID) (*user.Entity, error) {
				return nil, user.ErrUserNotFound
			}},
		)

		_, err := service.Create(context.Background(), CreateInput{OwnerUUID: ownerUUID})
		if !errors.Is(err, user.ErrUserNotFound) {
			t.Fatalf("expected %v, got %v", user.ErrUserNotFound, err)
		}
	})

	t.Run("propagates account creation errors", func(t *testing.T) {
		t.Parallel()

		someError := errors.New("database connection failed")
		service := NewService(
			&mockRepository{createFn: func(context.Context, *Entity) (*Entity, error) {
				return nil, someError
			}},
			&mockUserService{getByUUIDFn: func(context.Context, uuid.UUID) (*user.Entity, error) {
				return &user.Entity{ID: ownerID}, nil
			}},
		)

		_, err := service.Create(context.Background(), CreateInput{OwnerUUID: ownerUUID})
		if !errors.Is(err, someError) {
			t.Fatalf("expected %v, got %v", someError, err)
		}
	})

	t.Run("propagates duplicate account errors", func(t *testing.T) {
		t.Parallel()

		service := NewService(
			&mockRepository{createFn: func(context.Context, *Entity) (*Entity, error) {
				return nil, ErrAccountAlreadyExists
			}},
			&mockUserService{getByUUIDFn: func(context.Context, uuid.UUID) (*user.Entity, error) {
				return &user.Entity{ID: ownerID}, nil
			}},
		)

		_, err := service.Create(context.Background(), CreateInput{OwnerUUID: ownerUUID})
		if !errors.Is(err, ErrAccountAlreadyExists) {
			t.Fatalf("expected %v, got %v", ErrAccountAlreadyExists, err)
		}
	})
}
