package user

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	name := "Helio"

	user := Create(name)

	if user == nil {
		t.Fatalf("expected non-nil user")
	}

	if user.Name != name {
		t.Errorf("expected name %s, got %s", name, user.Name)
	}

	if user.UUID == (uuid.UUID{}) {
		t.Fatalf("expected generated UUID to be non-zero")
	}

	if user.ID != 0 {
		t.Fatalf("expected ID to be zero on creation, got %d", user.ID)
	}

	if !user.CreatedAt.IsZero() || !user.UpdatedAt.IsZero() {
		t.Fatalf("expected timestamps to be zero values before persistence")
	}
}

func TestCreateGeneratesNewUUIDEachTime(t *testing.T) {
	user1 := Create("User1")
	user2 := Create("User2")

	if user1.UUID == user2.UUID {
		t.Fatalf("expected different UUIDs for separate users")
	}
}
