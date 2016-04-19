package handler

import (
	"k8s.io/kubernetes/pkg/client/restclient"

	"github.com/asyoume/lib/pulic_type"
)

var kubeConfig *restclient.Config

func Init(conf ...*pulic_type.MicroSerType) error {
	kubeConfig = &restclient.Config{
		Host:     "http://115.29.113.249:8080",
		Username: "test",
		Password: "password",
	}

	return nil
}
