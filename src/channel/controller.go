package channel

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChannelHundler(ctx *gin.Context) {
	channelsRepository := NewChannelRepository()

	channels, err := channelsRepository.GetAll()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "test"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"channels": channels})
}

type postcChannelBody struct {
	Name  string `json:"name" binding:"required"`
	Owner uint   `json:"owner" binding:"required"`
}

func PostChannelHundler(ctx *gin.Context) {
	body := postcChannelBody{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channelRepository := NewChannelRepository()

	createParams := CreateParams{
		Name:  body.Name,
		Owner: body.Owner,
	}

	channel, err := channelRepository.Create(createParams)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "test"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"channel": channel})
}
