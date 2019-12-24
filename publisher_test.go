package broker_test

import (
	"testing"

	. "github.com/rafalmnich/broker"
	"github.com/rafalmnich/broker/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPublisher_Publish(t *testing.T) {
	topic := "topic"
	client := new(mocks.Client)

	actions := Actions{
		{
			Name:  "foo",
			Value: 100,
		},
		{
			Name:  "bar",
			Value: 0,
		},
	}

	msg := []uint8(`[{"name":"foo","value":100},{"name":"bar","value":0}]`)
	token := new(mocks.Token)
	token.On("Error").Return(nil)
	client.On("Publish", "topic", uint8(0), false, msg).Return(token)

	p := NewPublisher(client, topic)

	err := p.Publish(actions)
	assert.NoError(t, err)
}
