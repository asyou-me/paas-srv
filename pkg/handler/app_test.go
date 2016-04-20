package handler

import (
	"fmt"
	"testing"

	"github.com/asyoume/paas_srv/pkg/types"
)

func TestAppGet(t *testing.T) {
	Init()
	arg := new(types.GetParams)
	arg.Id = "redis"
	arg.Region = ""
	app := new(types.App)

	appHandler := AppHandler{}

	err := appHandler.Get(arg, app)

	if err != nil {
		t.Error("AppHandler get err:", err)
	}
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

	if err != nil {
		t.Error("AppHandler list err:", err)
	}
}
