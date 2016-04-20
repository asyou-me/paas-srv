package main

import (
	"fmt"
	"time"

	"github.com/asyoume/paas_srv/pkg/types"
	"github.com/asyoume/protorpc"
)

func main() {
	cli, err := protorpc.Dial("tcp", "127.0.0.1:1235")

	if err != nil {
		fmt.Println(err)
	}

	defer cli.Close()

	args := &types.Pod{}
	args.Name = "woyun2"
	args.ParentId = "default"
	args.Containers = []*types.Container{
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

	err = cli.Call("Pod.Post", args, reply)
	if err != nil {
		fmt.Println("Pod.Post: expected no error but got string %q", err.Error())
	} else {
		fmt.Println("Pod.Post: PASS")
	}

	err = cli.Call("Pod.Get", args, args)
	if err != nil {
		fmt.Println("Pod.Get: expected no error but got string %q", err.Error())
	} else {
		fmt.Println("Pod.Get: PASS")
		fmt.Println(args)
	}

	/*args.ParentId = "default"
	err = cli.Call("Pod.Put", args, reply)
	if err != nil {
		fmt.Println("Pod.Put error : ", err.Error())
	} else {
		fmt.Println("Pod.Put: PASS")
	}*/

	list_args := &types.ListParams{}
	list_args.ParentId = "default"
	list_reply := &types.PodList{}
	err = cli.Call("Pod.List", list_args, list_reply)
	if err != nil {
		fmt.Println("Pod.List: expected no error but got string %q", err.Error())
	} else {
		fmt.Println("Pod.List: PASS")
		fmt.Println(list_reply)
	}

	args2 := &types.DeleteParams{}
	args2.Id = "woyun2"
	args2.ParentId = "default"
	reply2 := new(types.Event)

	err = cli.Call("Pod.Delete", args2, reply2)
	if err != nil {
		fmt.Println("Pod.Delete: error %q", err.Error())
	} else {
		fmt.Println("Pod.Delete: PASS")
	}

	args3 := new(types.GetParams)
	args3.Id = "222"
	reply3 := new(types.App)
	err = cli.Call("AppTemp.Get", args3, reply3)
	if err != nil {
		fmt.Println("AppTemp.Get: error %q", err.Error())
	} else {
		fmt.Println("AppTemp.Get: PASS")
		fmt.Println(reply3)
	}

	time.Sleep(5 * time.Second)
}
