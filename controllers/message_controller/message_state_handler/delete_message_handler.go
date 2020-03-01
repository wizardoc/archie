package message_state_handler

import (
	"archie/models"
	"log"
)

func DeleteMessageHandler(id string) {
	message := models.Message{
		ID:       id,
		IsDelete: true,
	}

	if err := message.Update(); err != nil {
		log.Println("Delete message fail", err)
	}
}
