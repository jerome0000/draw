package stock

import (
	"context"
)

// Stock 库存
type Stock struct {
	ctx context.Context
}

// Handler stock handler
func (s *Stock) Handler(ctx context.Context) error {
	return nil
}
