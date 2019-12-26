package main

import (
	"log"
	"os"

	"github.com/msales/pkg/v3/clix"
	"gopkg.in/urfave/cli.v1"

	_ "github.com/joho/godotenv/autoload"
)

const (
	flagMQTTHost      = "mqtt-host"
	flagMQTTUser      = "mqtt-user"
	flagMQTTPass      = "mqtt-pass"
	flagMQTTPort      = "mqtt-port"
	flagMQTTTopicName = "mqtt-topic-name"

	flagAuthUser = "auth-user"
	flagAuthPass = "auth-pass"

	flagMassHost = "mass-host"
	flagMassURL  = "mass-url"
)

var mqttFlags = clix.Flags{
	cli.StringFlag{
		Name:   flagMQTTHost,
		Usage:  "MQTT server host",
		EnvVar: "MQTT_HOST",
	},
	cli.StringFlag{
		Name:   flagMQTTUser,
		Usage:  "MQTT user",
		EnvVar: "MQTT_USER",
	},
	cli.StringFlag{
		Name:   flagMQTTPass,
		Usage:  "MQTT password",
		EnvVar: "MQTT_PASS",
	},
	cli.IntFlag{
		Name:   flagMQTTPort,
		Usage:  "MQTT Port",
		EnvVar: "MQTT_PORT",
	},
	cli.StringFlag{
		Name:   flagMQTTTopicName,
		Usage:  "MQTT topic name",
		EnvVar: "MQTT_TOPIC_NAME",
	},
}

var serverFlags = clix.Flags{

	cli.StringFlag{
		Name:   flagAuthUser,
		Usage:  "Base auth user",
		EnvVar: "AUTH_USER",
	},
	cli.StringFlag{
		Name:   flagAuthPass,
		Usage:  "Base auth password",
		EnvVar: "AUTH_PASS",
	},
}.Merge(mqttFlags, clix.CommonFlags, clix.ServerFlags)

var subscriberFlags = clix.Flags{
	cli.StringFlag{
		Name:   flagMassHost,
		Usage:  "MASS http api host",
		EnvVar: "MASS_HOST",
	},
	cli.StringFlag{
		Name:   flagMassURL,
		Usage:  "MASS http api url",
		EnvVar: "MASS_URL",
	},
}.Merge(mqttFlags, clix.CommonFlags, clix.ServerFlags)

// Version is the compiled application version.
var Version = "¯\\_(ツ)_/¯"

var commands = []cli.Command{
	{
		Name:   "server",
		Usage:  "Run the server",
		Flags:  serverFlags,
		Action: runServer,
	},
	{
		Name:   "subscriber",
		Usage:  "Run the subscriber",
		Flags:  subscriberFlags,
		Action: runSubscriber,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "adsrv"
	app.Flags = clix.ProfilerFlags
	app.Before = clix.RunProfiler
	app.Commands = commands
	app.Version = Version

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
