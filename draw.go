package draw

import (
	"context"
	"fmt"
	"github.com/jerome0000/draw/process/strategy"
	"github.com/jerome0000/draw/util/gerror"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jerome0000/draw/config"
	"github.com/jerome0000/draw/model"
)

// IDraw draw interface
type IDraw interface {
	Do(ctx context.Context, uid int64, params map[string]any) (*config.WinInfo, error)
}

// Draw draw_struct
type Draw struct {
	ctx         context.Context
	redisClient *redis.Client
	drawConfig  *config.DrawConfig

	winInfo  *config.WinInfo
	userInfo map[string]int64

	reqTime time.Time
	redisM  *model.Redis
}

// New init draw
func New(ctx context.Context, redisClient *redis.Client, drawConfig *config.DrawConfig) *Draw {
	return &Draw{
		ctx:         ctx,
		redisClient: redisClient,
		drawConfig:  drawConfig,

		winInfo: &config.WinInfo{},

		redisM: model.New(ctx, redisClient),
	}
}

// Do 抽奖
func (d *Draw) Do(ctx context.Context, uid int64, params map[string]any) (info *config.WinInfo, err error) {
	d.reqTime = time.Now()
	rand.Seed(d.reqTime.UnixNano())

	if err = d.validate(); err != nil {
		return
	}

	if err = d.redisM.Lock(ctx, uid); err != nil {
		return
	}
	defer d.redisM.UnLock(ctx, uid)

	if err = d.redisM.GetUserInfo(ctx, uid, &d.userInfo); err != nil {
		return
	}

	if err = d.checkUserLimit(); err != nil {
		return
	}

	redisPipeline := d.redisClient.Pipeline()
	defer redisPipeline.Exec(ctx)

	redisPipeline.HIncrBy(ctx, fmt.Sprintf(model.User, uid), "draw", 1)
	redisPipeline.HIncrBy(ctx, fmt.Sprintf(model.User, uid), fmt.Sprintf("draw_%s", d.reqTime.Format("20060102")), 1)

	// 策略
	err = strategy.New(ctx, d.redisClient, d.drawConfig, d.winInfo, d.userInfo, d.reqTime, params).Handler(ctx)
	if err != nil {
		return
	}
	// 规则
	// 库存

	redisPipeline.HIncrBy(ctx, fmt.Sprintf(model.User, uid), "win", 1)
	redisPipeline.HIncrBy(ctx, fmt.Sprintf(model.User, uid), fmt.Sprintf("win_%s", d.reqTime.Format("20060102")), 1)

	return
}

// validate 基础检验
func (d *Draw) validate() error {
	if d.redisClient == nil {
		return gerror.ActError.WithLog("redis nil")
	}
	if d.drawConfig == nil {
		return gerror.ActError.WithLog("draw config error")
	}
	if d.drawConfig.Act.StartTime.Unix() <= d.reqTime.Unix() {
		return gerror.ActError.WithLog("act not start")
	}
	if d.drawConfig.Act.EndTime.Unix() >= d.reqTime.Unix() {
		return gerror.ActError.WithLog("act had end")
	}
	return nil
}

// checkUserLimit 检查用户状态
func (d *Draw) checkUserLimit() error {
	// 抽奖次数
	drawCount, _ := d.userInfo["draw"]
	drawCountDaily, _ := d.userInfo[fmt.Sprintf("draw_%s", d.reqTime.Format("20060102"))]
	// 中奖次数
	winCount, _ := d.userInfo["win"]
	winCountDaily, _ := d.userInfo[fmt.Sprintf("win_%s", d.reqTime.Format("20060102"))]

	if drawCount >= d.drawConfig.Act.DrawCount {
		return gerror.OutDrawLimit
	}
	if winCount >= d.drawConfig.Act.WinCount {
		return gerror.OutWinLimit
	}
	if drawCountDaily >= d.drawConfig.Act.DrawCountDaily {
		return gerror.OutDrawDayLimit
	}
	if winCountDaily >= d.drawConfig.Act.WinCountDaily {
		return gerror.OutWinDayLimit
	}
	return nil
}
