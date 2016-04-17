package main

import (
	"github.com/asyoume/paas_srv/pkg/types"
	"github.com/labstack/echo"
	"net/http"
)

// 获取 api 版本
func version(c echo.Context) error {
	return c.String(http.StatusOK, types.ApiVersion)
}

// 获取 api 版本信息
func versionInfo(c echo.Context) error {
	return c.String(http.StatusOK, types.ApiVersionInfo)
}
