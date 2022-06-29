package app

import (
	"fmt"

	"github.com/kzdv/cli/pkg/config"
	"github.com/kzdv/cli/pkg/login"
	"github.com/pkg/browser"
	"github.com/urfave/cli/v2"
)

func loginCommand() *cli.Command {
	return &cli.Command{
		Name:  "login",
		Usage: "Login to ZDV's OAuth2 Provider",
		Action: func(c *cli.Context) error {
			s := login.NewServer()

			go func() {
				s.Engine.Run(fmt.Sprintf(":%d", login.PORT))
			}()

			url := fmt.Sprintf("http://localhost:%d/authorize", login.PORT)

			browser.OpenURL(url)

			cred := <-s.Chan

			fmt.Println("Got response from OAuth2 Server... building credentials.")

			if cred.IdToken == "" {
				return fmt.Errorf("failed to login")
			}
			kc, err := config.GetKubeConfig(cred.IdToken, login.CLIENT_SECRET, cred.RefreshToken)
			if err != nil {
				return fmt.Errorf("error creating kubeconfig: %s", err)
			}
			err = config.WriteKubeConfig(kc)
			if err != nil {
				return fmt.Errorf("error writing kubeconfig: %s", err)
			}

			fmt.Println("Successfully logged in.")
			fmt.Println("You may close your browser now and use `zdv kubectl` to access the cluster with your credentials.")

			return nil
		},
	}
}
