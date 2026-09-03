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

func (m *mockRepository) create(ctx context.Context, user *Entity) (*Entity, error) {
	if m.createFn != nil {
		return m.createFn(ctx, user)
	}
	return nil, nil
}

func (m *mockRepository) getByEmail(ctx context.Context, email string) (*Entity, error) {
	if m.getByEmailFn != nil {
		return m.getByEmailFn(ctx, email)
	}
	return nil, nil
}

func (m *mockRepository) getByUUID(ctx context.Context, id uuid.UUID) (*Entity, error) {
	if m.getByUUIDFn != nil {
		return m.getByUUIDFn(ctx, id)
	}
	return nil, nil
}

func TestUserServiceCreate(t *testing.T) {
	someError := errors.New("database connection failed")

	tests := []struct {
		name        string
		input       CreateInput
		mockRepo    *mockRepository
		wantErr     error
		checkResult func(t *testing.T, created *Entity)
	}{
		{
			name: "given a email already in use when creating a user then it should return ErrEmailInUse",
			input: CreateInput{
				Name:  "Fulano",
				Email: "email@ja.em.uso",
			},
			mockRepo: &mockRepository{
				getByEmailFn: func(_ context.Context, _ string) (*Entity, error) {
					// Retornar entidade (não nil) e sem erro simula e-mail já existente
					return &Entity{Name: "Existente"}, nil
				},
			},
			wantErr: ErrEmailInUse,
		},
		{
			name: "given a repository error when creating a user then it should bubble up the error",
			input: CreateInput{
				Name:  "Fulano",
				Email: "email@valido.com",
			},
			mockRepo: &mockRepository{
				getByEmailFn: func(_ context.Context, _ string) (*Entity, error) {
					return nil, ErrUserNotFound
				},
				createFn: func(_ context.Context, _ *Entity) (*Entity, error) {
					return nil, someError
				},
			},
			wantErr: someError,
		},
		{
			name: "given a valid input when creating a user then it should succeed and return created entity",
			input: CreateInput{
				Name:  "Fulano",
				Email: "email@valido.com",
			},
			mockRepo: &mockRepository{
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
			},
			wantErr: nil,
			checkResult: func(t *testing.T, created *Entity) {
				if created == nil {
					t.Fatal("expected non-nil user entity")
				}
				if created.Name != "Fulano" {
					t.Errorf("expected name 'Fulano', got '%s'", created.Name)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := newService(tt.mockRepo)
			got, err := svc.Create(context.Background(), tt.input)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.checkResult != nil {
				tt.checkResult(t, got)
			}
		})
	}
}
