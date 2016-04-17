package handler

import (
	"errors"
	"fmt"
	"github.com/asyoume/paas_srv/pkg/types"
	"github.com/asyoume/paas_srv/pkg/utils"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/labels"
)

type AppHandler struct {
}

// 获取单台服务器信息
func (this *AppHandler) Get(args *types.GetParams, reply *types.App) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)
	pod, err := pods.Get(args.Id)
	if err != nil {
		return err
	}
	fmt.Println(pod)
	//*reply = *utils.PodToPbStruct(pod)
	return nil
}

// 获取服务器列表
func (this *AppHandler) List(args *types.ListParams, reply *types.AppList) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)

	selector := labels.Set(args.Labels).AsSelector()
	options := api.ListOptions{LabelSelector: selector}

	fmt.Println(pods)
	fmt.Println(options)
	fmt.Println(selector)

	/*
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
	*/
	return nil
}

// 创建服务器
func (this *AppHandler) Post(args *types.App, reply *types.Event) error {
	c := NewkubeClient()
	args.Id = "default"

	if len(args.Pods) == 0 {
		reply.Code = 500
		reply.Content = "Pods lenght 0"
		return errors.New("Pods lenght 0")
	}

	pods := c.Pods(args.Id)
	for _, v := range args.Pods {
		// 转换实例配置文件
		conf := utils.PodToKubeStruct(v)
		_, err := pods.Create(conf)
		if err != nil {
			return err
		}
	}

	sers := c.Services(args.Id)
	for _, v := range args.Services {
		// 转换网络配置文件
		conf := utils.ServiceTokubenetStruct(v)
		_, err := sers.Create(conf)
		if err != nil {
			return err
		}
	}

	reply.Code = 200
	reply.Content = "ok"

	return nil
}

/*
//
func (this *AppHandler) Put(args *types.Pod, reply *types.Event) error {
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

//
func (this *AppHandler) Patch(args *types.Pod, reply *types.Event) error {
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


*/

//
func (this *AppHandler) Delete(args *types.DeleteParams, reply *types.Event) error {
	c := NewkubeClient()
	pods := c.Pods(args.ParentId)
	err := pods.Delete(args.Id, nil)
	if err != nil {
		return err
	}
	return nil
}

func AppAttributeCode(app *types.App) {

}
