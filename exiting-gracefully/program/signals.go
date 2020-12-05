package program

import (
	"os"
	"os/signal"
	"syscall"
)

var exitSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific

// RegisterExitHandlers returns a stop channel to wait on exit signals
func RegisterExitHandlers() (stop chan struct{}) {
	stop = make(chan struct{})
	sig := make(chan os.Signal, len(exitSignals))
	signal.Notify(sig, exitSignals...)

	go func() {
		// Wait for a signal from the OS before dispatching
		// a stop signal to all other goroutines observing this channel.
		<-sig
		close(stop)
	}()

	return stop
}
