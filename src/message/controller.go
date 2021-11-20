package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postMessageBody struct {
	UserID     uint   `json:"user_id" binding:"required"`
	ChannnelId uint   `json:"channnel_id" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Body       string `json:"body" binding:"required"`
}

func PostMessageHundler(ctx *gin.Context) {

	requestParams := postMessageBody{}
	if err := ctx.ShouldBindJSON(&requestParams); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageRepository := NewMessageRepository(requestParams.ChannnelId)

	messageCreateParams := MessageCreateParams{
		UserID: requestParams.UserID,
		Type:   requestParams.Type,
		Body:   requestParams.Body,
	}

	message, err := messageRepository.Create(messageCreateParams)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
