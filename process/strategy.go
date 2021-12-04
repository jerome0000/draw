package process

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

	c "github.com/jerome0000/draw/conf"
	"github.com/jerome0000/draw/util"
)

// StrategyHandler strategy handler
func StrategyHandler(ctx context.Context, reqTime time.Time, info *c.Info, conf *c.Conf, params map[string]interface{}) error {
	// 检查策略配置
	if len(conf.Strategies) == 0 {
		return util.NotHitStrategy
	}

	// 筛选满足条件的策略
	hitStrategies := make([]*c.Strategy, 0)
	var wg sync.WaitGroup
	for _, strategy := range conf.Strategies {
		wg.Add(1)
		go func(st *c.Strategy) {
			defer wg.Done()
			if checkStrategyStatus(st, reqTime, params) {
				hitStrategies = append(hitStrategies, st)
			}
		}(strategy)
	}
	wg.Wait()

	if len(hitStrategies) == 0 {
		return util.NotHitStrategy
	}

	// 初始化默认概率
	initRate := rand.Float64() * 100

	// 根据强制命中条件进行选择
	conditionStrategy := checkCondition(hitStrategies, params)
	if conditionStrategy != nil {
		// 判断概率
		if conditionStrategy.Rate >= initRate {
			return util.NotHitStrategy
		}
		info.StrategyInfo = conditionStrategy
		return nil
	}

	// 根据概率命中条件进行选择
	sort.Sort(WeightsSort(hitStrategies))
	weightsStrategy := hitStrategies[0]
	if weightsStrategy.Rate >= initRate {
		return util.NotHitStrategy
	}
	info.StrategyInfo = weightsStrategy
	return nil
}

func checkStrategyStatus(strategy *c.Strategy, reqTime time.Time, params map[string]interface{}) bool {
	if strategy == nil || len(strategy.Rules) == 0 {
		return false
	}
	if strategy.StartTime.Unix() <= reqTime.Unix() {
		return false
	}
	if strategy.EndTime.Unix() >= reqTime.Unix() {
		return false
	}

	Ymd := time.Now().Format("20060102")
	startTimeDay := fmt.Sprintf("%s %s", Ymd, strategy.StartTimeDay)
	endTimeDay := fmt.Sprintf("%s %s", Ymd, strategy.EndTimeDay)

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

func checkCondition(strategies []*c.Strategy, params map[string]interface{}) *c.Strategy {
	// todo 强制命中补充
	return nil
}

type WeightsSort []*c.Strategy

func (w WeightsSort) Len() int {
	return len(w)
}
func (w WeightsSort) Swap(x, y int) {
	w[x], w[y] = w[y], w[x]
}
func (w WeightsSort) Less(x, y int) bool {
	return w[x].Weights > w[y].Weights
}
