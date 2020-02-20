package routes

import (
	"archie/controllers/message_controller"
	"github.com/gin-gonic/gin"
)

func messageRoutes(router *gin.Engine) {
	message := router.Group("/message")

	message.GET("all")
	message.GET("/connect", message_controller.ConnectWS)
}
