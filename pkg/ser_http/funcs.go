package main

import (
	"fmt"
	"github.com/asyoume/PaasControl/pkg/handler"
	"github.com/labstack/echo"
	"net/http"
)

var podhander = handler.PodHandler{}

func PodPut(c *echo.Context) error {
	r, err := podhander.Put()
	if err == nil {
		return c.String(http.StatusOK, r)
	} else {
		return c.String(http.StatusInternalServerError, err.Error())
	}
}
