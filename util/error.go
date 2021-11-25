package util

import "errors"

var (
	// OutDrawLimit 超过总抽奖次数
	OutDrawLimit = errors.New("out draw limit")
	// OutWinLimit 超过总中奖次数
	OutWinLimit = errors.New("out win limit")
	// OutDrawDayLimit 超过每天抽奖次数
	OutDrawDayLimit = errors.New("out draw day limit")
	// OutWinDayLimit 超过每天中奖次数
	OutWinDayLimit = errors.New("out win day limit")

	// NotHitStrategy 没有命中策略
	NotHitStrategy = errors.New("not hit strategy")
	// NotHitRule 没有命中规则
	NotHitRule = errors.New("not hit rule")

	// OutCountLimit 奖品数量不足
	OutCountLimit = errors.New("out count limit")
	// OutCycleCountLimit 周期内奖品数量不足
	OutCycleCountLimit = errors.New("out cycle count limit")
)
