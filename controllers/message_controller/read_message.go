package message_controller

import "github.com/gin-gonic/gin"

func ReadMessage(context *gin.Context) {
	context.Set(SIGNAL, READ)
	context.Next()
}
