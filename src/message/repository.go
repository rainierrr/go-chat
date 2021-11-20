package message

import "github.com/rainierrr/go-chat/db"

type Repository interface {
	Create(params MessageCreateParams) (*Message, error)
}

func NewMessageRepository(channelID uint) Repository {
	return &MessageRepository{ChannelID: channelID}
}

type MessageRepository struct {
	ChannelID uint
}

type MessageCreateParams struct {
	UserID uint
	Type   string
	Body   string
}

func (r MessageRepository) Create(params MessageCreateParams) (*Message, error) {
	message := Message{
		UserID:    params.UserID,
		ChannelID: r.ChannelID,
		Type:      params.Type,
		Body:      params.Body,
	}

	if result := db.DB.Create(&message); result.Error != nil {
		return &message, result.Error
	}

	return &message, nil
}
