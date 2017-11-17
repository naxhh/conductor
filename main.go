package main

import (
	docker "github.com/docker/docker/client"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	clientset := getClientset()
	dockerClient := getDockerClient()
	builder := NewBuilder(dockerClient)
	deployer := NewDeployer(clientset)
	server := NewServer(clientset, builder, deployer)
	server.Start()
}

func getDockerClient() *docker.Client {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return cli
}

func getClientset() *kubernetes.Clientset {
	config := inClusterConfig()

	return kubernetes.NewForConfigOrDie(config)
}

func inClusterConfig() *rest.Config {
	config, err := rest.InClusterConfig()

	if err != nil {
		panic(err.Error())
	}

	return config
}
