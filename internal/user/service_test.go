package user

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

type mockRepository struct {
	createFunc func(ctx context.Context, user *User) (*User, error)
}

func (m *mockRepository) Create(ctx context.Context, user *User) (*User, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, user)
	}
	return user, nil
}

func TestNewService(t *testing.T) {
	repo := &mockRepository{}
	service := newService(repo)

	if service == nil {
		t.Fatalf("expected non-nil service")
	}

	if service.repo != repo {
		t.Errorf("expected repo to be set correctly")
	}
}

func TestNewServiceWithNilRepo(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic when repo is nil")
		}
	}()

	newService(nil)
}

func TestServiceCreate(t *testing.T) {
	name := "Helio Albano"
	expectedUser := &User{
		ID:   1,
		Name: name,
		UUID: uuid.New(),
	}

	repo := &mockRepository{
		createFunc: func(ctx context.Context, user *User) (*User, error) {
			return expectedUser, nil
		},
	}

	service := newService(repo)
	createdUser, err := service.Create(context.Background(), name)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if createdUser != expectedUser {
		t.Errorf("expected returned user to match repository result")
	}
}

func TestServiceCreateWithRepositoryError(t *testing.T) {
	expectedErr := errors.New("database error")

	repo := &mockRepository{
		createFunc: func(ctx context.Context, user *User) (*User, error) {
			return nil, expectedErr
		},
	}

	service := newService(repo)
	createdUser, err := service.Create(context.Background(), "Test User")

	if createdUser != nil {
		t.Errorf("expected nil user on error")
	}

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected wrapped error, got %v", err)
	}
}

func TestServiceCreatePassesUserDataCorrectly(t *testing.T) {
	name := "Helio Albano"
	var capturedUser *User

	repo := &mockRepository{
		createFunc: func(ctx context.Context, user *User) (*User, error) {
			capturedUser = user
			return user, nil
		},
	}

	service := newService(repo)
	service.Create(context.Background(), name)

	if capturedUser.Name != name {
		t.Errorf("expected user name %s, got %s", name, capturedUser.Name)
	}

	if capturedUser.UUID == (uuid.UUID{}) {
		t.Errorf("expected UUID to be generated")
	}
}

func TestServiceCreatePreservesContext(t *testing.T) {
	var capturedCtx context.Context

	repo := &mockRepository{
		createFunc: func(ctx context.Context, user *User) (*User, error) {
			capturedCtx = ctx
			return user, nil
		},
	}

	service := newService(repo)
	testCtx := context.WithValue(context.Background(), "key", "value")
	service.Create(testCtx, "Test")

	if capturedCtx.Value("key") != "value" {
		t.Errorf("expected context to be preserved")
	}
}
