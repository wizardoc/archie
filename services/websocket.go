package services

import (
	"archie/models"
	"archie/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type ParsedMessage struct {
	models.Message
	From models.User `json:"from"`
}

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
	fmt.Println("add conn")
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
	var conns []*websocket.Conn

	//utils.ArrayMap(userIDs, func(id interface{}) interface{} {
	//	return pool.conns[id.(string)]
	//}, &conns)

	for _, userID := range userIDs {
		fmt.Println(pool.conns, userID)
		conn := pool.conns[userID]

		if conn == nil {
			continue
		}

		conns = append(conns, conn)
	}

	fmt.Println(conns)

	return sendMsgMulti(conns, cm)
}

// broadcast to all connections
func (pool *WebsocketPool) Broadcast(cm *ChannelMessage) error {
	return sendMsgMulti(pool.Conns(), cm)
}

func sendMsgMulti(conns []*websocket.Conn, cm *ChannelMessage) error {
	m, err := persistentMessage(cm)

	if err != nil {
		return err
	}

	user := models.User{}
	if err := user.Find("id", m.From); err != nil {
		return err
	}

	pm := ParsedMessage{
		Message: *m,
		From:    user,
	}

	// marshal msg
	msg, err := json.Marshal(pm)

	if err != nil {
		return err
	}

	for _, conn := range conns {
		sendMsg(conn, msg)
	}

	return nil
}

func persistentMessage(cm *ChannelMessage) (*models.Message, error) {
	m := models.Message{}
	utils.CpStruct(cm, &m)

	main, err := json.Marshal(cm.Main)

	if err != nil {
		return nil, err
	}

	m.Main = string(main)

	if err := m.Create(cm.To); err != nil {
		return nil, err
	}

	return &m, nil
}

func sendMsg(conn *websocket.Conn, msg []byte) {
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println("send message error", err)

		if err == websocket.ErrCloseSent {
			if err := conn.Close(); err != nil {
				log.Println("close websocket connection error", err)
			}
		}
	}
}
