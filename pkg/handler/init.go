package handler

import (
	"fmt"
	//"k8s.io/kubernetes/pkg/api"
	"github.com/Shopify/sarama"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

var kubeConfig *client.Config
var AccessLogProducer *sarama.AsyncProducer

func Init(conf_path string) error {
	kubeConfig = &client.Config{
		Host:     "http://localhost:8080",
		Username: "test",
		Password: "password",
	}
	//kafka log
	//AccessLogProducer = newLogProducer([]string{"10.64.3.140:9092"})
	conf, err := ConfigInit(conf_path)
	if err != nil {
		return err
	}
	fmt.Println("conf:", conf)
	//models.Init(conf.MicroSer["db"], conf.MicroSer["dblog"])
	if err != nil {
		return err
	}
	return nil
}

/*ajax返回数据*/
type AjaxRel struct {
	Status int
	Mag    string
	Data   string
}
