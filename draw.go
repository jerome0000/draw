package draw

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jerome0000/draw/process"
	"github.com/jerome0000/draw/util"
	"time"
)

// IDraw draw interface
type IDraw interface {
	Do(ctx context.Context, redisClient *redis.Client, conf *Conf, uid int64, params map[string]interface{}) (*Info, error)
}

// Draw draw_struct
type Draw struct {
}

// Do do_draw
func (d *Draw) Do(ctx context.Context, redisClient *redis.Client, conf *Conf, uid int64, params map[string]interface{}) (info *Info, err error) {
	reqTime := time.Now()

	var userInfo util.UserInfo

	if err = checkCommonStatus(redisClient, conf); err != nil {
		return
	}

	if err = checkActStatus(conf.Act, reqTime); err != nil {
		return
	}

	if err = util.Lock(ctx, redisClient, uid); err != nil {
		return
	}
	defer util.UnLock(ctx, redisClient, uid)

	if userInfo, err = util.GetUserInfo(ctx, redisClient, uid); err != nil {
		return
	}

	if err = checkUserLimit(userInfo, conf.Act); err != nil {
		return
	}

	redisPipeline := redisClient.Pipeline()
	defer redisPipeline.Exec(ctx)

	redisPipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), "draw", 1)
	redisPipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), "draw_daily", 1)

	if err = process.StrategyHandler(ctx, reqTime, info, conf); err != nil {
		return
	}

	if err = process.RuleHandler(); err != nil {
		return
	}

	if err = process.StockHandler(); err != nil {
		return
	}

	redisPipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), "win", 1)
	redisPipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), "win_daily", 1)

	return
}

// checkCommonStatus 检查常规参数
func checkCommonStatus(redisClient *redis.Client, conf *Conf) error {
	if redisClient == nil {
		return util.RedisNil
	}
	if conf == nil {
		return util.ConfNil
	}
	return nil
}

// checkAct 检查活动状态
func checkActStatus(act Act, time time.Time) error {
	if act.StartTime.Unix() <= time.Unix() {
		return util.ActError
	}
	if act.EndTime.Unix() >= time.Unix() {
		return util.ActError
	}
	return nil
}

// checkUserLimit 检查用户状态
func checkUserLimit(userInfo util.UserInfo, act Act) error {
	if userInfo.Draw >= act.DrawCount {
		return util.OutDrawLimit
	}
	if userInfo.Win >= act.WinCount {
		return util.OutWinLimit
	}
	if userInfo.DrawDaily >= act.DrawCountDaily {
		return util.OutDrawDayLimit
	}
	if userInfo.WinDaily >= act.WinCountDaily {
		return util.OutWinDayLimit
	}
	return nil
}
