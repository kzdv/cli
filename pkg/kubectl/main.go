package kubectl

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/kzdv/cli/pkg/config"
	"github.com/urfave/cli/v2"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	kubecli "k8s.io/component-base/cli"
	"k8s.io/kubectl/pkg/cmd"
	"k8s.io/kubectl/pkg/cmd/util"
)

func Run(ctx *cli.Context) error {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	dir, err := config.GetConfigDir()
	if err != nil {
		panic(err)
	}

	exe := filepath.Join(exPath, "kubectl")
	args := append(os.Args[1:], "--kubeconfig="+filepath.Join(dir, "kubeconfig"))

	return syscall.Exec(exe, args, os.Environ())
}

func Main() {
	var kubenv string
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "--kubeconfig=") {
			kubenv = strings.Split(arg, "=")[1]
		} else if strings.HasPrefix(arg, "--kubeconfig") && i+1 < len(os.Args) {
			kubenv = os.Args[i+1]
		}
	}
	// Force set our kubeconfig
	if kubenv == "" {
		dir, err := config.GetConfigDir()
		if err != nil {
			panic(err)
		}
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		exe := filepath.Join(exPath, "kubectl")
		args := append(os.Args, "--kubeconfig="+filepath.Join(dir, "kubeconfig"))

		syscall.Exec(exe, args, os.Environ())
		return
	}

	rand.Seed(time.Now().UnixNano())
	command := cmd.NewDefaultKubectlCommand()
	if err := kubecli.RunNoErrOutput(command); err != nil {
		fmt.Printf("Error running kubectl: %s\n", err)
		util.CheckErr(err)
	}
}
