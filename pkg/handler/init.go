package handler

import (
	"github.com/asyoume/lib/pulic_type"
	"k8s.io/kubernetes/pkg/client/restclient"
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
