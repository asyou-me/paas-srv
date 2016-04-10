package utils

import (
	"github.com/asyoume/paas_srv/pkg/types"
	"k8s.io/kubernetes/pkg/api"
	api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/util/intstr"
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
		port[k] = &types.ServicePort{
			Name:       v.Name,
			Protocol:   string(v.Protocol),
			Port:       int32(v.Port),
			TargetPort: v.TargetPort.IntVal,
			NodePort:   int32(v.NodePort),
		}
	}
	todata.Port = port
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
	var ports []api.ServicePort = make([]api.ServicePort, len(args.Port))
	for k1, v := range args.Port {
		// 解析转换端口信息
		ports[k1] = api.ServicePort{
			Name:     v.Name,
			Protocol: api.Protocol(v.Protocol),
			Port:     int(v.Port),
			TargetPort: intstr.IntOrString{
				IntVal: v.TargetPort,
			}, NodePort: int(v.NodePort),
		}
	}
	todata.Spec.Ports = ports
	return todata
}
