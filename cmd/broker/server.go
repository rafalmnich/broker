package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/msales/pkg/v3/clix"
	"github.com/msales/pkg/v3/health"
	"github.com/msales/pkg/v3/httpx/middleware"
	"github.com/msales/pkg/v3/log"
	"github.com/rafalmnich/broker/server"
	"gopkg.in/urfave/cli.v1"
)

func runServer(c *cli.Context) error {
	ctx, err := clix.NewContext(c)
	if err != nil {
		panic(err)
	}
	defer ctx.Close()

	go runHealthCheck(ctx)

	publisher := newPublisher(ctx)
	mux := server.NewMux(publisher)
	mux = middleware.WithContext(ctx, mux)

	port := c.String(clix.FlagPort)

	log.Info(ctx, "Starting server on port "+port)
	s := &http.Server{Addr: ":" + port, Handler: mux, IdleTimeout: time.Second, ReadTimeout: time.Second, WriteTimeout: time.Second}
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	<-clix.WaitForSignals()

	return nil
}

func runHealthCheck(ctx *clix.Context) {
	defer closer(ctx, health.StopServer)

	err := health.StartServer(":"+ctx.String(clix.FlagPort), newHealthCheck(ctx))
	if err != nil {
		fatal(ctx, err)
	}
}

// fatal prints out the error in a panic.
func fatal(ctx context.Context, err error) {
	log.Fatal(ctx, "error", "details", fmt.Sprintf("%+v", err))
}

// closer is a helper function that executes the closeFn and logs an error if it occurs.
func closer(ctx context.Context, closeFn func() error) {
	if err := closeFn(); err != nil {
		fatal(ctx, err)
	}
}
