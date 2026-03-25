package account

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type CreateInput struct {
	Name      string    `json:"name" validate:"required,min=2,max=120"`
	OwnerUUID uuid.UUID `json:"owner_uuid" validate:"required"`
}

func (i *CreateInput) Validate() error {
	return validate.Struct(i)
}
