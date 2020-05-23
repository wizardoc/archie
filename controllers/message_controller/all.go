package message_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/services"
	"archie/utils"
	"archie/utils/helper"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type ParsedMessage struct {
	models.Message
	Main services.ChannelMessageMain `json:"main"`
	From models.User                 `json:"from"`
}

// Get all messages by user
func GetAllMessages(ctx *gin.Context) {
	res := helper.Res{}

	// parse JWT
	claims, err := middlewares.GetClaims(ctx)

	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	user := models.User{ID: claims.UserId}

	// cannot find all messages
	if err := user.FindAllMessages(); err != nil && !gorm.IsRecordNotFoundError(err) {
		res.Status(http.StatusNotFound).Error(ctx, err)
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

	if err := models.FindAllUsersByFrom(froms, fromIDs); err != nil {
		res.Status(http.StatusInternalServerError).Error(ctx, err)

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

	res.Send(ctx, gin.H{
		"notifies": notifies,
		"chats":    chats,
	})
}
