package main

import (
	"github.com/gin-gonic/gin"
	eventapp "github.com/waro163/feishu_robot/event-app"
	_ "github.com/waro163/feishu_robot/settings"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	event := router.Group("/api/event")
	eventapp.RegisterRouter(event)

	router.Run(":8080")
}
