package model

import "time"

// Message struct
type Message struct {
	ID        int
	Chat      *Chat
	Author    *User
	Text      string
	CreatedAt time.Time
}
