package redis_cli

import (
	"context"
	"goframework/src/framework/utils/cfg"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisConnect *redis.ClusterClient

func Init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.T.Redis.Host,
		Password: cfg.T.Redis.Password,
		PoolSize: cfg.T.Redis.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.T.Redis.PoolTimeout))
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("连接Redis数据库失败，%s", err)
	}
	RedisConnect = client
}
