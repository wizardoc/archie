package routes

import (
	"archie/controllers/message_controller"
	"archie/middlewares"
	"github.com/gin-gonic/gin"
)

func messageRoutes(router *gin.Engine) {
	passedMessage := router.Group("/message")

	passedMessage.GET("/connect", message_controller.ConnectWS)

	message := router.Group("/message", middlewares.ValidateToken)

	message.GET("/all", message_controller.GetAllMessages)
	message.PUT("/read/:id", message_controller.ReadMessage, message_controller.ChangeMessageState)
	message.DELETE("/delete/:id", message_controller.DeleteMessage, message_controller.ChangeMessageState)
	message.PUT("/revoke/:id", message_controller.RevokeMessage, message_controller.ChangeMessageState)
	message.POST("/send", message_controller.SendMessage)
}
