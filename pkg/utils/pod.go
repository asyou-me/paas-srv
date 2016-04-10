package utils

import (
	"github.com/asyoume/paas_srv/pkg/types"
	"k8s.io/kubernetes/pkg/api"
	api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
)

// kubenetes pod to protobuf struct
func PodToPbStruct(pod *api.Pod) *types.Pod {
	todata := &types.Pod{
		Id:   string(pod.GetUID()),
		Name: pod.GetName(),
	}
	pod_container := pod.Spec.Containers
	// 解析容器信息
	container := make([]*types.Container, len(pod_container))
	for k, v := range pod_container {
		container[k] = &types.Container{
			Name:  v.Name,
			Image: v.Image,
		}
		// 解析空间挂载信息
		VolumeMounts := make([]*types.VolumeMount, len(v.VolumeMounts))
		for k2, v2 := range v.VolumeMounts {
			VolumeMounts[k2] = &types.VolumeMount{
				Name:      v2.Name,
				ReadOnly:  v2.ReadOnly,
				MountPath: v2.MountPath,
			}
		}
		// 解析端口映射信息
		containerPort := make([]*types.ContainerPort, len(v.Ports))
		for k2, port := range v.Ports {
			containerPort[k2] = &types.ContainerPort{
				Name:          port.Name,
				Protocol:      string(port.Protocol),
				ContainerPort: int64(port.ContainerPort),
			}
		}
		container[k].Port = containerPort
	}
	todata.Containers = container
	return todata
}

// protobuf struct  to  kubenetes pod
func PodToKubeStruct(args *types.Pod) *api.Pod {
	// 建立配置文件
	todata := &api.Pod{
		TypeMeta: api_unversioned.TypeMeta{
			Kind:       "Pod",
			APIVersion: types.ApiVersion,
		},
		ObjectMeta: api.ObjectMeta{
			Name:   args.Name,
			Labels: args.Labels,
		},
		Spec:   api.PodSpec{},
		Status: api.PodStatus{},
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
