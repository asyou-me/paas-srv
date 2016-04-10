package utils

import (
	"github.com/asyoume/paas_srv/pkg/types"
	"k8s.io/kubernetes/pkg/api"
	api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
)

// kubenetes Service to protobuf struct
func ServiceToPbStruct(Service *api.Service) *types.Service {
	todata := &types.Service{
		Name:     Service.GetName(),
		Labels:   Service.ObjectMeta.Labels,
		Selector: Service.Spec.Selector,
	}
	Service_port := Service.Spec.Ports
	// 解析容器信息
	port := make([]*types.ServicePort, len(Service_port))
	for k, v := range Service_port {
		port[k] = &types.Container{
			Name:       v.Name,
			Protocol:   string(v.Protocol),
			Port:       int32(v.Port),
			TargetPort: int32(v.TargetPort),
			NodePort:   int32(v.NodePort),
		}
	}
	todata.Containers = container
	return todata
}

// protobuf struct  to  kubenetes Service
func ServiceTokubenetStruct(args *types.Service) *api.Service {
	// 建立配置文件
	todata := &api.Service{
		TypeMeta: api_unversioned.TypeMeta{
			Kind:       "Service",
			APIVersion: types.ApiVersion,
		},
		ObjectMeta: api.ObjectMeta{
			Name:   args.Name,
			Labels: args.Labels,
		},
		Spec:   api.ServiceSpec{},
		Status: api.ServiceStatus{},
	}

	// 解析转换容器信息
	var container []api.Container = make([]api.Container, len(args.Containers))
	for k1, v := range args.Containers {
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
	todata.Spec.Containers = container
	return todata
}
