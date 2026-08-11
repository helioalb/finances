package user

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

type mockRepository struct {
	createFn     func(ctx context.Context, user *Entity) (*Entity, error)
	getByEmailFn func(ctx context.Context, email string) (*Entity, error)
	getByUUIDFn  func(ctx context.Context, uuid uuid.UUID) (*Entity, error)
}

func (m *mockRepository) create(ctx context.Context, user *Entity) (*Entity, error) {
	if m.createFn != nil {
		return m.createFn(ctx, user)
	}
	return nil, nil
}

func (m *mockRepository) getByEmail(ctx context.Context, email string) (*Entity, error) {
	return m.getByEmailFn(ctx, email)
}

func (m *mockRepository) getByUUID(ctx context.Context, uuid uuid.UUID) (*Entity, error) {
	return m.getByUUIDFn(ctx, uuid)
}

func TestUserServiceCreate(t *testing.T) {
	t.Run("email already in use", func(t *testing.T) {
		repo := &mockRepository{
			getByEmailFn: func(_ context.Context, _ string) (*Entity, error) {
				return nil, nil // err nil means the email already exists
			},
		}

		svc := newService(repo)

		input := CreateInput{
			Name:  "Fulano",
			Email: "email@ja.em.uso",
		}

		_, err := svc.Create(context.Background(), input)

		fmt.Println(err)
		if !errors.Is(err, errEmailInUse) {
			t.Errorf("Expected errEmailInUse")
		}
	})

	t.Run("fail", func(t *testing.T) {
		someError := errors.New("some error")

		repo := &mockRepository{
			getByEmailFn: func(ctx context.Context, email string) (*Entity, error) {
				return nil, errUserNotFound
			},
			createFn: func(ctx context.Context, user *Entity) (*Entity, error) {
				return nil, someError
			},
		}

		svc := newService(repo)
		input := CreateInput{
			Name:  "Fulano",
			Email: "email@valido.com",
		}

		_, err := svc.Create(context.Background(), input)

		if !errors.Is(err, someError) {
			t.Errorf("Expected repository error")
		}
	})

	t.Run("success", func(t *testing.T) {
		repo := &mockRepository{
			getByEmailFn: func(ctx context.Context, email string) (*Entity, error) {
				return nil, errUserNotFound
			},
			createFn: func(ctx context.Context, user *Entity) (*Entity, error) {
				return user, nil
			},
		}

		svc := newService(repo)
		input := CreateInput{
			Name:  "Fulano",
			Email: "email@valido.com",
		}

		_, err := svc.Create(context.Background(), input)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
