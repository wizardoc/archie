package message_state_handler

import (
	"archie/models"
	"log"
)

func ReadMessageHandler(id string) {
	message := models.Message{
		ID:     id,
		IsRead: true,
	}

	if err := message.Update(); err != nil {
		log.Println("Read message fail", err)
	}
}
