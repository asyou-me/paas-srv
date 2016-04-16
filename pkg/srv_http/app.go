package main

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/types"
	//"github.com/asyoume/paas_srv/pkg/re_act"
	"github.com/labstack/echo"
	"net/http"
)

var app_h *handler.appHandler = new(handler.appHandler)

func appGet(c echo.Context) error {
	var app = new(types.app)
	app.Region = c.QueryParam("region")
	app.ParentId = c.QueryParam("appid")
	app.Id = c.QueryParam("id")

	err := app_h.Get(app, app)
	if err != nil {
		return protoRetrun(c, http.StatusNotFound, app)
	}
	return protoRetrun(c, http.StatusOK, app)
}

func appPost(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}

func appPut(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}

func appPatch(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}

func appDelete(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}

/*protoBuffer, err := utils.RecvFrame(c.Request().Body())
if err != nil {
  return err
}

err = proto.Unmarshal(protoBuffer, app)
if err != nil {
  return err
}*/
