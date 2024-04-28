package main

import (
	"net/http"
	"server/listeners"
	"server/routes"
	"time"

	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	main_router := routes.SetupRouter()
	main_router.Run() // default listens on :8080

	println("Base server is running on port 8080")

	stdHttpListener := &http.Server{
		Addr:         ":42069",
		Handler:      listeners.StdHttpRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	torHttpListener := &http.Server{
		Addr:         ":42070",
		Handler:      listeners.TorHttpRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Run the tor listener
	g.Go(func() error {
		return stdHttpListener.ListenAndServe()
	})

	// Run the standard listener

	g.Go(func() error {
		return torHttpListener.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		println(err.Error())
	}

}
