package handler

import (
	"github.com/asyoume/paas_srv/pkg/types"
)

type AppTempHandler struct {
}

// 获取应用模板
func (this *AppTempHandler) Get(args *types.GetParams, reply *types.App) error {
	*reply = types.App{
		Region: "shanghai",
		Pods: [](*types.Pod){
			&types.Pod{
				Name: "cs",
				Labels: map[string]string{
					"pod": "cs",
				},
				Containers: []*types.Container{
					&types.Container{
						Name:    "cs",
						Image:   "hub.gmcloud.io/cs_server",
						Version: "1.7",
						Port: []*types.ContainerPort{
							&types.ContainerPort{
								Name:          "cs",
								ContainerPort: 27015,
								Protocol:      "udp",
							},
						},
					},
				},
			},
		},
		Services: []*types.Service{
			&types.Service{
				Name: "cs-server",
				Port: []*types.ServicePort{
					&types.ServicePort{
						Name:       "cs",
						Protocol:   "udp",
						Port:       27015,
						TargetPort: 27015,
						NodePort:   30001,
					},
				},
			},
		},
		Conf:     map[string]string{},
		Editable: []string{""},
	}

	return nil
}
