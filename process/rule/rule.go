package rule

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jerome0000/draw/config"
	"github.com/jerome0000/draw/model"
	"time"
)

// Rule 规则
type Rule struct {
	ctx         context.Context
	redisClient *redis.Client
	drawConfig  *config.DrawConfig

	winInfo  *config.WinInfo
	userInfo map[string]int64

	reqTime time.Time
	redisM  *model.Redis

	params map[string]any
}

// Handler rule handler
func (r *Rule) Handler(ctx context.Context) error {
	if err := r.validate(ctx); err != nil {
		return err
	}
	if err := r.run(ctx); err != nil {
		return err
	}
	return nil
}

func (r *Rule) validate(ctx context.Context) error {
	return nil
}

func (r *Rule) run(ctx context.Context) error {
	return nil
}
