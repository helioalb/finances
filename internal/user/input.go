package user

import "regexp"

type createInput struct {
	Name  string `json:"name" validate:"required,min=2,max=120"`
	Email string `json:"email" validate:"required,email,max=120"`
}

func (ci *createInput) ToEntity() *User {
	return &User{
		Name:  ci.Name,
		Email: ci.Email,
	}
}

func (ci *createInput) Validate() error {
	if ci.Name == "" {
		return ErrInvalidName
	}

	if len(ci.Name) < 2 || len(ci.Name) > 120 {
		return ErrInvalidName
	}

	if !isValidEmail(ci.Email) {
		return ErrInvalidEmail
	}

	return nil
}

func isValidEmail(email string) bool {
	// RFC 5322 simplified regex for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}
