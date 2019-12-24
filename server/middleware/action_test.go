package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/rafalmnich/broker"
	"github.com/rafalmnich/broker/mocks"
	. "github.com/rafalmnich/broker/server/middleware"
	"github.com/stretchr/testify/assert"
)

func TestSendAction(t *testing.T) {
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
	r, _ := http.NewRequest("get", "/", strings.NewReader(string(aJson)))
	w := httptest.NewRecorder()

	p := new(mocks.Publisher)
	p.On("Publish", a).Return(nil)
	handler := SendAction(p)

	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestSendActionNoBodySent(t *testing.T) {
	r, _ := http.NewRequest("get", "/", nil)
	w := httptest.NewRecorder()

	p := new(mocks.Publisher)
	handler := SendAction(p)

	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestSendActionErorredReader(t *testing.T) {
	r, _ := http.NewRequest("get", "/", errReader(0))
	w := httptest.NewRecorder()

	p := new(mocks.Publisher)
	handler := SendAction(p)

	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSendActionBadJson(t *testing.T) {
	r, _ := http.NewRequest("get", "/", strings.NewReader("bad json"))
	w := httptest.NewRecorder()

	p := new(mocks.Publisher)
	handler := SendAction(p)

	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSendActionPublishError(t *testing.T) {
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
	r, _ := http.NewRequest("get", "/", strings.NewReader(string(aJson)))
	w := httptest.NewRecorder()

	p := new(mocks.Publisher)
	p.On("Publish", a).Return(errors.New("test error"))
	handler := SendAction(p)

	handler.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
