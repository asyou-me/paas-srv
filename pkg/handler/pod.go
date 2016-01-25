package handler

import (
	"fmt"
	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
)

func PodList(c *client.Client, namespace string, label labels.Selector, field fields.Selector) *api.PodList {
	pods := c.Pods(namespace)
	podList, _ := pods.List(api.ListOptions{
		LabelSelector: label,
		FieldSelector: field,
	})
	return podList
}

func PodGet(c *client.Client, namespace string, name string) {
	pods := c.Pods(namespace)
	_, err := pods.Get(name)
	if err != nil {
		fmt.Println(err)
	}
}
