package broker

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/json-iterator/go"
)

// Action is the action message sent through http api
type Action struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Actions is for many actions
type Actions []*Action

// Publisher is an Action publishing interface
type Publisher interface {
	Publish(actions Actions) error
}

type publisher struct {
	client mqtt.Client
	topic  string
}

// NewPublisher creates a default MQTT Publisher instance
func NewPublisher(client mqtt.Client, topic string) Publisher {
	return &publisher{client: client, topic: topic}
}

// Publish publishes the actions to mqtt topic
func (p publisher) Publish(actions Actions) error {
	// error is not possible here, as actions is well defined
	msg, _ := jsoniter.Marshal(actions)

	return p.client.Publish(p.topic, 0, false, msg).Error()
}
