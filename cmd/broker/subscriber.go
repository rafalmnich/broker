package main

import (
	"github.com/msales/pkg/v3/clix"
	"gopkg.in/urfave/cli.v1"
)

func runSubscriber(c *cli.Context) error {
	ctx, err := clix.NewContext(c)
	if err != nil {
		panic(err)
	}
	defer ctx.Close()

	subscriber := newSubscriber(ctx)
	subscriber.Subscribe(ctx.String(flagMQTTTopicName))

	<-clix.WaitForSignals()

	return nil
}
