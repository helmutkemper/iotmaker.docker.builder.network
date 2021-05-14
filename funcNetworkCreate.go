package iotmaker_docker_builder_network

import (
	"github.com/docker/docker/api/types"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.1"
)

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
