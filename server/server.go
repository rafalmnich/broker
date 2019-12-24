package server

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/msales/pkg/v3/health"
	"github.com/rafalmnich/broker"
	"github.com/rafalmnich/broker/server/middleware"
)

const (
	username = "iqccuser"
	password = "ihe*(#(DUjoas389rfj"
)

type Server interface {
	health.Reporter
}

func NewMux(p broker.Publisher) http.Handler {
	mux := bone.New()

	mux.Post("/action", ActionHandler(p))

	return mux
}

func ActionHandler(p broker.Publisher) http.Handler {
	// innermost
	h := middleware.SendAction(p)
	h = middleware.BasicAuth(h, username, password, "Please enter your username and password for this site")
	// outermost
	return h
}
