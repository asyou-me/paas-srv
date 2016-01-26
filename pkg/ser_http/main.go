package main

import (
	"flag"
	"fmt"
	"github.com/asyoume/PaasControl/pkg/handler"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"os"
	"path/filepath"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	//获取执行参数
	conf_path := flag.String("conf", "", "配置文件路径")
	flag.Parse()
	//获取配置文件的路径
	complete_path(conf_path)
	//初始化控制层
	err := handler.Init(*conf_path)
	if err != nil {
		fmt.Println("handler初始化失败:", err)
	}

	e := echo.New()
	//调试测试
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	//添加跨域cors
	e.Use(cors.Default().Handler)

	e.Post("/Pod/Put", PodPut)

	// Routes
	e.Run(":9091")
}

func complete_path(conf_path *string) {
	if len((*conf_path)) > 0 && string((*conf_path)[0]) != "/" {
		var curr_path string = ""
		curr_path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		*conf_path = curr_path + "/" + *conf_path
	}
}
