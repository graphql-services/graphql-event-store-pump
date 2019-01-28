package cmd

import (
	"log"

	"github.com/graphql-services/graphql-event-store-pump/src"
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
		},
		Action: func(c *cli.Context) error {
			aggregatorURL := c.String("aggregator-url")
			if aggregatorURL == "" {
				log.Fatal("Missing AGGREGATOR_URL variable")
			}

			bootupOptions := src.PerformBootupOptions{AggregatorURL: aggregatorURL}
			if err := src.PerformBootup(bootupOptions); err != nil {
				return cli.NewExitError(err, 1)
			}

			return nil
		},
	}
}
