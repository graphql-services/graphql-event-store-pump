package main

import (
	"os"

	"github.com/graphql-services/graphql-event-store-pump/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "event-store-pump"
	app.Usage = "..."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		cmd.BootCmd(),
		cmd.StartCmd(),
	}

	app.Run(os.Args)
}
