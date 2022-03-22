package strategy

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jerome0000/draw/config"
	"github.com/jerome0000/draw/util/gerror"
	"time"
)

// Strategy 策略
type Strategy struct {
	ctx         context.Context
	redisClient *redis.Client
	drawConfig  *config.DrawConfig

	winInfo  *config.WinInfo
	userInfo map[string]int64

	reqTime time.Time

	params map[string]any

	// filter strategy
	filterStrategy []*config.Strategy
}

// New 初始化策略
func New(ctx context.Context, redisClient *redis.Client, drawConfig *config.DrawConfig, winInfo *config.WinInfo, userInfo map[string]int64, reqTime time.Time, params map[string]any) *Strategy {
	return &Strategy{
		ctx:         ctx,
		redisClient: redisClient,
		drawConfig:  drawConfig,
		winInfo:     winInfo,
		userInfo:    userInfo,
		reqTime:     reqTime,
		params:      params,
	}
}

// Handler strategy handler
func (s *Strategy) Handler(ctx context.Context) error {
	if err := s.validate(ctx); err != nil {
		return err
	}
	if err := s.run(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Strategy) validate(ctx context.Context) error {
	if len(s.drawConfig.Strategies) == 0 {
		return gerror.NoStrategy
	}
	s.filterStrategy = make([]*config.Strategy, 0)
	for _, strategy := range s.drawConfig.Strategies {
		if strategy == nil {
			continue
		}
		if strategy.StartTime.Unix() < s.reqTime.Unix() {
			continue
		}
		if strategy.EndTime.Unix() > s.reqTime.Unix() {
			continue
		}

		s.filterStrategy = append(s.filterStrategy, strategy)
	}
	if len(s.filterStrategy) == 0 {
		return gerror.NoStrategy.WithLog("strategy filtered")
	}
	return nil
}

func (s *Strategy) run(ctx context.Context) error {
	return nil
}
