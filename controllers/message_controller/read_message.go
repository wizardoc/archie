package message_controller

import "github.com/gin-gonic/gin"

func ReadMessage(ctx *gin.Context) {
	ctx.Set(SIGNAL, READ)
	ctx.Next()
}
