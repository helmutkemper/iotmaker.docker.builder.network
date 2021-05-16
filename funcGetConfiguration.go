package iotmaker_docker_builder_network

import (
	"errors"
	"github.com/docker/docker/api/types/network"
)

func (e *ContainerBuilderNetwork) GetConfiguration() (IP string, networkConfiguration *network.NetworkingConfig, err error) {
	IP, networkConfiguration, err = e.generator.GetNext()

	if err != nil {
		err = errors.New("GetNext().error: " + err.Error())
	}

	return
}
