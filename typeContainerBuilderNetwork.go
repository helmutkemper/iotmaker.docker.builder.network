package iotmaker_docker_builder_network

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

type ContainerBuilderNetwork struct {
	dockerSys   iotmakerdocker.DockerSystem
	generator   *iotmakerdocker.NextNetworkAutoConfiguration
	networkID   string
	networkName string
}

// NetworkCreate (Português):
//   subnet "18.18.0.0/16"
//   gateway "18.18.0.1"
func (e *ContainerBuilderNetwork) NetworkCreate(name, subnet, gateway string) (err error) {
	e.networkName = name

	var networkList []types.NetworkResource
	networkList, err = e.dockerSys.NetworkList()
	if err != nil {
		return
	}

	for _, networkData := range networkList {
		if networkData.Name == name {
			e.networkID = networkData.ID
			e.generator = e.dockerSys.NetworkGetGenerator(name)
			return
		}
	}

	e.networkID, e.generator, err = e.dockerSys.NetworkCreate(name, iotmakerdocker.KNetworkDriveBridge, "local", subnet, gateway)
	return
}

func (e *ContainerBuilderNetwork) Init() (err error) {
	e.dockerSys = iotmakerdocker.DockerSystem{}
	err = e.dockerSys.Init()
	if err != nil {
		return
	}

	return
}

func (e *ContainerBuilderNetwork) GetConfiguration() (networkConfiguration *network.NetworkingConfig, err error) {
	networkConfiguration, err = e.generator.GetNext()

	if err != nil {
		err = errors.New("GetNext().error: " + err.Error())
	}

	return
}

func (e *ContainerBuilderNetwork) Remove() (err error) {
	err = e.dockerSys.NetworkRemove(e.networkID)
	return
}
