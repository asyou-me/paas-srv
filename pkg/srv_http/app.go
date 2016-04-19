package main

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"

	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/types"
)

var appHandler *handler.AppHandler = new(handler.AppHandler)

func appGet(c echo.Context) error {

	ty := c.QueryParam("type")
	var returnValue proto.Message
	var err error
	var Region = c.QueryParam("region")

	if args.Region == "" {
		args.Region = types.DefaultRegion
	}

	if ty == "value" { // 获取单个应用信息
		args := new(types.GetParams)
		app := new(types.App)
		returnValue = app
		args.Region = Region
		args.Id = c.QueryParam("id")
		err = appHandler.Get(args, app)
	} else { // 获取应用列表
		args := new(types.ListParams)
		applist := new(types.AppList)
		returnValue = applist
		args.Region = Region
		args.Offset = 0  // 列表偏移量
		args.Length = 10 // 列表长度
		err = appHandler.List(args, applist)
	}

	// 获取数据失败
	if err != nil {
		return SendProto(c, http.StatusNotFound, returnValue)
	}

	return SendProto(c, http.StatusOK, returnValue)
}

func appPost(c echo.Context) error {
	app := new(types.App)
	reply := new(types.Event)
	err := RecvProto(c, app)

	if err != nil {
		reply.Code = http.StatusBadRequest
		reply.Content = "数据格式为非protobuf"
		return SendProto(c, http.StatusBadRequest, reply)
	}

	err = appHandler.Post(app, reply)

	if err != nil {
		return SendProto(c, http.StatusInternalServerError, reply)
	}

	return SendProto(c, http.StatusOK, reply)
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
