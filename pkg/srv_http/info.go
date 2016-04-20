package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/asyoume/paas_srv/pkg/types"
)

// 获取 api 版本
func version(c echo.Context) error {
	return c.String(http.StatusOK, types.ApiVersion)
}

// 获取 api 版本信息
func info(c echo.Context) error {
	return c.String(http.StatusOK, types.Info)
}

// 获取 api 版本信息
func apiMap(c echo.Context) error {
	return c.String(http.StatusOK, types.Info)
}
