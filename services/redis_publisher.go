package services

import (
	"archie/connection"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
)

type RedisPublisher struct {
	Message ChannelMessage
}

func NewPublisher(message *ChannelMessage) RedisPublisher {
	publisher := RedisPublisher{
		Message: *message,
	}

	return publisher
}

func (publisher RedisPublisher) Publish() error {
	data, err := json.Marshal(publisher.Message)

	if err != nil {
		return err
	}

	connection.GetRedisConnMust(func(conn redis.Conn) {
		if err := conn.Send("PUBLISH", NOTIFY_CHANNEL, data); err != nil {
			log.Println(err)
		}

		if err := conn.Flush(); err != nil {
			log.Println(err)
		}
	})

	return nil
}
