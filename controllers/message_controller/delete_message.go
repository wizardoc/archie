package message_controller

import "github.com/gin-gonic/gin"

func DeleteMessage(ctx *gin.Context) {
	ctx.Set(SIGNAL, DELETE)
	ctx.Next()
}
