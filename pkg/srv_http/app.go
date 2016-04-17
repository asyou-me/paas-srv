package main

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/types"
	//"github.com/asyoume/paas_srv/pkg/re_act"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
	"net/http"
)

var appHandler *handler.AppHandler = new(handler.AppHandler)

func appGet(c echo.Context) error {

	ty := c.QueryParam("type")
	var returnValue proto.Message
	var err error

	if ty == "value" { // 获取单个应用信息
		args := new(types.GetParams)
		app := new(types.App)
		returnValue = app
		args.Region = c.QueryParam("region")
		args.Id = c.QueryParam("id")
		err = appHandler.Get(args, app)
	} else if ty == "list" { // 获取应用列表
		args := new(types.ListParams)
		applist := new(types.AppList)
		returnValue = applist
		args.Region = c.QueryParam("region")
		args.Id = c.QueryParam("id")
		args.Offset = 0  // 列表偏移量
		args.Length = 10 // 列表长度
		err = appHandler.List(args, applist)
	} else { // 没有类型返回错误
		return c.String(http.StatusNotFound, "")
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
		fmt.Println(err)
		fmt.Println(reply)
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
