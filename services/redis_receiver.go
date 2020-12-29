package services

import (
	"archie/connection/redis_conn"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"log"
)

const (
	NOTIFY_CHANNEL = "notify_channel"
)

var (
	Receiver = NewRedisReceiver()
)

type RedisReceiver struct {
	//pool *redis.Pool
	pool     WebsocketPool
	messages chan *ChannelMessage
	new      chan WebsocketConnPayload
	del      chan WebsocketConnPayload
}

type WebsocketConnPayload struct {
	UserID string
	conn   *websocket.Conn
}

func NewRedisReceiver() RedisReceiver {
	return RedisReceiver{
		pool: WebsocketPool{
			conns: make(map[string]*websocket.Conn),
		},
		messages: make(chan *ChannelMessage, 1000),
		new:      make(chan WebsocketConnPayload),
		del:      make(chan WebsocketConnPayload),
	}
}

// initialize redis subscribe and wait msg from provider
// process msg, for example, broadcast to all user that cache in WebsocketPool
// and send msg to a specify user or multi users
func (receiver *RedisReceiver) Run() {
	redis_conn.GetRedisConnMust(func(conn redis.Conn) error {
		psConn := redis.PubSubConn{Conn: conn}
		if err := psConn.
			Subscribe(NOTIFY_CHANNEL); err != nil {
			log.Fatal(err)
		}

		go receiver.messageDispatch()

		for {
			switch t := psConn.Receive().(type) {
			case redis.Message:
				receiver.pushMessage(t.Data)
			case redis.Subscription:
				log.Printf("Redis subscription received, %s %d %s", t.Kind, t.Count, t.Channel)
			case error:
			default:
				log.Println("Unknown type of Receive during subscription")
			}
		}
	})
}

// establish ChannelMessage and leaving data to the messageDispatch for further processing
// The message is come from anywhere that you can get UserID
func (receiver *RedisReceiver) pushMessage(msg []byte) {
	channelMessage := ChannelMessage{}

	if err := json.Unmarshal(msg, &channelMessage); err != nil {
		log.Fatal("Invalid ChannelMessage")
	}

	receiver.messages <- &channelMessage
}

// dispatch handler center
func (receiver *RedisReceiver) messageDispatch() {
	for {
		select {
		case msg := <-receiver.messages:
			processMsg(receiver.pool, msg)
		case payload := <-receiver.new:
			receiver.pool.AddConn(payload.UserID, payload.conn)
		case payload := <-receiver.del:
			receiver.pool.RemoveConn(payload.UserID)
		}
	}
}

// put connection of websocket in conn pool
func (receiver *RedisReceiver) Register(userID string, conn *websocket.Conn) {
	receiver.pool.AddConn(userID, conn)
}

// process message that from redis channel
func processMsg(pool WebsocketPool, msg *ChannelMessage) {
	if err := msg.valid(); err != nil {
		log.Println(err)
	}

	logMust := func(err error) {
		if err != nil {
			log.Printf("Send msg to %s fail", err)
		}
	}

	switch msg.Type {
	case BROADCAST:
		logMust(pool.Broadcast(msg))
	case DIRECTIONAL:
		logMust(pool.SendDirectionalMsg(msg, msg.To...))
	default:
		log.Println("Unknown type of msg during subscription")
	}
}
