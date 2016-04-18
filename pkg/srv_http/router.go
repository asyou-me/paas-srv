package main

import (
	"flag"
	"fmt"
	base_utils "github.com/asyoume/lib/utils"
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/re_act"
	"github.com/asyoume/paas_srv/pkg/re_act/types"
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
		os.Exit(2)
	}

	// 初始化系统日志
	re_act.InitLog(conf.MicroSer["log1"])

	// 记录系统日志
	log := types.NewSystemLog()
	log.Type = "system"
	log.Msg = "start "
	re_act.Info(log)

	// 初始化控制层
	err = handler.Init()
	if err != nil {
		fmt.Println("handler初始化失败:", err)
		os.Exit(2)
	}

	// 选择数据序列化方式
	selectProto("proto")

	// echo引擎
	e := echo.New()
	// 调试测试
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	// 添加跨域 cors
	e.Use(mw.CORS())

	// 容器
	//e.Get("/pod", podGet)
	//e.Post("/pod", podPost)
	//e.Put("/pod", podPut)
	//e.Patch("/pod", podPatch)
	//e.Delete("/pod", podDelete)

	// 容器网络
	//e.Get("/ser", podGet)
	//e.Post("/ser", podPost)
	//e.Put("/ser", podPut)
	//e.Patch("/ser", podPatch)
	//e.Delete("/ser", podDelete)

	// 应用
	e.Get("/app", appGet)
	e.Post("/app", appPost)
	e.Put("/app", podPut)
	e.Patch("/app", podPatch)
	e.Delete("/app", podDelete)

	// 应用模板
	e.Get("/app/temp", appTempGet)
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

	// 版本信息
	e.Get("/version", version)
	e.Get("/version/info", info)

	// 运行http服务器
	e.Run(fasthttp.New(":1234"))
}
