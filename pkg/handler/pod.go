package handler

import (
	"errors"

	"k8s.io/kubernetes/pkg/api"
	//api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	//"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"

	"github.com/asyoume/paas_srv/pkg/types"
	"github.com/asyoume/paas_srv/pkg/utils"
)

type PodHandler struct {
}

func (this *PodHandler) Get(args *types.GetParams, reply *types.Pod) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)
	pod, err := pods.Get(args.Id)
	if err != nil {
		return err
	}
	*reply = *utils.PodToPbStruct(pod)
	return nil
}

func (this *PodHandler) List(args *types.ListParams, reply *types.PodList) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	selector := labels.Set(args.Labels).AsSelector()
	options := api.ListOptions{LabelSelector: selector}

	podList, err := pods.List(options)
	if err != nil {
		//reply.Code = 500
		reply.Region = args.Region
		return err
	}
	content := make([]*types.Pod, len(podList.Items))
	for k, v := range podList.Items {
		content[k] = utils.PodToPbStruct(&v)
	}
	reply.Content = content
	return err
}

func (this *PodHandler) Post(args *types.Pod, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	if len(args.Containers) == 0 {
		return errors.New("一个实例至少有一个容器")
	}

	// 转换配置文件
	conf := utils.PodToKubeStruct(args)

	_, err := pods.Create(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *PodHandler) Put(args *types.Pod, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	// 转换配置文件
	conf := utils.PodToKubeStruct(args)

	reply.Id = conf.GetName()

	_, err := pods.Update(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *PodHandler) Patch(args *types.Pod, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	// 转换配置文件
	conf := utils.PodToKubeStruct(args)

	reply.Id = conf.GetName()

	_, err := pods.Update(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *PodHandler) Delete(args *types.DeleteParams, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)
	err := pods.Delete(args.Id, nil)
	if err != nil {
		return err
	}
	return nil
}
