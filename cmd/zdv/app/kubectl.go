package app

import (
	"github.com/kzdv/cli/pkg/kubectl"
	"github.com/urfave/cli/v2"
)

func kubectlCommand() *cli.Command {
	return &cli.Command{
		Name:            "kubectl",
		Usage:           "kubectl command line tool for accessing the kubernetes cluster and automation",
		SkipFlagParsing: true,
		Action:          kubectl.Run,
	}
}
