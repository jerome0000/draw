package draw

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/jerome0000/draw/conf"
	"github.com/jerome0000/draw/process"
	"github.com/jerome0000/draw/util"
)

// IDraw draw interface
type IDraw interface {
	Do(ctx context.Context, redisClient *redis.Client, conf *conf.Conf, uid int64, params map[string]interface{}) (*conf.Info, error)
}

// Draw draw_struct
type Draw struct {
}

// Do do_draw
func (d *Draw) Do(ctx context.Context, redisClient *redis.Client, conf *conf.Conf, uid int64, params map[string]interface{}) (info *conf.Info, err error) {
	reqTime := time.Now()
	reqData := reqTime.Format("20060102")

	// 随机种子
	rand.Seed(reqTime.UnixNano())

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
	redisPipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), fmt.Sprintf("draw_%s", reqData), 1)

	if err = process.StrategyHandler(ctx, reqTime, info, conf, params); err != nil {
		return
	}

	if err = process.RuleHandler(ctx, info, conf); err != nil {
		return
	}

	if err = process.StockHandler(ctx, reqTime, uid, info, conf, redisPipeline); err != nil {
		return
	}

	redisPipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), "win", 1)
	redisPipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), fmt.Sprintf("win_%s", reqData), 1)

	return
}

// checkCommonStatus 检查常规参数
func checkCommonStatus(redisClient *redis.Client, conf *conf.Conf) error {
	if redisClient == nil {
		return util.RedisNil
	}
	if conf == nil {
		return util.ConfNil
	}
	return nil
}

// checkAct 检查活动状态
func checkActStatus(act conf.Act, time time.Time) error {
	if act.StartTime.Unix() <= time.Unix() {
		return util.ActError
	}
	if act.EndTime.Unix() >= time.Unix() {
		return util.ActError
	}
	return nil
}

// checkUserLimit 检查用户状态
func checkUserLimit(userInfo util.UserInfo, act conf.Act) error {
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
