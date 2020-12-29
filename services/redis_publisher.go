package services

import (
	"archie/connection/redis_conn"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
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

	redis_conn.GetRedisConnMust(func(conn redis.Conn) error {
		if err := conn.Send("PUBLISH", NOTIFY_CHANNEL, data); err != nil {
			conn.Close()

			return err
		}

		if err := conn.Flush(); err != nil {
			return err
		}

		return nil
	})

	return nil
}
