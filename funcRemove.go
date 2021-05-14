package iotmaker_docker_builder_network

func (e *ContainerBuilderNetwork) Remove() (err error) {
	err = e.dockerSys.NetworkRemove(e.networkID)
	return
}
