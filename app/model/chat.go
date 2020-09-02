package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Lengths of chat name
const (
	ChatNameMinLength = 4
	ChatNameMaxLength = 20
)

// Chat struct
type Chat struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Users     []string `json:"users"`
	CreatedAt time.Time
}

// Validate ...
func (chat *Chat) Validate() error {
	return validation.ValidateStruct(
		chat,
		validation.Field(
			&chat.Name,
			validation.Required,
			validation.Length(ChatNameMinLength, ChatNameMaxLength),
		),
		// validation.Field(
		// 	&chat.Users,
		// 	validation.Each(validation.By())
		// )
	)
}
