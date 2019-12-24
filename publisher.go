package broker

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/json-iterator/go"
)

type Action struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Actions []*Action

type Publisher interface {
	Publish(actions Actions) error
}

type publisher struct {
	client mqtt.Client
	topic  string
}

func NewPublisher(client mqtt.Client, topic string) Publisher {
	return &publisher{client: client, topic: topic}
}

func (p publisher) Publish(actions Actions) error {
	// error is not possible here, as actions is well defined
	msg, _ := jsoniter.Marshal(actions)

	return p.client.Publish(p.topic, 0, false, msg).Error()
}
