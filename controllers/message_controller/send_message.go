package message_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/robust"
	"archie/services"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendMessageParams struct {
	Title string `validate:"required" json:"title" form:"title"`
	Body  string `validate:"required" json:"body" form:"body"`
	// username
	To string `validate:"required" json:"to" form:"to"`
}

func SendMessage(ctx *gin.Context) {
	res := helper.Res{}

	var params SendMessageParams
	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	claims, err := middlewares.GetClaims(ctx)

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	// find user by username
	user := models.User{}
	if err := user.FindByUsername(params.To); err != nil {
		res.Status(http.StatusBadRequest).Error(robust.MESSAGE_CANNOT_FIND_TO).Send(ctx)
		return
	}

	if user.ID == claims.User.ID {
		res.Status(http.StatusBadRequest).Error(robust.MESSAGE_SEND_TO_YOURSELF).Send(ctx)
		return
	}

	msg := services.Message{
		Title: params.Title,
		Body:  params.Body,
		To:    user.ID,
		From:  claims.User.ID,
	}

	if err := msg.SendPersonalMessage(); err != nil {
		res.Status(http.StatusInternalServerError).Error(err).Send(ctx)
		return
	}

	res.Send(ctx)
}
