package main

import (
	"context"
	"log"

	"github.com/tonybase/golang-best-practices/exiting-gracefully/app"
	"github.com/tonybase/golang-best-practices/exiting-gracefully/grpc"
	"github.com/tonybase/golang-best-practices/exiting-gracefully/http"
)

func main() {
	httpSrv := http.NewServer()
	grpcSrv := grpc.NewServer()

	a := app.New()
	a.Append(app.Hook{
		OnStart: func(ctx context.Context) error {
			return httpSrv.Start(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return httpSrv.Stop(ctx)
		},
	})
	a.Append(app.Hook{
		OnStart: func(ctx context.Context) error {
			return grpcSrv.Start(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return grpcSrv.Stop(ctx)
		},
	})
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
