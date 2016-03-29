package handler

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/types"
	//"github.com/asyoume/paas_srv/pkg/utils"
	"k8s.io/kubernetes/pkg/api"
	api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	//"k8s.io/kubernetes/pkg/fields"
	"errors"
	"k8s.io/kubernetes/pkg/labels"
)

type PodHandler struct {
}

func (this *PodHandler) List(args *types.ListParams, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	selector := labels.Set(args.Labels).AsSelector()
	options := api.ListOptions{LabelSelector: selector}

	podList, err := pods.List(options)
	fmt.Println(podList)
	return err
}

func (this *PodHandler) Get(args *types.DeleteParams, reply *types.Instance) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)
	_, err := pods.Get(args.Id)
	if err != nil {
		return err
	}
	return nil
}

func (this *PodHandler) Put(args *types.Instance, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	if len(args.Container) == 0 {
		return errors.New("一个实例至少有一个容器")
	}

	// 建立配置文件
	conf := &api.Pod{
		TypeMeta: api_unversioned.TypeMeta{
			Kind:       "Pod",
			APIVersion: types.ApiVersion,
		},
		ObjectMeta: api.ObjectMeta{
			Name:   args.Id,
			Labels: args.Labels,
		},
		Spec:   api.PodSpec{},
		Status: api.PodStatus{},
	}

	// 解析转换容器信息
	var container []api.Container = make([]api.Container, len(args.Container))
	for k1, v := range args.Container {
		// 解析转换端口信息
		var containerPort = make([]api.ContainerPort, len(v.Port))
		for k2, port := range v.Port {
			containerPort[k2] = api.ContainerPort{
				Protocol:      api.Protocol(port.Protocol),
				ContainerPort: int(port.ContainerPort),
			}
		}
		// 写入镜像信息
		container[k1] = api.Container{
			Name:  v.Name,
			Image: v.Image,
			Ports: containerPort,
		}
	}
	conf.Spec.Containers = container
	reply.Id = conf.GetName()

	_, err := pods.Create(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *PodHandler) Update(args *types.Instance, reply *types.Instance) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	// 转换配置文件
	conf := &api.Pod{
		TypeMeta: api_unversioned.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: api.ObjectMeta{
			Name:   args.Id,
			Labels: args.Labels,
		},
		Spec: api.PodSpec{
			Containers: []api.Container{
				api.Container{
					Name:  "postgresql",
					Image: "hub.asyou.me/postgresql",
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
	}

	reply.Id = conf.GetName()

	_, err := pods.Create(conf)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (this *PodHandler) Del(args *types.DeleteParams, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)
	err := pods.Delete(args.Id, nil)
	if err != nil {
		return err
	}
	return nil
}
