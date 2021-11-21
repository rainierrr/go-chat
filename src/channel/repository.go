package channel

import (
	"github.com/rainierrr/go-chat/db"
)

type Repository interface {
	//GetByID() (*Channels, error)
	Create(params CreateParams) (*Channel, error)
}
type ChannelRepository struct{}

func NewChannelRepository() Repository {
	return &ChannelRepository{}
}

type CreateParams struct {
	Name  string
	Owner uint
}

func (r ChannelRepository) Create(params CreateParams) (*Channel, error) {
	channel := Channel{
		Name:  params.Name,
		Owner: params.Owner,
	}

	if result := db.DB.Create(&channel); result.Error != nil {
		return &channel, result.Error
	}

	return &channel, nil
}
