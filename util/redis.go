package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// UserLock 用户锁
// eg: user_lock_[uid]
const UserLock = "user_lock_%v"

// User 用户计数
// eg: user_[uid]
const User = "user_%v"

// Rule 规则计数
// eg: rule_[strategyID]_[ruleID]_[slice]_[cycle]
const Rule = "rule_%v_%v_%v_%v"

// Lock lock_user
func Lock(ctx context.Context, redisClient *redis.Client, uid int64) error {
	boolean, _ := redisClient.SetXX(ctx, fmt.Sprintf(UserLock, uid), uid, time.Second*3).Result()
	if !boolean {
		return LockUserError
	}
	return nil
}

// UnLock unlock_user
func UnLock(ctx context.Context, redisClient *redis.Client, uid int64) {
	redisClient.Del(ctx, fmt.Sprintf(UserLock, uid))
}

// UserInfo 用户信息用户计数
type UserInfo struct {
	Draw      int64 `redis:"draw"`
	Win       int64 `redis:"win"`
	DrawDaily int64 `redis:"draw_daily"`
	WinDaily  int64 `redis:"win_daily"`
}

// GetUserInfo 获取用户信息用户计数
func GetUserInfo(ctx context.Context, redisClient *redis.Client, uid int64) (userInfo UserInfo, err error) {
	err = redisClient.HGetAll(ctx, fmt.Sprintf(User, uid)).Scan(&userInfo)
	return
}
