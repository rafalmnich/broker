package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/rafalmnich/broker"
	"github.com/rafalmnich/broker/mocks"
	"github.com/rafalmnich/broker/server"
	"github.com/stretchr/testify/assert"
)

var serverOpts = server.Opts{
	Username: "iqccuser",
	Password: "ihe*((DUjoas389rfj",
}

func TestNewMux(t *testing.T) {
	p := new(mocks.Publisher)

	mux := server.NewMux(p, serverOpts)

	a := broker.Actions{
		{
			Name:  "foo",
			Value: 10,
		},
		{
			Name:  "bar",
			Value: 0,
		},
	}

	aJson, err := jsoniter.Marshal(a)
	assert.NoError(t, err)

	p.On("Publish", a).Return(nil)

	r, _ := http.NewRequest("POST", "/action", strings.NewReader(string(aJson)))
	r.SetBasicAuth("iqccuser", "ihe*((DUjoas389rfj")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
