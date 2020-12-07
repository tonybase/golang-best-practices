package main

import (
	"log"

	"github.com/tonybase/golang-best-practices/exiting-gracefully/app"
	"github.com/tonybase/golang-best-practices/exiting-gracefully/grpc"
	"github.com/tonybase/golang-best-practices/exiting-gracefully/http"
)

func main() {
	httpSrv := http.NewServer()
	grpcSrv := grpc.NewServer()

	a := app.New()
	a.Append(app.Hook{OnStart: httpSrv.Start, OnStop: httpSrv.Stop})
	a.Append(app.Hook{OnStart: grpcSrv.Start, OnStop: grpcSrv.Stop})

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
