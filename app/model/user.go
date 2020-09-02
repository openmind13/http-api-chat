package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Legth of usernames
const (
	UsernameMinLength = 5
	UsernameMaxLength = 20
)

// User struct
type User struct {
	ID        int
	Username  string
	CreatedAt time.Time
}

// Validate ...
func (user *User) Validate() error {
	return validation.ValidateStruct(
		user,
		validation.Field(
			&user.Username,
			validation.Required,
			validation.Length(UsernameMinLength, UsernameMaxLength)),
	)
}
