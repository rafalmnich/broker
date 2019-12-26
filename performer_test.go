package broker_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/dghubble/sling"
	"github.com/rafalmnich/broker"
	"github.com/stretchr/testify/assert"
)

func Test_performer_MakeActions(t *testing.T) {
	client := mockClient(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(nil)),
		}, nil
	})

	p := broker.NewHTTPPerformer(client, "/actions")

	actions := broker.Actions{
		{
			Name:  "foo",
			Value: 10,
		},
		{
			Name:  "bar",
			Value: 300,
		},
	}
	err := p.MakeActions(actions)
	assert.NoError(t, err)
}

func Test_performer_MakeActions_errored(t *testing.T) {
	client := mockClient(func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("test error")
	})

	p := broker.NewHTTPPerformer(client, "/actions")

	actions := broker.Actions{
		{
			Name:  "foo",
			Value: 10,
		},
		{
			Name:  "bar",
			Value: 300,
		},
	}
	err := p.MakeActions(actions)
	assert.Error(t, err)
}

func mockClient(fn MockDoer) *sling.Sling {
	return sling.New().Doer(fn)
}

type MockDoer func(req *http.Request) (*http.Response, error)

func (fn MockDoer) Do(req *http.Request) (*http.Response, error) {
	return fn(req)
}
