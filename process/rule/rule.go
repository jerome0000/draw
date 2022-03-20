package rule

import (
	"context"
)

// Rule 规则
type Rule struct {
	ctx context.Context
}

// Handler rule handler
func (r *Rule) Handler(ctx context.Context) error {
	return nil
}
