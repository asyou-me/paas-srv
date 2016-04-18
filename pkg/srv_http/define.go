package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

type (
	RecvFunc func(echo.Context, proto.Message) error
	SendFunc func(echo.Context, int, proto.Message) error
)

var (
	protoTy  = "json"
	RecvData RecvFunc
	SendData SendFunc
)
