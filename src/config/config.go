package config

import (
	"log"
)

type AppConfig struct {
	AppName     string
	Port        int
	Description string
}

type MysqlConfig struct {
	AppName string
	Port    int
}

func init() {
	//设置前缀
	//log.SetPrefix("MYLOG: ")
	//对log输出格式进行配置：日期|时间|文件路径
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}
