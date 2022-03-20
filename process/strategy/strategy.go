package strategy

import (
	"context"
)

// Strategy 策略
type Strategy struct {
	ctx context.Context
}

// Handler strategy handler
func (s *Strategy) Handler(ctx context.Context) error {
	return nil
}
