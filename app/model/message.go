package model

import "time"

// Message struct
type Message struct {
	ID        int
	Chat      int // chat_id
	Author    int // user_id
	Text      string
	CreatedAt time.Time
}
