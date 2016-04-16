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

	// echo引擎
	e := echo.New()
	// 调试测试
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	// 添加跨域 cors
	e.Use(mw.CORS())

	// 容器
	e.Get("/pod", podGet)
	e.Post("/pod", podPost)
	e.Put("/pod", podPut)
	e.Patch("/pod", podPatch)
	e.Delete("/pod", podDelete)

	// 容器网络
	e.Get("/ser", podGet)
	e.Post("/ser", podPost)
	e.Put("/ser", podPut)
	e.Patch("/ser", podPatch)
	e.Delete("/ser", podDelete)

	// 应用
	e.Get("/app", podGet)
	e.Post("/app", podPost)
	e.Put("/app", podPut)
	e.Patch("/app", podPatch)
	e.Delete("/app", podDelete)

	// 应用模板
	e.Get("/app/temp", podGet)
	//e.Post("/app/temp", podPost)
	//e.Put("/app/temp", podPut)
	//e.Patch("/app/temp", podPatch)
	//e.Delete("/app/temp", podDelete)

	// 用户
	e.Get("/user", podGet)
	e.Post("/user", podPost)
	e.Put("/user", podPut)
	e.Patch("/user", podPatch)
	e.Delete("/user", podDelete)

	// 运行http服务器
	e.Run(fasthttp.New(":1234"))
}
