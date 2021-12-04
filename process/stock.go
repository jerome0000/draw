package process

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	c "github.com/jerome0000/draw/conf"
	"github.com/jerome0000/draw/util"
)

// StockHandler stock handler
func StockHandler(ctx context.Context, reqTime time.Time, uid int64, info *c.Info, conf *c.Conf, pipeline redis.Pipeliner) error {
	rule := info.RuleInfo
	if !checkRuleStatus(rule, reqTime) {
		return util.NotHitPrize
	}

	prize := getPrizeInfo(rule.PrizeID, conf)
	if prize == nil {
		return util.NotHitPrize
	}

	pipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), fmt.Sprintf("draw_stock_%v_%v", info.StrategyInfo.ID, rule.ID), 1)
	pipeline.HIncrBy(ctx, fmt.Sprintf(util.User, uid), fmt.Sprintf("draw_stock_%v_%v_%v", reqTime.Format("20060102"), info.StrategyInfo.ID, rule.ID), 1)

	// 继续检查数据

	return nil
}

func checkRuleStatus(rule *c.Rule, reqTime time.Time) bool {
	if rule.StartTime.Unix() >= reqTime.Unix() {
		return false
	}
	if rule.EndTime.Unix() <= reqTime.Unix() {
		return false
	}

	Ymd := time.Now().Format("20060102")
	startTimeDay := fmt.Sprintf("%s %s", Ymd, rule.StartTimeDay)
	endTimeDay := fmt.Sprintf("%s %s", Ymd, rule.EndTimeDay)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	startTimeDayT, err := time.ParseInLocation("20060102 15:04:05", startTimeDay, loc)
	if err != nil {
		return false
	}

	endTimeDayUnixT, err := time.ParseInLocation("20060102 15:04:05", endTimeDay, loc)
	if err != nil {
		return false
	}

	if startTimeDayT.Unix() <= reqTime.Unix() {
		return false
	}
	if endTimeDayUnixT.Unix() >= reqTime.Unix() {
		return false
	}
	return true
}

func getPrizeInfo(prizeID int64, info *c.Conf) *c.Prize {
	for _, item := range info.Prizes {
		if item.ID == prizeID {
			return item
		}
	}
	return nil
}
