package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getLatestMessageQuery struct {
	ChannelId int `form:"channel_id" binding:"required"`
	Limit     int `form:"limit" binding:"required"`
}

func GetLatestMessageHundler(ctx *gin.Context) {
	query := getLatestMessageQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageRepository := NewMessageRepository(query.ChannelId)

	messages, err := messageRepository.GetLatestMessageByID(query.Limit)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "test"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}

type postMessageBody struct {
	UserID     int    `json:"user_id" binding:"required"`
	ChannnelId int    `json:"channnel_id" binding:"required"`
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

	messageCreateParams := CreateParams{
		UserID: requestParams.UserID,
		Type:   requestParams.Type,
		Body:   requestParams.Body,
	}

	message, err := messageRepository.Create(messageCreateParams)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
