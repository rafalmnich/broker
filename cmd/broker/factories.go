package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dghubble/sling"
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
	client := connect(prepareMqttOpts(ctx))
	topic := ctx.String(flagMQTTTopicName)

	return broker.NewPublisher(client, topic)
}

func prepareMqttOpts(ctx *clix.Context) *mqtt.ClientOptions {
	return mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%s:%d", ctx.String(flagMQTTHost), ctx.Int(flagMQTTPort))).
		SetUsername(ctx.String(flagMQTTUser)).
		SetPassword(ctx.String(flagMQTTPass))
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

func newSubscriber(ctx *clix.Context) broker.Subscriber {
	client := connect(prepareMqttOpts(ctx))

	slingClient := newSlingClient(ctx)
	performer := broker.NewHTTPPerformer(slingClient, ctx.String(flagMassURL))

	return broker.NewSubscriber(ctx, client, performer)

}

func newSlingClient(ctx *clix.Context) *sling.Sling {
	doer := sling.New().Base(ctx.String(flagMassHost))

	return doer
}
