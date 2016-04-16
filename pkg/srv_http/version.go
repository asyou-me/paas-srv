package main

import (
	"github.com/asyoume/paas_srv/pkg/utils"
	"github.com/labstack/echo"
)

// 获取 api 版本
func version(c echo.Context) error {
	return c.Sring(http.StatusOK, utils.ApiVersion)
}

// 获取 api 版本信息
func versionInfo(c echo.Context) error {
	return c.Sring(http.StatusOK, utils.ApiVersionInfo)
}
