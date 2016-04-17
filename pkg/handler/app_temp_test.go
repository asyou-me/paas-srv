package handler

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/types"
	"testing"
)

func TestAppTempGet(t *testing.T) {
	Init()
	arg := new(types.GetParams)
	arg.Id = "redis"
	arg.Region = ""
	app := new(types.App)
	appTempHandler := AppTempHandler{}
	err := appTempHandler.Get(arg, app)
	if err != nil {
		fmt.Println("AppHandler get err:", err)
	} else {
		fmt.Println(app)
	}
}
