package services

import (
	"archie/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

// websocket connection pool that is used to push some data into the connection
// and provide CURD interface for this connection pool
type WebsocketPool struct {
	conns map[string]*websocket.Conn
}

func (pool *WebsocketPool) RemoveConn(userID string) {
	delete(pool.conns, userID)
}

func (pool *WebsocketPool) AddConn(userID string, targetConn *websocket.Conn) {
	pool.conns[userID] = targetConn
}

func (pool *WebsocketPool) Conns() []*websocket.Conn {
	var values []*websocket.Conn

	for _, conn := range pool.conns {
		values = append(values, conn)
	}

	return values
}

// send msg to a specify user or multi users
func (pool *WebsocketPool) SendDirectionalMsg(cm *ChannelMessage, userIDs ...string) error {
	conns := make([]*websocket.Conn, len(userIDs))

	utils.ArrayMap(userIDs, func(item interface{}) interface{} {
		return pool.conns[item.(string)]
	}, &conns)

	return sendMsgMulti(conns, cm)
}

// broadcast to all connections
func (pool *WebsocketPool) Broadcast(cm *ChannelMessage) error {
	return sendMsgMulti(pool.Conns(), cm)
}

func sendMsgMulti(conns []*websocket.Conn, cm *ChannelMessage) error {
	fmt.Println(conns)
	// marshal msg
	msg, err := json.Marshal(cm)

	if len(conns) == 1 {
		sendMsg(conns[0], msg)
		return nil
	}

	if err != nil {
		return err
	}

	for _, conn := range conns {
		sendMsg(conn, msg)
	}

	return nil
}

func sendMsg(conn *websocket.Conn, msg []byte) {
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println("send message error", err)

		if err := conn.Close(); err != nil {
			log.Println("close websocket connection error", err)
		}
	}
}
