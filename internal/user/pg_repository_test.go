package user

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestNewPgRepository(t *testing.T) {
	t.Run("panic when db is nil", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic when creating repository with nil db")
			}
		}()

		newPgRepository(nil)
	})

	t.Run("success when db is not nil", func(t *testing.T) {
		mockDB := &sql.DB{}
		repo := newPgRepository(mockDB)

		if repo == nil {
			t.Error("Expected repository to not be nil")
		}

		if repo.db != mockDB {
			t.Error("Expected repository.db to equal provided db")
		}
	})
}

func TestRepository_Create(t *testing.T) {
	t.Run("should return error when user is nil", func(t *testing.T) {
		mockDB := &sql.DB{}
		repo := newPgRepository(mockDB)
		ctx := context.Background()

		_, err := repo.Create(ctx, nil)
		if err == nil {
			t.Error("Expected error when creating user with nil user")
		}
	})

	t.Run("should create user successfully", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create sqlmock: %v", err)
		}
		defer db.Close()

		uu := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
		expectedUser := &User{
			ID:        1,
			UUID:      uu,
			Name:      "Test User",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(uu, "Test User").
			WillReturnRows(sqlmock.NewRows([]string{"id", "uuid", "name", "created_at", "updated_at"}).
				AddRow(expectedUser.ID, expectedUser.UUID, expectedUser.Name, expectedUser.CreatedAt, expectedUser.UpdatedAt))

		repo := newPgRepository(db)
		ctx := context.Background()
		userToCreate := &User{
			UUID: uu,
			Name: "Test User",
		}

		createdUser, err := repo.Create(ctx, userToCreate)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if createdUser.ID != expectedUser.ID || createdUser.UUID != expectedUser.UUID || createdUser.Name != expectedUser.Name {
			t.Errorf("Created user does not match expected user. Got %+v, expected %+v", createdUser, expectedUser)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %v", err)
		}
	})
}
