package model

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx context.Context
	r   *Redis
)

func init() {
	ctx = context.Background()

	redisClient := redis.NewClient(&redis.Options{
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

	r = New(ctx, redisClient)
}

func TestGetUserInfo(t *testing.T) {
	var result map[string]int64
	err := r.GetUserInfo(ctx, 123123, &result)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("result: %v, error: %v", result, err)
}

func TestLock(t *testing.T) {
	var err error

	err = r.Lock(ctx, 123123)
	fmt.Println(err)
	err = r.Lock(ctx, 123123)
	fmt.Println(err)

	time.Sleep(time.Second * 3)
	err = r.Lock(ctx, 123123)
	fmt.Println(err)
}

func TestUnLock(t *testing.T) {
	r.UnLock(ctx, 123123)
}
