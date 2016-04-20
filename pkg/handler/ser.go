package handler

import (
	//"errors"

	"k8s.io/kubernetes/pkg/api"
	//api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	//"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"

	"github.com/asyoume/paas_srv/pkg/types"
	"github.com/asyoume/paas_srv/pkg/utils"
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
		content[k] = utils.ServiceToPbStruct(&v)
	}
	reply.Content = content
	return err
}

func (this *SerHandler) Get(args *types.Service, reply *types.Service) error {
	c := NewkubeClient()
	sers := c.Services(args.ParentId)
	ser, err := sers.Get(args.Name)
	if err != nil {
		return err
	}
	*reply = *utils.ServiceToPbStruct(ser)
	return nil
}

func (this *SerHandler) Post(args *types.Service, reply *types.Event) error {
	c := NewkubeClient()
	sers := c.Services(args.ParentId)

	// 转换配置文件
	conf := utils.ServiceTokubenetStruct(args)

	_, err := sers.Create(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *SerHandler) Put(args *types.Service, reply *types.Event) error {
	c := NewkubeClient()
	sers := c.Services(args.ParentId)

	// 转换配置文件
	conf := utils.ServiceTokubenetStruct(args)

	reply.Id = conf.GetName()

	_, err := sers.Update(conf)
	if err != nil {
		return err
	}
	return nil
}

func (this *SerHandler) Delete(args *types.DeleteParams, reply *types.Event) error {
	c := NewkubeClient()
	sers := c.Services(args.ParentId)
	err := sers.Delete(args.Id)
	if err != nil {
		return err
	}
	return nil
}
