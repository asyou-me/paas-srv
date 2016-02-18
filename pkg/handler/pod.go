package handler

import (
	"fmt"
	"k8s.io/kubernetes/pkg/api"
	api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	//client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
)

type PodHandler struct {
}

func (this *PodHandler) List(label labels.Selector, field fields.Selector, device_id string) *api.PodList {
	c := NewkubeClient()
	pods := c.Pods(device_id)
	podList, _ := pods.List(api.ListOptions{
		LabelSelector: label,
		FieldSelector: field,
	})
	return podList
}

func (this *PodHandler) Get(name string, device_id string) {
	c := NewkubeClient()
	pods := c.Pods(device_id)
	_, err := pods.Get(name)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *PodHandler) Put() {
	c := NewkubeClient()
	pods := c.Pods("development")
	_, err := pods.Create(&api.Pod{
		TypeMeta: api_unversioned.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: api.ObjectMeta{
			Name:   "postgresql",
			Labels: map[string]string{"name": "gitlab"},
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				api.Container{
					Name:  "postgresql",
					Image: "hub.asyou.me:5000/postgresql",
					Ports: []api.ContainerPort{
						api.ContainerPort{
							Protocol:      "TCP",
							ContainerPort: 5432,
						},
					},
				},
			},
		},
		Status: api.PodStatus{},
	})
	if err != nil {
		fmt.Println(err)
	}
}
