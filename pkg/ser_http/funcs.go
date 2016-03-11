package main

import (
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/labstack/echo"
	"net/http"
)

var podhander = handler.PodHandler{}

func PodPut(c *echo.Context) error {
	podhander.Put()
	var err error = nil
	if err == nil {
		return c.String(http.StatusOK, "")
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}
