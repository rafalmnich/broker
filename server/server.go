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
	password = "ihe*((DUjoas389rfj"
)

// Server is the server interface
type Server interface {
	health.Reporter
}

// Opts are base auth options for server
type Opts struct {
	Username string
	Password string
}

// NewMux returns a http handler with action endpoint
func NewMux(p broker.Publisher, opts Opts) http.Handler {
	mux := bone.New()

	mux.Post("/action", actionHandler(p, opts))

	return mux
}

func actionHandler(p broker.Publisher, opts Opts) http.Handler {
	// innermost
	h := middleware.SendAction(p)
	h = middleware.BasicAuth(h, opts.Username, opts.Password, "Please enter your username and password for this site")
	// outermost
	return h
}
