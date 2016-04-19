package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/types"
	//"github.com/asyoume/paas_srv/pkg/re_act"
)

var appTempHandler *handler.AppTempHandler = new(handler.AppTempHandler)

func appTempGet(c echo.Context) error {

	args := new(types.GetParams)
	app := new(types.App)

	args.Region = c.QueryParam("region")
	if args.Region == "" {
		args.Region = types.DefaultRegion
	}

	args.Id = c.QueryParam("id")
	err := appTempHandler.Get(args, app)
	// 获取数据失败
	if err != nil {
		return c.String(http.StatusNotFound, "")
	}
	return SendProto(c, http.StatusOK, app)
}
