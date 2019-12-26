package broker_test

import (
	"context"
	"testing"

	. "github.com/rafalmnich/broker"
	"github.com/rafalmnich/broker/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_subscriber_Subscribe(t *testing.T) {
	c := new(mocks.Client)
	p := new(mocks.Performer)
	token := new(mocks.Token)

	topic := "topic"

	ctx := context.Background()

	token.On("Error").Return(nil)

	c.On("Subscribe", topic, uint8(0), mock.Anything).Return(token)

	s := NewSubscriber(ctx, c, p)
	err := s.Subscribe(topic)
	assert.NoError(t, err)
}
