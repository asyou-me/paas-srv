package main

import (
	"fmt"
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/types"
	//"github.com/asyoume/paas_srv/pkg/re_act"
	"github.com/labstack/echo"
	"net/http"
)

var pod_h *handler.PodHandler = new(handler.PodHandler)

func podGet(c echo.Context) error {

	var pod = new(types.Pod)
	pod.Region = c.QueryParam("region")
	pod.ParentId = c.QueryParam("appid")
	pod.Id = c.QueryParam("id")

	err := pod_h.Get(pod, pod)
	if err != nil {
		return protoRetrun(c, http.StatusNotFound, pod)
	}
	return protoRetrun(c, http.StatusOK, pod)
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
