package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/types"
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
		return SendData(c, http.StatusNotFound, pod)
	}
	return SendData(c, http.StatusOK, pod)
}

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
