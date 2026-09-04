package user

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

type mockRepository struct {
	createFn     func(ctx context.Context, user *Entity) (*Entity, error)
	getByEmailFn func(ctx context.Context, email string) (*Entity, error)
	getByUUIDFn  func(ctx context.Context, id uuid.UUID) (*Entity, error)
}

func (m *mockRepository) Create(ctx context.Context, user *Entity) (*Entity, error) {
	if m.createFn != nil {
		return m.createFn(ctx, user)
	}
	return nil, nil
}

func (m *mockRepository) GetByEmail(ctx context.Context, email string) (*Entity, error) {
	if m.getByEmailFn != nil {
		return m.getByEmailFn(ctx, email)
	}
	return nil, nil
}

func (m *mockRepository) GetByUUID(ctx context.Context, id uuid.UUID) (*Entity, error) {
	if m.getByUUIDFn != nil {
		return m.getByUUIDFn(ctx, id)
	}
	return nil, nil
}

func TestUserServiceCreate(t *testing.T) {
	someError := errors.New("database connection failed")

	t.Run("given a email already in use when creating a user then it should return ErrEmailInUse", func(t *testing.T) {
		t.Parallel()

		svc := NewService(&mockRepository{
			getByEmailFn: func(_ context.Context, _ string) (*Entity, error) {
				// Retornar entidade (não nil) e sem erro simula e-mail já existente
				return &Entity{Name: "Existente"}, nil
			},
		})
		_, err := svc.Create(context.Background(), CreateInput{
			Name:  "Fulano",
			Email: "email@ja.em.uso",
		})

		if !errors.Is(err, ErrEmailInUse) {
			t.Fatalf("expected error %v, got %v", ErrEmailInUse, err)
		}
	})

	t.Run("given a repository error when creating a user then it should bubble up the error", func(t *testing.T) {
		t.Parallel()

		svc := NewService(&mockRepository{
			getByEmailFn: func(_ context.Context, _ string) (*Entity, error) {
				return nil, ErrUserNotFound
			},
			createFn: func(_ context.Context, _ *Entity) (*Entity, error) {
				return nil, someError
			},
		})
		_, err := svc.Create(context.Background(), CreateInput{
			Name:  "Fulano",
			Email: "email@valido.com",
		})

		if !errors.Is(err, someError) {
			t.Fatalf("expected error %v, got %v", someError, err)
		}
	})

	t.Run("given a valid input when creating a user then it should succeed and return created entity", func(t *testing.T) {
		t.Parallel()

		svc := NewService(&mockRepository{
			getByEmailFn: func(_ context.Context, _ string) (*Entity, error) {
				return nil, ErrUserNotFound
			},
			createFn: func(_ context.Context, u *Entity) (*Entity, error) {
				// Valida se os dados passados para a camada de persistência estão corretos
				if u.Name != "Fulano" || u.Email != "email@valido.com" {
					t.Fatalf("create received unexpected entity fields: %+v", u)
				}
				return u, nil
			},
		})
		got, err := svc.Create(context.Background(), CreateInput{
			Name:  "Fulano",
			Email: "email@valido.com",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got == nil {
			t.Fatal("expected non-nil user entity")
		}
		if got.Name != "Fulano" {
			t.Errorf("expected name 'Fulano', got '%s'", got.Name)
		}
	})
}
