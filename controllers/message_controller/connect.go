package message_controller

import (
	"archie/controllers/message_controller/message_handler"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	INIT = iota
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	readMessageChan = make(chan BaseMessage, 2)
)

type BaseMessage struct {
	Type    int
	Payload interface{}
}

func ConnectWS(context *gin.Context) {
	//authRes := helper.Res{Status: http.StatusUnauthorized}

	//parsedClaims, err := middlewares.GetClaims(context)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	authRes.Err = err
	//	authRes.Send(context)
	//	return
	//}

	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		log.Println("Unable to upgrade websocket", err)
		return
	}

	//services.Receiver.Register(parsedClaims.UserId, conn)
	go messageDispatcher(conn)
	go readMessage(conn)
}

func readMessage(conn *websocket.Conn) {
	baseMessage := BaseMessage{}

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			if err := conn.Close(); err != nil {
				log.Println("close the conn fail", err)
				return
			}

			log.Println(err)
			return
		}

		// decode baseMessage
		if err := json.Unmarshal(msg, &baseMessage); err != nil {
			log.Println("Unexpect BaseMessage")
			continue
		}

		readMessageChan <- baseMessage
	}
}

func messageDispatcher(conn *websocket.Conn) {
	for msg := range readMessageChan {
		switch msg.Type {
		case INIT:
			message_handler.InitMessageHandler(msg.Payload.(string), conn)
		}
	}
}
