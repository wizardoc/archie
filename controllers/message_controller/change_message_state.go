package message_controller

import (
	"archie/controllers/message_controller/message_state_handler"
	"archie/robust"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const SIGNAL = "message_state_signal"

const (
	READ = iota
	DELETE
	REVOKE
)

func ChangeMessageState(ctx *gin.Context) {
	res := helper.Res{}

	signal, isExist := ctx.Get(SIGNAL)
	if !isExist {
		res.Status(http.StatusNotFound).Error(ctx, robust.MESSAGE_SIGNAL_NOT_EXIST)
		return
	}

	id := ctx.Params.ByName("id")

	switch signal {
	case READ:
		go message_state_handler.ReadMessageHandler(id)
	case REVOKE:
		go message_state_handler.RevokeMessageHandler(id)
	case DELETE:
		go message_state_handler.DeleteMessageHandler(id)
	default:
		log.Println("The signal is not exist")
	}

	res.Send(ctx, nil)
}
