package util

import (
	"fmt"
	"os"
	"runtime"

	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/spf13/cobra"
)

// Docker Client initialization
var DockerClient *docker.Client

func DockerConnect(cmd *cobra.Command) {
	var err error
	verbose := cmd.Flags().Lookup("verbose").Changed

	if runtime.GOOS == "linux" {
		endpoint := "unix:///var/run/docker.sock"

		if verbose {
			fmt.Println("Connecting to the Docker Client via:", endpoint)
		}

		DockerClient, err = docker.NewClient(endpoint)
		if err != nil {
			// TODO: better error handling
			fmt.Println(err)
			os.Exit(1)
		}

		if verbose {
			fmt.Println("Successfully connected to Docker daemon")
		}

	} else {

		path := os.Getenv("DOCKER_CERT_PATH")

		if verbose {
			fmt.Println("Connecting to the Docker Client via:", os.Getenv("DOCKER_HOST"))
		}

		DockerClient, err = docker.NewTLSClient(os.Getenv("DOCKER_HOST"), fmt.Sprintf("%s/cert.pem", path), fmt.Sprintf("%s/key.pem", path), fmt.Sprintf("%s/ca.pem", path))

		if err != nil {
			// TODO: better error handling
			fmt.Println(err)
			os.Exit(1)
		}

		if verbose {
			fmt.Println("Successfully connected to Docker daemon")
		}
	}
}