package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/rafalmnich/broker/server/middleware"
	"github.com/stretchr/testify/assert"
)

func TestBasicAuthBlocking(t *testing.T) {
	r, err := http.NewRequest("get", "/", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	h := func(w http.ResponseWriter, r *http.Request) {

	}
	user := "username"
	pass := "password"

	auth := BasicAuth(h, user, pass, "gimme password or u die")
	auth.ServeHTTP(w, r)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
func TestBasicAuthPass(t *testing.T) {
	r, err := http.NewRequest("get", "username:password@/", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	h := func(w http.ResponseWriter, r *http.Request) {

	}
	user := "username"
	pass := "password"

	r.SetBasicAuth(user, pass)
	auth := BasicAuth(h, user, pass, "gimme password or u die")
	auth.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
