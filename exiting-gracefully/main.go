package main

import (
	"golang-best-practices/exiting-gracefully/http"
	"golang-best-practices/exiting-gracefully/program"
)

func main() {
	g := program.NewExitGroup()

	srv := http.NewServer()
	g.Add(srv.Start, srv.Shutdown)
	g.Wait()
}
