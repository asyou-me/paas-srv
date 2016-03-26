package handler

import (
	"github.com/asyoume/lib/pulic_type"
	"github.com/asyoume/paas_srv/pkg/types"
	"github.com/asyoume/paas_srv/pkg/utils"
	"k8s.io/kubernetes/pkg/client/restclient"
)

var kubeConfig *restclient.Config

func Init(conf *pulic_type.ConfType) error {
	kubeConfig = &restclient.Config{
		Host:     "http://115.29.113.249:8080",
		Username: "test",
		Password: "password",
	}

	utils.InitLog(conf.MicroSer["log1"])

	utils.Log.PrintKey = true

	log := types.NewSystemLog()
	log.Type = "system"

	//models.Init(conf.MicroSer["db"], conf.MicroSer["dblog"])
	/*if err != nil {
		utils.Error(log)
		return err
	}*/

	log.Msg = "start "
	utils.Info(log)

	return nil
}

/*ajax返回数据*/
type AjaxRel struct {
	Status int
	Mag    string
	Data   string
}
