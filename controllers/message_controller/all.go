package message_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/services"
	"archie/utils"
	"archie/utils/helper"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ParsedMessage struct {
	models.Message
	Main services.ChannelMessageMain `json:"main"`
	From models.User                 `json:"from"`
}

type GetAllMessageParams struct {
	utils.PageInfo
}

// Get all messages by user
func GetAllMessages(ctx *gin.Context) {
	res := helper.Res{}

	var params GetAllMessageParams
	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	params.ParsePageInfo()

	// parse JWT
	claims, err := middlewares.GetClaims(ctx)

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	user := models.User{ID: claims.User.ID}

	// cannot find all messages
	if err := user.FindAllMessages(params.Page, params.PageSize); err != nil {
		res.Status(http.StatusNotFound).Error(err).Send(ctx)
		return
	}

	notifies := []ParsedMessage{}
	chats := []ParsedMessage{}

	// fill user info
	var fromIDs []string
	froms := make(map[string]models.User)

	utils.ArrayMap(user.Messages, func(item interface{}) interface{} {
		return item.(models.Message).From
	}, &fromIDs)

	// no message
	if len(froms) == 0 {
		res.Success(gin.H{
			"notifies": notifies,
			"chats":    chats,
		}).Send(ctx)
		return
	}

	if err := models.FindAllUsersByFrom(froms, fromIDs); err != nil {
		res.Status(http.StatusInternalServerError).Error(err).Send(ctx)

		return
	}

	for _, m := range user.Messages {
		main := services.ChannelMessageMain{}

		if err := json.Unmarshal([]byte(m.Main), &main); err != nil {
			log.Println("unmarshal main of message fail", err)
			continue
		}

		parsedMsg := ParsedMessage{
			Message: m,
			Main:    main,
			From:    froms[m.From],
		}

		// dispatch notify messages or chat messages
		if m.MessageType == services.NOTIFY {
			notifies = append(notifies, parsedMsg)
		} else {
			chats = append(chats, parsedMsg)
		}
	}

	res.Success(gin.H{
		"notifies": notifies,
		"chats":    chats,
	}).Send(ctx)
}
