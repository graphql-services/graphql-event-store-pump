package cmd

import (
	"github.com/urfave/cli"
)

// BootCmd ...
func BootCmd() cli.Command {
	return cli.Command{
		Name:        "boot",
		Description: "perform boot check and run import if needed",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "aggregator-url",
				EnvVar: "AGGREGATOR_URL",
			},
			cli.StringFlag{
				Name:   "eventstore-url",
				EnvVar: "EVENTSTORE_URL",
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
