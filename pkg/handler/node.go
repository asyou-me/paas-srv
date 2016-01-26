package handler

import (
	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
)

type NodeHandler struct {
}

func (this *NodeHandler) Put(c *client.Client, label labels.Selector, field fields.Selector) (*api.NodeList, error) {
	nodes := c.Nodes()
	return nodes.List(api.ListOptions{
		LabelSelector: label,
		FieldSelector: field,
	})
}

func (this *NodeHandler) List(c *client.Client, label labels.Selector, field fields.Selector) (*api.NodeList, error) {
	nodes := c.Nodes()
	return nodes.List(api.ListOptions{
		LabelSelector: label,
		FieldSelector: field,
	})
}
