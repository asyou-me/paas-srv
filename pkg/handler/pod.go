package handler

import (
	"fmt"
	"k8s.io/kubernetes/pkg/api"
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
	client := NewkubeClient()
	pods := client.Pods("test")
	_, err := pods.Create(&api.Pod{
		TypeMeta: api.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1beta3",
		},
		ObjectMeta: api.ObjectMeta{
			Name:   "podtest",
			Labels: map[string]string{"name": "foo2"},
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				api.Container{
					Name:  "containertest",
					Image: "nginx",
					Ports: []api.ContainerPort{
						api.ContainerPort{
							HostPort:      8081,
							ContainerPort: 80,
						},
					},
				},
			},
		},
		Status: api.PodStatus{},
	})
}
