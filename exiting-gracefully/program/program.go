package program

import (
	"errors"
)

var (
	// ErrExitSignal receive a exit signal.
	ErrExitSignal = errors.New("program: receive a exit signal")
)

// NewExitGroup new a exit group with signals.
func NewExitGroup() *Group {
	g := new(Group)
	g.Add(func() error {
		<-RegisterExitHandlers()
		return ErrExitSignal
	}, func(err error) {
		// nothing
	})
	return g
}
