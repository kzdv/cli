package app

import (
	"fmt"

	"github.com/kzdv/cli/pkg/version"
	"github.com/urfave/cli/v2"
	"hawton.dev/log4g"
)

var log = log4g.Category("app")

func NewRootCommand() *cli.App {
	app := &cli.App{
		Name:  "zdv",
		Usage: "zdv command line tool for accessing the kubernetes cluster and automation",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "enable debug mode",
				Aliases: []string{"d"},
			},
		},
		Commands: []*cli.Command{
			loginCommand(),
			versionCommand(),
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				log4g.SetLogLevel(log4g.DEBUG)
			}

			return nil
		},
	}

	return app
}

func versionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Print the version",
		Action: func(c *cli.Context) error {
			fmt.Printf("%s\n", version.FriendlyVersion())
			return nil
		},
	}
}
