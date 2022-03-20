package process

import "context"

// IProcess .
type IProcess interface {
	Handler(ctx context.Context) error
}

// IHandler .
type IHandler interface {
	Validate() error
	Run() error
}
