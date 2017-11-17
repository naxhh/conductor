package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	clientset := getClientset()
	server := NewServer(clientset)
	server.Start()
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
