package config

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func ClientSet() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", EnvConfig("CONFIG_PATH"))
	if err != nil {
		panic(fmt.Errorf("fail to build the k8s config. Error - %s", err))
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Errorf("fail to create the k8s client set. Errorf - %s", err))
	}

	return clientSet
}

func ClientSetCoreV1() v1.CoreV1Interface {
	return ClientSet().CoreV1()
}
