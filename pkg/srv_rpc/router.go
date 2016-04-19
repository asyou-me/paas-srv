package main

import (
	//"github.com/asyoume/protorpc/protobuf"
	//gogo_proto "github.com/gogo/protobuf/proto"
	"fmt"
	"github.com/asyoume/protorpc"
	"net"
	"net/rpc"

	"flag"
	base_utils "github.com/asyoume/lib/utils"
	"github.com/asyoume/paas_srv/pkg/handler"
	"github.com/asyoume/paas_srv/pkg/re_act"
	"os"
	"os/signal"
	"syscall"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	// 信号
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 监控信号
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	//获取执行参数
	conf_path := flag.String("conf", "", "配置文件路径")
	flag.Parse()
	//获取配置文件的路径
	base_utils.CompletePath(conf_path)

	conf, err := base_utils.ConfigInit(*conf_path)
	if err != nil {
		re_act.Error()
	}

	fmt.Println(conf)

	//初始化控制层
	err = handler.Init()
	if err != nil {
		fmt.Println("handler初始化失败:", err)
	}

	// 注册 pod 服务
	podHandler := new(handler.PodHandler)
	err = rpc.RegisterName("Pod", podHandler)
	if err != nil {
		fmt.Println(err)
	}
	// 注册网络服务
	serHandler := new(handler.SerHandler)
	err = rpc.RegisterName("Ser", serHandler)
	if err != nil {
		fmt.Println(err)
	}
	// 注册 App 服务
	appHandler := new(handler.AppHandler)
	err = rpc.RegisterName("App", appHandler)
	if err != nil {
		fmt.Println(err)
	}
	// 注册 AppTemp 服务
	apptempHandler := new(handler.AppTempHandler)
	err = rpc.RegisterName("AppTemp", apptempHandler)
	if err != nil {
		fmt.Println(err)
	}

	//  监听端口
	l, e := net.Listen("tcp", ":1235")
	if e != nil {
		fmt.Println("listen error:", e)
	}

	// 开启服务 goroutine
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("conn:", err)
			}
			fmt.Println(err)
			go protorpc.ServeConn(conn)
		}
	}()

	// 服务退出操作
	<-done
	fmt.Println("exiting")
}
