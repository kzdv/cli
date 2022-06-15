package main

import (
	"fmt"
	"os"

	"github.com/kzdv/cli/cmd/zdv/app"
	"github.com/kzdv/cli/pkg/config"
)

func main() {
	err := config.MakeConfigDir()
	if err != nil {
		fmt.Printf("Error creating config path: %v\n", err)
		panic(err)
	}

	app := app.NewRootCommand()
	err = app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
