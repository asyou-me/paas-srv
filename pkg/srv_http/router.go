package main

import (
	"flag"
	"fmt"
	base_utils "github.com/asyoume/lib/utils"
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/re_act"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	mw "github.com/labstack/echo/middleware"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	// 获取执行参数
	conf_path := flag.String("conf", "", "配置文件路径")
	flag.Parse()
	// 获取配置文件的路径
	base_utils.CompletePath(conf_path)

	// 初始化log日志系统
	conf, err := base_utils.ConfigInit(*conf_path)
	if err != nil {
		re_act.Error()
	}

	// 初始化控制层
	err = handler.Init(conf)
	if err != nil {
		fmt.Println("handler初始化失败:", err)
	}

	e := echo.New()
	// 调试测试
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	// 添加跨域 cors
	e.Use(mw.CORS())

	// 实例 restfull
	e.Get("/pod", podGet)
	e.Post("/pod", podPost)
	e.Put("/pod", podPut)
	e.Patch("/pod", podPatch)
	e.Delete("/pod", podDelete)

	// 实例网络 restfull

	// RUN
	e.Run(fasthttp.New(":1234"))
}
