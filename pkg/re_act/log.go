package re_act

import (
	"fmt"
	"os"

	"github.com/asyoume/lib/log_client"
	"github.com/asyoume/lib/pulic_type"
)

var Log *log_client.Logger

// log初始化
func InitLog(conf pulic_type.MicroSerType) {
	var err error
	Log, err = log_client.New(&[]log_client.LogConf{
		log_client.LogConf{ // 主log服务
			Addr:   conf.Addr,
			Area:   conf.Attr["dir"].(string),
			Type:   conf.Attr["type"].(string),
			Spare:  false,
			Weight: 1,
		},
	})

	// 当日志服务不可用,取消启动服务器
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

//传入debug日志
func Debug(obj ...log_client.LogBase) {
	Log.Debug(obj...)
}

//传入info日志
func Info(obj ...log_client.LogBase) {
	Log.Info(obj...)
}

//传入Print日志
func Print(obj ...log_client.LogBase) {
	Log.Print(obj...)
}

//传入Warn日志
func Warn(obj ...log_client.LogBase) {
	Log.Warn(obj...)
}

//传入Error日志
func Error(obj ...log_client.LogBase) {
	Log.Error(obj...)
}

//传入Fatal日志
func Fatal(obj ...log_client.LogBase) {
	Log.Fatal(obj...)
}
