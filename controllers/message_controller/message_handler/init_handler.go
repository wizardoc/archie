package message_handler

import (
	"archie/middlewares"
	"archie/services"
	"archie/utils/jwt_utils"
	"github.com/gorilla/websocket"
	"log"
)

func InitMessageHandler(jwt string, conn *websocket.Conn) {
	claims := jwt_utils.LoginClaims{}

	if err := middlewares.ParseToken2Claims(jwt, &claims); err != nil {
		log.Println("Parse jwt fail", err)
	}

	services.Receiver.Register(claims.User.ID, conn)
}
