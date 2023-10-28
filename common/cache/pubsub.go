package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	//pub *redis.PubSub
	sub *redis.PubSub
)

func SetSubScribe(data *redis.PubSub) {
	sub = data
}

func Channel() <-chan *redis.Message {
	return sub.Channel()
}

func Publish(ctx context.Context, channel, message string) error {
	_, err := PubRedisClient.Publish(ctx, channel, message).Result()
	return err
}
