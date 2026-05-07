package cache

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/richxcame/gosms/src/utils"
)

func NewCacheClient(ctx context.Context) (*redis.Client, error) {
	redisDb := utils.GetEnvD("REDIS_DB", "0")
	database, err := strconv.Atoi(redisDb)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.GetEnvD("REDIS_ADDR", "localhost:6379"),
		Password: utils.GetEnvD("REDIS_PASS", ""), // no password set
		DB:       database,                        // use default DB
	})

	s := rdb.Ping(ctx)
	if s.Err() != nil {
		return nil, s.Err()
	}

	return rdb, nil
}
