package process

import (
	"context"
	"fmt"
	c "github.com/jerome0000/draw/conf"
	"github.com/jerome0000/draw/util"
	"sort"
	"sync"
	"time"
)

// StrategyHandler strategy handler
func StrategyHandler(ctx context.Context, reqTime time.Time, info *c.Info, conf *c.Conf, params map[string]interface{}) error {
	if len(conf.Strategies) == 0 {
		return util.NotHitStrategy
	}

	// 筛选满足条件的策略
	hitStrategies := make([]c.Strategy, 0)
	var wg sync.WaitGroup
	for _, strategy := range conf.Strategies {
		wg.Add(1)
		go func(st c.Strategy) {
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

	// 根据权重进行排序
	sort.Sort(WeightsSort(hitStrategies))
	info.StrategyInfo = hitStrategies[0]

	return nil
}

func checkStrategyStatus(strategy c.Strategy, reqTime time.Time, params map[string]interface{}) bool {
	if len(strategy.Rules) == 0 {
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

type WeightsSort []c.Strategy

func (w WeightsSort) Len() int {
	return len(w)
}
func (w WeightsSort) Swap(x, y int) {
	w[x], w[y] = w[y], w[x]
}
func (w WeightsSort) Less(x, y int) bool {
	return w[x].Weights > w[y].Weights
}
