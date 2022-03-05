package process

import (
	"context"
	"github.com/jerome0000/draw/config"
	"github.com/jerome0000/draw/util"
	"math/rand"
)

// RuleHandler rule handler
func RuleHandler(ctx context.Context, info *config.DrawInfo, conf *config.DrawConfig) error {
	// 检查规则配置
	if len(conf.Rules) == 0 {
		return util.NotHitRule
	}

	// 命中的全部规则
	strategy := info.StrategyInfo
	rules := strategy.Rules
	if len(rules) == 0 {
		return util.NotHitRule
	}

	// 有效的全部规则
	allRules := make([]*config.Rule, 0)
	for _, item := range conf.Rules {
		if util.Int64InArray(item.ID, rules) {
			allRules = append(allRules, item)
		}
	}
	if allRules == nil || len(allRules) == 0 {
		return util.NotHitRule
	}

	// 获取命中的规则
	rule := getRuleByRate(allRules)
	if rule == nil {
		return util.NotHitRule
	}

	info.RuleInfo = rule
	return nil
}

func getRuleByRate(rules []*config.Rule) *config.Rule {
	rate := 0.0
	for _, item := range rules {
		rate += item.Rate
	}

	hitRate := rand.Float64() * rate

	for _, item := range rules {
		hitRate = hitRate - item.Rate
		if hitRate <= 0 {
			return item
		}
	}

	return nil
}
