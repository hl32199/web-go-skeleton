package components

import "web-go-skeleton/library/glog"

func InitLogger()  {
	//注册文件日志
	path := Config.MustValue("log", "path","./logs/")
	fileNamePrefix := Config.MustValue("log", "file_name_prefix","app")
	fileLogConfig := glog.FileOutputConfig{Path:path,FileNamePrefix:fileNamePrefix}
	fileLogger := glog.NewFileOutput(fileLogConfig)
	glog.RegisterOutput(fileLogger)
	//注册标准输出日志
	stdout := glog.NewStdout()
	glog.RegisterOutput(stdout)
}
