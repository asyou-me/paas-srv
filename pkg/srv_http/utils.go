package main

import (
	"github.com/asyoume/paas_srv/pkg/utils"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

func RecvProto(c echo.Context, data proto.Message) error {
	req := c.Request()
	return utils.RecvFrame(req.Body(), data)
}

// 将结果写入到http请求
func SendProto(c echo.Context, code int, data proto.Message) error {
	resp := c.Response()
	err := utils.SendFrame(resp, data)
	if err != nil {
		return err
	}
	resp.WriteHeader(code)
	return nil
}
