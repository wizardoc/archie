package message_controller

import "github.com/gin-gonic/gin"

func RevokeMessage(ctx *gin.Context) {
	ctx.Set(SIGNAL, REVOKE)
	ctx.Next()
}
