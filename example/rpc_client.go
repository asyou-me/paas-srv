package main

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/types"
	"github.com/asyoume/protorpc"
	"time"
)

func main() {
	cli, err := protorpc.Dial("tcp", "127.0.0.1:1235")

	if err != nil {
		fmt.Println(err)
	}

	defer cli.Close()

	args := &types.Instance{}
	args.Id = "woyun2"
	args.ParentId = "default"
	args.Container = []*types.Container{
		&types.Container{
			Name:  "redis",
			Image: "hub.asyou.me/gitlab/redis",
			Port: []*types.ContainerPort{
				&types.ContainerPort{
					Name:          "redis",
					ContainerPort: 6379,
				},
			},
		},
	}
	reply := new(types.Event)

	err = cli.Call("Instance.Put", args, reply)
	if err != nil {
		fmt.Println("Instance.Put: expected no error but got string %q", err.Error())
	} else {
		fmt.Println("Instance.Put: PASS")
	}

	list_args := &types.ListParams{}
	list_args.ParentId = "default"
	err = cli.Call("Instance.List", list_args, reply)
	if err != nil {
		fmt.Println("Instance.List: expected no error but got string %q", err.Error())
	} else {
		fmt.Println("Instance.List: PASS")
	}

	args2 := &types.DeleteParams{}
	args2.Id = "woyun2"
	args2.ParentId = "default"
	reply2 := new(types.Event)

	err = cli.Call("Instance.Del", args2, reply2)
	if err != nil {
		fmt.Println("Instance.Del: expected no error but got string %q", err.Error())
	} else {
		fmt.Println("Instance.Del: PASS")
	}

	time.Sleep(5 * time.Second)
}
