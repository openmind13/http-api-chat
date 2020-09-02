package model

import "time"

// Message struct
type Message struct {
	ID        int
	Chat      int // chat id
	Author    int // user id
	Text      string
	CreatedAt time.Time
}
