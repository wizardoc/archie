package message_handler

import (
	"archie/middlewares"
	"archie/services"
	"github.com/gorilla/websocket"
	"log"
)

func InitMessageHandler(jwt string, conn *websocket.Conn) {
	claims, err := middlewares.ParseToken2Claims(jwt)

	if err != nil {
		log.Println("Parse jwt fail", err)
	}

	services.Receiver.Register(claims.UserId, conn)
}
