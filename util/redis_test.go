package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var (
	redisClient *redis.Client
	ctx         context.Context
)

func init() {
	ctx = context.Background()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if redisClient == nil {
		panic("redis client nil")
	}

	result, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(fmt.Sprintf("result: %v, err: %v", result, err))
}

func TestGetUserInfo(t *testing.T) {
	userInfo, err := GetUserInfo(ctx, redisClient, 123123)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("result: %v, error: %v", userInfo, err)
}

func TestLock(t *testing.T) {
	var err error

	err = Lock(ctx, redisClient, 123123)
	fmt.Println(err)
	err = Lock(ctx, redisClient, 123123)
	fmt.Println(err)

	time.Sleep(time.Second * 3)
	err = Lock(ctx, redisClient, 123123)
	fmt.Println(err)
}

func TestUnLock(t *testing.T) {
	UnLock(ctx, redisClient, 123123)
}
