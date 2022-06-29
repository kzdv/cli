package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/reexec"
	"github.com/kzdv/cli/cmd/zdv/app"
	"github.com/kzdv/cli/pkg/config"
	"github.com/kzdv/cli/pkg/kubectl"
)

func init() {
	reexec.Register("kubectl", kubectl.Main)
}

func main() {
	cmd := os.Args[0]
	os.Args[0] = filepath.Base(os.Args[0])
	if reexec.Init() {
		return
	}
	os.Args[0] = cmd

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
