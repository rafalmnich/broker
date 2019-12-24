package main

import (
	"github.com/msales/pkg/v3/clix"
	"github.com/msales/pkg/v3/health"
	"github.com/rafalmnich/broker"
)

func newHealthCheck(ctx *clix.Context) health.ReporterFunc {
	return func() error {
		return nil //todo: extend health check
	}
}

func newPublisher(ctx *clix.Context) broker.Publisher {

	return nil
}
