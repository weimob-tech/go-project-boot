package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-boot/pkg/wlog"
)

func NewRedis(ctx context.Context) (rdb *redis.Client, err error) {
	// todo: cluster, sentinel
	// todo: timeout
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.address"),
		Password: config.GetString("redis.password"), // no password set
		DB:       config.GetInt("redis.database"),    // use default DB
	})
	// test redis connection
	if "false" != config.GetString("redis.test") {
		var result string
		result, err = rdb.Ping(ctx).Result()
		if err != nil {
			panic(fmt.Errorf("redis connection test failed: %w", err))
		}
		wlog.I().Str("result", result).Msg("redis ping result")
	}

	return rdb, nil
}
