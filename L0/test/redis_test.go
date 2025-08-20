package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"main.go/utils"
	"testing"
)

func createRedis(t *testing.T) *redis.Client {
	t.Helper()
	err := utils.NewRedisConfig("../.data/dev.yaml")
	require.NoError(t, err)

	red := redis.NewClient(&redis.Options{
		Addr:        utils.RedisCFG.RedisAddr,
		DB:          utils.RedisCFG.RedisDB,
		MaxRetries:  utils.RedisCFG.MaxRetries,
		DialTimeout: utils.RedisCFG.DialTimeout,
		PoolTimeout: utils.RedisCFG.Timeout,
		IdleTimeout: utils.RedisCFG.Timeout,
	})
	return red
}

func TestRedisConn(t *testing.T) {
	red := createRedis(t)
	stat := red.Ping(context.Background())
	require.NotEmpty(t, stat)

	fmt.Println(stat)
}

func TestRedisSet(t *testing.T) {
	red := createRedis(t)
	stat := red.Set(context.Background(), "some-order-uid", "sdfsdfsdf", utils.RedisCFG.TTL)
	require.NotEmpty(t, stat)

	fmt.Println(stat)
}
