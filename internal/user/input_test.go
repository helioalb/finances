package user_test

import (
	"testing"

	"github.com/helioalb/finances/internal/user"
)

func TestCreateInputValidate(t *testing.T) {
	t.Run("given a valid input when validated then it should be valid", func(t *testing.T) {
		input := &user.CreateInput{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		if err := input.Validate(); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("given an invalid input when validated then it should return an error", func(t *testing.T) {
		input := &user.CreateInput{
			Name:  "J",
			Email: "invalid-email",
		}

		if err := input.Validate(); err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}

func TestCreateInputToEntity(t *testing.T) {
	t.Run("given a valid input when converted to entity then it should return a valid entity", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "john.doe@example.com"

		input := &user.CreateInput{
			Name:  expectedName,
			Email: expectedEmail,
		}

		entity := input.ToEntity()

		if entity.Name != expectedName || entity.Email != expectedEmail {
			t.Errorf("Expected a valid entity, got %v", entity)
		}
	})
}
