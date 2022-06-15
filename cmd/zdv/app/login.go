package app

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func loginCommand() *cli.Command {
	return &cli.Command{
		Name:  "login",
		Usage: "Login to ZDV's OAuth2 Provider",
		Action: func(c *cli.Context) error {
			fmt.Printf("login command\n")
			return nil
		},
	}
}
