package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "event-store-pump"
	app.Usage = "..."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{}

	app.Run(os.Args)
}
