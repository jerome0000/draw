package conf

import "time"

// Info 抽奖info,贯穿整个抽奖流程
type Info struct {
	StrategyInfo *Strategy // 命中的策略信息
	RuleInfo     *Rule     // 命中的规则信息
	WinPrize     *Prize    // 命中的奖品信息
}

// Conf 抽奖的配置信息
type Conf struct {
	Act        Act
	Strategies []*Strategy
	Rules      []*Rule
	Prizes     []*Prize
}

// Act 抽奖信息
type Act struct {
	StartTime      time.Time // 开始时间
	EndTime        time.Time // 结束时间
	DrawCount      int64     // 抽奖次数
	WinCount       int64     // 中奖次数
	DrawCountDaily int64     // 每日抽奖次数
	WinCountDaily  int64     // 每日中奖次数
}

// Strategy 策略信息
type Strategy struct {
	ID           int64     // 策略ID
	Rules        []int64   // 规则ID集合
	StartTime    time.Time // 策略开始时间
	EndTime      time.Time // 策略结束时间
	StartTimeDay string    // 每天开始时间(H:i:s)
	EndTimeDay   string    // 每天结束时间(H:i:s)
	Weights      int64     // 权重
	Condition    string    // 强制命中条件
	Rate         float64   // 概率
}

// Rule 规则信息
type Rule struct {
	ID           int64     // 规则ID
	Rate         float64   // 概率(此处的比例是个拟合的概率,可以认为是一个权重)
	PrizeID      int64     // 发放的奖品ID
	StartTime    time.Time // 奖品发放开始时间
	EndTime      time.Time // 奖品发放结束时间
	StartTimeDay string    // 奖品每天开始时间(H:i:s)
	EndTimeDay   string    // 奖品每天结束时间(H:i:s)
	Slice        int64     // 奖品分片
	Range        int64     // 奖品打散(单位秒)
	// 规则维度计数
	Count    int64 // 奖品总数
	CountDay int64 // 每日奖品总数
	// 该规则下用户维度计数
	UserDrawCount      int64 // 用户总抽奖个数
	UserWinCount       int64 // 用户总中奖个数
	UserDrawCountDaily int64 // 用户每日抽奖个数
	UserWinCountDaily  int64 // 用户每日中奖个数
}

// Prize 奖品信息
type Prize struct {
	ID int64 // 奖品ID
}
