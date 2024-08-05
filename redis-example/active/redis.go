package active

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

// https://www.liwenzhou.com/posts/Go/redis/
// 初始化全局
var Rdb *redis.Client
var ctx = context.Background()

func Init() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       2,
		PoolSize: 100,
	})
	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	slog.Info("redis start...")
	return err
}
func Close() {
	_ = Rdb.Close()
}
