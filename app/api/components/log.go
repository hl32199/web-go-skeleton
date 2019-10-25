package components

import "web-go-skeleton/library/glog"

func init()  {
	//注册文件日志
	fileLogConfig := glog.FileOutputConfig{Path:"./logs/",FileNamePrefix:"app"}
	fileLogger := glog.NewFileOutput(fileLogConfig)
	glog.RegisterOutput(fileLogger)
	//注册标注错误日志
	stderr := glog.NewStderr()
	glog.RegisterOutput(stderr)
}
