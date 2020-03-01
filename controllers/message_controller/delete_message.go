package message_controller

import "github.com/gin-gonic/gin"

func DeleteMessage(context *gin.Context) {
	context.Set(SIGNAL, DELETE)
	context.Next()
}
