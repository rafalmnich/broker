package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/msales/pkg/v3/clix"
	"github.com/msales/pkg/v3/health"
	"github.com/rafalmnich/broker"
)

func newHealthCheck(ctx *clix.Context) health.ReporterFunc {
	return func() error {
		return nil //todo: extend health check
	}
}

func newPublisher(ctx *clix.Context) broker.Publisher {
	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%s:%d", ctx.String(flagMQTTHost), ctx.Int(flagMQTTPort))).
		SetUsername(ctx.String(flagMQTTUser)).
		SetPassword(ctx.String(flagMQTTPass))

	client := connect(opts)
	topic := ctx.String(flagMQTTTopicName)
	return broker.NewPublisher(client, topic)
}

func connect(opts *mqtt.ClientOptions) mqtt.Client {
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
