package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

// Hook is a pair of start and stop callbacks.
type Hook struct {
	OnStart func(context.Context) error
	OnStop  func(context.Context) error
}

// Option is a life cycle option.
type Option func(o *Options)

// Options is a life cycle options.
type Options struct {
	startTimeout time.Duration
	stopTimeout  time.Duration
	exitSignals  []os.Signal
}

// WithStartTimeout with the start timeout.
func WithStartTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.startTimeout = timeout
	}
}

// WithStopTimeout with the stop timeout.
func WithStopTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.stopTimeout = timeout
	}
}

// WithExitSignal with the exit signals.
func WithExitSignal(sigs ...os.Signal) Option {
	return func(o *Options) {
		o.exitSignals = sigs
	}
}

// App is manage the application component life cycle.
type App struct {
	opts    Options
	hooks   []Hook
	started int
}

// New new a application manage.
func New(opts ...Option) *App {
	options := Options{
		startTimeout: time.Second * 30,
		stopTimeout:  time.Second * 30,
		exitSignals: []os.Signal{
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGINT,
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return &App{opts: options}
}

// Append register callbacks that are executed on application start and stop.
func (a *App) Append(hook Hook) {
	a.hooks = append(a.hooks, hook)
}

// Run starts the application, blocks on the signals channel, and then gracefully shuts the application down.
func (a *App) Run() error {
	startCtx, startCancel := context.WithTimeout(context.Background(), a.opts.startTimeout)
	defer startCancel()
	if err := a.Start(startCtx); err != nil {
		return err
	}

	<-a.Signal()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), a.opts.stopTimeout)
	defer stopCancel()
	return a.Stop(stopCtx)
}

// Start executes all OnStart hooks registered with the application's Lifecycle.
func (a *App) Start(ctx context.Context) error {
	for _, hook := range a.hooks {
		if hook.OnStart != nil {
			if err := hook.OnStart(ctx); err != nil {
				return err
			}
		}
		a.started++
	}
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop(ctx context.Context) error {
	var errs error
	for ; a.started > 0; a.started-- {
		if hook := a.hooks[a.started-1]; hook.OnStop != nil {
			if err := hook.OnStop(ctx); err != nil {
				errs = errors.Wrapf(err, "OnStop error: %v", errs)
			}
		}
	}
	return errs
}

// Signal returns a stop channel to wait on exit signals.
func (a *App) Signal() <-chan os.Signal {
	sig := make(chan os.Signal, len(a.opts.exitSignals))
	signal.Notify(sig, a.opts.exitSignals...)
	return sig
}
