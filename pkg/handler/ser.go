package handler

import (
	"github.com/asyoume/paas_srv/pkg/types"
	//"github.com/asyoume/paas_srv/pkg/utils"
	"k8s.io/kubernetes/pkg/api"
	//api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	//"k8s.io/kubernetes/pkg/fields"
	//"errors"
	"k8s.io/kubernetes/pkg/labels"
)

type SerHandler struct {
}

func (this *SerHandler) List(args *types.ListParams, reply *types.ServiceList) error {
	c := NewkubeClient()
	sers := c.Services(args.ParentId)

	selector := labels.Set(args.Labels).AsSelector()
	options := api.ListOptions{LabelSelector: selector}

	serList, err := sers.List(options)
	if err != nil {
		reply.Code = 500
		reply.Region = args.Region
		return err
	}
	content := make([]*types.Service, len(serList.Items))
	for k, v := range serList.Items {
		//content[k] = utils.PodToPbStruct(&v)
	}
	reply.Content = content
	return err
}

/*

func (this *SerHandler) Get(args *types.Pod, reply *types.Pod) error {
	c := NewkubeClient()
	pods := c.Services(args.ParentId)
	pod, err := pods.Get(args.Name)
	if err != nil {
		return err
	}
	*reply = *utils.PodToPbStruct(pod)
	return nil
}

func (this *SerHandler) Post(args *types.Pod, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Services(args.ParentId)

	if len(args.Containers) == 0 {
		return errors.New("一个实例至少有一个容器")
	}

	// 转换配置文件
	conf := utils.PodToPodStruct(args)

	_, err := pods.Create(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *SerHandler) Put(args *types.Pod, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Services(args.ParentId)

	// 转换配置文件
	conf := utils.PodToPodStruct(args)

	reply.Id = conf.GetName()

	_, err := pods.Update(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *SerHandler) Delete(args *types.DeleteParams, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Services(args.ParentId)
	err := pods.Delete(args.Id, nil)
	if err != nil {
		return err
	}
	return nil
}
*/
