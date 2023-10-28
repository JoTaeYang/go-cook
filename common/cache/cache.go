package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	// WriteRedisClient ...
	WriteRedisClient *redis.ClusterClient
	// ReadRedisClient ...
	ReadRedisClient *redis.ClusterClient

	// pub
	PubRedisClient *redis.ClusterClient

	// sub
	SubRedisClient *redis.ClusterClient
)

const appName = "macovill.develop"

type Config struct {
	PAddr []string `yaml:"primary_addr"`
}

func InitRedis(conf *Config, pubSub ...bool) {
	WriteRedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: conf.PAddr,
	})
	if WriteRedisClient == nil {
		log.Println("redis connect fail")
		return
	}

	ReadRedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: conf.PAddr,
	})
	if ReadRedisClient == nil {
		log.Println("redis connect fail")
		return
	}

	log.Println("redis connect success")

	log.Println(ReadRedisClient.Ping(context.Background()))

	if pubSub != nil {
		PubRedisClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: conf.PAddr,
		})
		if PubRedisClient == nil {
			log.Println("redis connect fail")
			return
		}
		log.Println("redis pubsub ready")
	}

	log.Println("redis connect success")
}
