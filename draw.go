package draw

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// IDraw draw interface
type IDraw interface {
	Do(ctx context.Context, redisClient redis.Client, params map[string]interface{}) (interface{}, error)
}

// Draw draw_struct
type Draw struct {
}

// Do do_draw
func (d *Draw) Do(ctx context.Context, redisClient redis.Client, params map[string]interface{}) (interface{}, error) {

	// format draw conf

	// check user draw/win all/day limit

	// do strategy handler

	// do rule handler

	// do stock handler

	return nil, nil
}