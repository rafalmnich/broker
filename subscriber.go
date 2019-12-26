package broker

import (
	"context"
	"fmt"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/json-iterator/go"
	"github.com/msales/pkg/v3/log"
)

type Subscriber interface {
	Subscribe(topic string) error
}

type MQTTSubscriber struct {
	ctx       context.Context
	client    mqtt.Client
	performer Performer
}

func NewSubscriber(ctx context.Context, c mqtt.Client, p Performer) *MQTTSubscriber {
	return &MQTTSubscriber{ctx: ctx, client: c, performer: p}
}

func (s *MQTTSubscriber) Subscribe(topic string) error {
	token := s.client.Subscribe(topic, 0, s.applyAction)

	return token.Error()
}

func (s *MQTTSubscriber) applyAction(client mqtt.Client, message mqtt.Message) {
	p := message.Payload()
	var actions Actions
	err := jsoniter.Unmarshal(p, &actions)
	if err != nil {
		return
	}

	err = s.performer.MakeActions(actions)
	if err != nil {
		log.Error(s.ctx, fmt.Sprintf("Cannot perform action: %s", err.Error()))
	}
}
