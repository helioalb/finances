package user

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type CreateInput struct {
	Name  string `json:"name" validate:"required,min=2,max=120"`
	Email string `json:"email" validate:"required,email"`
}

func (ci *CreateInput) Validate() error {
	return validate.Struct(ci)
}

func (ci *CreateInput) ToEntity() *User {
	return &User{
		Name:  ci.Name,
		Email: ci.Email,
	}
}
