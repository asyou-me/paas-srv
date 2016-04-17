package handler

import (
	"fmt"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func NewkubeClient() *client.Client {
	kubeclient, err := client.New(kubeConfig)
	if err != nil {
		fmt.Println(err)
	}
	return kubeclient
}

/*
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
*/
