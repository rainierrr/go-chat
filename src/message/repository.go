package message

import "github.com/rainierrr/go-chat/db"

type Repository interface {
	GetLatestMessageByID(linit uint) (*Messages, error)
	Create(params CreateParams) (*Message, error)
}
type MessageRepository struct {
	ChannelID uint
}

func NewMessageRepository(channelID uint) Repository {
	return &MessageRepository{ChannelID: channelID}
}

func (r MessageRepository) GetLatestMessageByID(linit uint) (*Messages, error) {
	messages := Messages{}

	if result := db.DB.Where(&Message{ChannelID: r.ChannelID}).Limit(int(linit)).Find(&messages); result.Error != nil {
		return &messages, result.Error
	}

	return &messages, nil
}

type CreateParams struct {
	UserID uint
	Type   string
	Body   string
}

func (r MessageRepository) Create(params CreateParams) (*Message, error) {
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
