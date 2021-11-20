package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rainierrr/go-chat/db"
	"github.com/rainierrr/go-chat/message"
)

func main() {
	router := gin.Default()

	if err := db.Initialize(); err != nil {
		log.Fatalf("Failed To Initialize DB : %v", err)
	}

	router.POST("/message", message.PostMessageHundler)
	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
