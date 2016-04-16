package handler

import (
	"fmt"
	//"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func NewkubeClient() *client.Client {
	kubeclient, err := client.New(kubeConfig)
	if err != nil {
		fmt.Println(err)
	}
	return kubeclient
}
