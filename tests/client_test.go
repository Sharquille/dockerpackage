package tests

import (
	"dockerpackage/pkg"
	"testing"
)

func TestNewClient(t *testing.T) {
	// create docker instance to call function
	docker := &pkg.DockerClient{}
	newDockCli, err := docker.NewClient()
	if err != nil {
		t.Logf("error while creating new client: %v", err)
	}
	defer newDockCli.Close()
}
