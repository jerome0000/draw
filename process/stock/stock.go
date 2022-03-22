package stock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jerome0000/draw/config"
	"github.com/jerome0000/draw/model"
	"time"
)

// Stock 库存
type Stock struct {
	ctx         context.Context
	redisClient *redis.Client
	drawConfig  *config.DrawConfig

	winInfo  *config.WinInfo
	userInfo map[string]int64

	reqTime time.Time
	redisM  *model.Redis

	params map[string]any
}

// Handler stock handler
func (s *Stock) Handler(ctx context.Context) error {
	if err := s.validate(ctx); err != nil {
		return err
	}
	if err := s.run(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Stock) validate(ctx context.Context) error {
	return nil
}

func (s *Stock) run(ctx context.Context) error {
	return nil
}
