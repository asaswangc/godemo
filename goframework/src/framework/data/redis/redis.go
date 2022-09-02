package redis

import (
	"context"
	"goframework/src/framework/utils/config"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.ClusterClient

func Init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    config.Toml.Redis.Host,
		Password: config.Toml.Redis.Password,
		PoolSize: config.Toml.Redis.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Toml.Redis.PoolTimeout))
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("连接Redis数据库失败，%s", err)
	}
	RDB = client
}
