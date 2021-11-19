package main

import (
	"log"
	"net/http"

	"ghe.rakuten-it.com/dmc/jwt_issuer/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := db.Initialize(); err != nil {
		log.Fatalf("Failed To Initialize DB : %v", err)
	}

	router.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
