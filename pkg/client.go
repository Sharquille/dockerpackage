/*
Package pkg is used to manage docker containers in a grocery store application.

This package provides functions for creating, starting, stopping, and removing docker containers that are used in the grocery store application.

Functions:
- CreateContainer: Creates a new docker container with the specified image and options.
- StartContainer: Starts a docker container with the specified container ID.
- StopContainer: Stops a running docker container with the specified container ID.
- RemoveContainer: Removes a docker container with the specified container ID.
*/

package pkg

import (
	"github.com/docker/docker/client"
)

// docker represents a Docker client with methods for managing
// Docker operations
type DockerClient struct {
	cli *client.Client
}

// NewClient creates a new Docker client and returns it along with any error encountered.
// It initializes a new Docker client using the client.NewClientWithOpts function with client.FromEnv option.
//
// Parameters:
// - d: a pointer to a docker struct, which holds the Docker client instance
//
// Returns:
// - *client.Client: a pointer to the newly created Docker client
// - error: an error, if any, encountered during the creation of the Docker client
//
// Example usage:
//
//	docker := &docker{}
//	dockerClient, err := docker.NewClient()
//	if err != nil {
//	    log.Fatalf("Error creating Docker client: %v", err)
//	}
//	defer dockerClient.Close()
func (d *DockerClient) NewClient() (*client.Client, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	d.cli = dockerClient
	return dockerClient, nil
}
