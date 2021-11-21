package channel

import "time"

type Channel struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Owner     uint      `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Channels []*Channel
