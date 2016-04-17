package handler

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/types"
	"testing"
)

func TestAppGet(t *testing.T) {
	Init()
	arg := new(types.GetParams)
	arg.Id = "redis"
	arg.Region = ""
	app := new(types.App)
	appHandler := AppHandler{}
	err := appHandler.Get(arg, app)
	fmt.Println("AppHandler get err:", err)
}

func TestAppList(t *testing.T) {
	Init()
	arg := new(types.ListParams)
	arg.Id = "test"
	arg.Region = ""
	arg.Length = 10
	arg.Offset = 0
	list := new(types.AppList)
	appHandler := AppHandler{}
	err := appHandler.List(arg, list)
	fmt.Println("AppHandler list err:", err)
}
