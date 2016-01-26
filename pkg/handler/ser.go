package handler

import (
	"fmt"
	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"
)

type SerHandler struct {
}

func (this *SerHandler) List(c *client.Client, namespace string, label labels.Selector) *api.ServiceList {
	Sers := c.Services(namespace)
	SerList, _ := Sers.List(api.ListOptions{
		LabelSelector: label,
	})
	return SerList
}

func (this *SerHandler) Get(c *client.Client, namespace string, name string) {
	Sers := c.Pods(namespace)
	_, err := Sers.Get(name)
	if err != nil {
		fmt.Println(err)
	}
}
