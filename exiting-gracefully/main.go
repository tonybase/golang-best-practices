package main

import (
	"context"
	"log"
	"time"

	"github.com/tonybase/golang-best-practices/exiting-gracefully/app"
	"github.com/tonybase/golang-best-practices/exiting-gracefully/http"
)

func main() {
	srv := http.NewServer()

	a := app.New()
	a.Append(app.Hook{
		OnStart: func(ctx context.Context) error {
			return srv.Start(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	if err := a.Run(); err != nil {
		log.Println(err)
	}

	time.Sleep(time.Second * 5)
}
