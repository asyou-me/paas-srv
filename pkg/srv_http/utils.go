package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"

	"github.com/asyoume/paas_srv/pkg/utils"
)

func selectProto(ty string) {
	protoTy = ty
	if ty == "proto" {
		RecvData = RecvProto
		SendData = SendProto
	} else {
		RecvData = RecvJson
		SendData = SendJson
	}
}

// 从http中获取数据(ProtoBuffer 格式)
func RecvProto(c echo.Context, data proto.Message) error {
	req := c.Request()
	return utils.RecvFrame(req.Body(), data)
}

// 将结果写入到http请求(ProtoBuffer 格式)
func SendProto(c echo.Context, code int, data proto.Message) error {
	resp := c.Response()
	err := utils.SendFrame(resp, data)
	if err != nil {
		return err
	}
	resp.WriteHeader(code)
	return nil
}

// 从http中获取数据(Json 格式)
func RecvJson(c echo.Context, data proto.Message) error {
	req := c.Request()
	body := req.Body()
	buffer, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffer, data)
	return err
}

// 将结果写入到http请求(Json 格式)
func SendJson(c echo.Context, code int, data proto.Message) error {
	resp := c.Response()
	buffer, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp.Write(buffer)
	resp.WriteHeader(code)
	resp.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return nil
}
