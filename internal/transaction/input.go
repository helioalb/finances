package transaction

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type CreateInput struct {
	AccountUUID string  `json:"account_uuid"`
	Amount      int     `json:"amount" validate:"gt=0"`
	Description *string `json:"description,omitempty"`
}

func (ci *CreateInput) Validate() error {
	return validate.Struct(ci)
}
