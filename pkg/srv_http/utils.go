package main

import (
	"github.com/asyoume/paas_srv/pkg/utils"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

func protoRetrun(c echo.Context, code int, data proto.Message) error {
	req := c.Response()
	utils.SendFrame(req, data)
	req.WriteHeader(code)
	return nil
}
