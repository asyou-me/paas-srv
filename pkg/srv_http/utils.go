package main

import (
	"github.com/asyoume/paas_srv/pkg/utils"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

// 将结果写入到http请求
func protoRetrun(c echo.Context, code int, data proto.Message) error {
	req := c.Response()
	utils.SendFrame(req, data)
	req.WriteHeader(code)
	return nil
}
