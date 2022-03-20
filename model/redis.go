package model

import (
	"context"
	"fmt"
	"github.com/jerome0000/draw/util/gerror"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	// UserLock 用户锁
	// eg: user_lock_[uid]
	UserLock = "user_lock_%v"

	// User 用户计数
	// eg: user_[uid]
	User = "user_%v"
)

// Redis redis_struct
type Redis struct {
	ctx         context.Context
	redisClient *redis.Client
}

// New init redis dao
func New(ctx context.Context, client *redis.Client) *Redis {
	return &Redis{
		ctx:         ctx,
		redisClient: client,
	}
}

// Lock lock_user
func (r *Redis) Lock(ctx context.Context, uid int64) error {
	boolean, _ := r.redisClient.SetNX(ctx, fmt.Sprintf(UserLock, uid), uid, time.Second*3).Result()
	if !boolean {
		return gerror.LockErr
	}
	return nil
}

// UnLock unlock_user
func (r *Redis) UnLock(ctx context.Context, uid int64) {
	r.redisClient.Del(ctx, fmt.Sprintf(UserLock, uid))
}

// GetUserInfo 获取用户信息用户计数
func (r *Redis) GetUserInfo(ctx context.Context, uid int64, userInfo *map[string]int64) error {
	return r.redisClient.HGetAll(ctx, fmt.Sprintf(User, uid)).Scan(&userInfo)
}
