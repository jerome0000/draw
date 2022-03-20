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
	ActError = &Error{Errno: 1000, ErrMsg: "act error"}

	LockErr = &Error{Errno: 2001, ErrMsg: "lock error"}

	OutDrawLimit    = &Error{Errno: 3001, ErrMsg: "out draw limit"}
	OutWinLimit     = &Error{Errno: 3002, ErrMsg: "out win limit"}
	OutDrawDayLimit = &Error{Errno: 3003, ErrMsg: "out draw day limit"}
	OutWinDayLimit  = &Error{Errno: 3004, ErrMsg: "out win day limit"}
)
