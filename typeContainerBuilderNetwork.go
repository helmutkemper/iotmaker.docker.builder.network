package iotmaker_docker_builder_network

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

type ContainerBuilderNetwork struct {
	dockerSys   iotmakerdocker.DockerSystem
	generator   *iotmakerdocker.NextNetworkAutoConfiguration
	networkID   string
	networkName string
}
