package routes

import (
	"archie/controllers/message_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func messageRoutes(router *gin.Engine) {
	message := router.Group("/message")

	message.GET("/all", middlewares.ValidateToken, message_controller.GetAllMessages)
	message.GET("/connect", message_controller.ConnectWS)
}
