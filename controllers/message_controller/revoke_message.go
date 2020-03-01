package message_controller

import "github.com/gin-gonic/gin"

func RevokeMessage(context *gin.Context) {
	context.Set(SIGNAL, REVOKE)
	context.Next()
}
