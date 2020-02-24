package services

import (
	"archie/connection/redis_conn"
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

	redis_conn.GetRedisConnMust(func(conn redis.Conn) {
		if err := conn.Send("PUBLISH", NOTIFY_CHANNEL, data); err != nil {
			log.Println(err)
			conn.Close()
		}

		if err := conn.Flush(); err != nil {
			log.Println(err)
		}
	})

	return nil
}
