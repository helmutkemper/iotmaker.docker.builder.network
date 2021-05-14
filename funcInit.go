package iotmaker_docker_builder_network

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

func (e *ContainerBuilderNetwork) Init() (err error) {
	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		return
	}

	return
}
