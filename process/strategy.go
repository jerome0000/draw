package process

import (
	"context"
	"github.com/jerome0000/draw"
	"github.com/jerome0000/draw/util"
	"sync"
	"time"
)

// StrategyHandler strategy handler
func StrategyHandler(ctx context.Context, reqTime time.Time, info *draw.Info, conf *draw.Conf) error {
	if len(conf.Strategies) == 0 {
		return util.NotHitStrategy
	}

	// 筛选满足条件的策略
	hitStrategies := make([]draw.Strategy, 0)
	var wg sync.WaitGroup
	for _, strategy := range conf.Strategies {
		wg.Add(1)
		go func(st draw.Strategy) {
			defer wg.Done()
			if checkStrategyStatus(st, reqTime) {
				hitStrategies = append(hitStrategies, st)
			}
		}(strategy)
	}
	wg.Wait()

	// 排除获得能够命中的策略

	return nil
}

func checkStrategyStatus(strategy draw.Strategy, reqTime time.Time) bool {
	if strategy.StartTime.Unix() <= reqTime.Unix() {
		return false
	}
	if strategy.EndTime.Unix() >= reqTime.Unix() {
		return false
	}
	if len(strategy.Rules) == 0 {
		return false
	}
	return true
}
