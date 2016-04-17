package main

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/types"
	//"github.com/asyoume/paas_srv/pkg/re_act"
	"github.com/labstack/echo"
	"net/http"
)

var podHandler *handler.PodHandler = new(handler.PodHandler)

func podGet(c echo.Context) error {

	var arg = new(types.GetParams)
	arg.Region = c.QueryParam("region")
	arg.ParentId = c.QueryParam("appid")
	arg.Id = c.QueryParam("id")
	var pod = new(types.Pod)

	err := podHandler.Get(arg, pod)
	if err != nil {
		return SendProto(c, http.StatusNotFound, pod)
	}
	return SendProto(c, http.StatusOK, pod)
}

/*protoBuffer, err := utils.RecvFrame(c.Request().Body())
if err != nil {
	return err
}

err = proto.Unmarshal(protoBuffer, pod)
if err != nil {
	return err
}*/

func podPost(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}

func podPut(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}

func podPatch(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}

func podDelete(c echo.Context) error {
	data := []byte{}
	fmt.Println(data)
	return c.JSONBlob(http.StatusOK, data)
}
