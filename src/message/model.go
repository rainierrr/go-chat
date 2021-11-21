package message

import "time"

type Message struct {
	ID        uint      `json:"id"`
	ChannelID uint      `json:"channel_id"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Messages []*Message
