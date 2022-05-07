package gerror

// Error .
type Error struct {
	Errno  int
	ErrMsg string
}

// Error.
func (e *Error) Error() string {
	return e.ErrMsg
}

// WithLog .
func (e *Error) WithLog(str string) *Error {
	e.ErrMsg = str
	return e
}

var (
	// ActError 活动异常
	ActError = &Error{Errno: 1000, ErrMsg: "act error"}

	// LockErr 用户锁异常
	LockErr = &Error{Errno: 2001, ErrMsg: "lock error"}

	// OutDrawLimit 超出总抽奖次数
	OutDrawLimit = &Error{Errno: 3001, ErrMsg: "out draw limit"}
	// OutWinLimit 超出总中奖次数
	OutWinLimit = &Error{Errno: 3002, ErrMsg: "out win limit"}
	// OutDrawDayLimit 超出每日抽奖次数
	OutDrawDayLimit = &Error{Errno: 3003, ErrMsg: "out draw day limit"}
	// OutWinDayLimit 超出每日中奖次数
	OutWinDayLimit = &Error{Errno: 3004, ErrMsg: "out win day limit"}

	// NoStrategy 没有命中的策略
	NoStrategy = &Error{Errno: 4001, ErrMsg: "no strategy"}
	// NoRule 没有命中的规则
	NoRule = &Error{Errno: 4002, ErrMsg: "no rule"}
	// NoStack 没有命中的库存
	NoStack = &Error{Errno: 4001, ErrMsg: "no stack"}
)
