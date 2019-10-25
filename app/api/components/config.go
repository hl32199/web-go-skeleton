package components

import (
	"github.com/Unknwon/goconfig"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	Config *goconfig.ConfigFile
)

func init(){
	var err error
	configFile := fmt.Sprintf("app/api/conf/%s.conf",gin.Mode())
	Config, err = goconfig.LoadConfigFile(configFile)    //加载配置文件
	if err != nil {
		panic(fmt.Sprintf("get config file %s error:%s",configFile,err))
	}
}