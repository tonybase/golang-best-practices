package program

type worker struct {
	start    func() error
	shutdown func(error)
}

// Group is graceful shutdown of worker goroutines.
type Group struct {
	w []worker
}

// Add adds worker funcs to the group.
func (g *Group) Add(start func() error, shutdown func(error)) {
	g.w = append(g.w, worker{start: start, shutdown: shutdown})
}

// Wait blocks until the Group returns shutdown.
func (g *Group) Wait() error {
	if len(g.w) == 0 {
		return nil
	}

	errors := make(chan error, len(g.w))
	for _, w := range g.w {
		go func(w worker) {
			errors <- w.start()
		}(w)
	}

	err := <-errors

	for _, w := range g.w {
		w.shutdown(err)
	}

	for len(errors) > 0 {
		<-errors
	}
	return err
}
