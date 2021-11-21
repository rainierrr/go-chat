package message

import "time"

type Message struct {
	ID        int       `json:"id"`
	ChannelID int       `json:"channel_id"`
	UserID    int       `json:"user_id"`
	Type      string    `json:"type"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Messages []*Message
