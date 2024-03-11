package main

import (
	"SpeakPeak/dao/mysql"
	"SpeakPeak/logger"
	"SpeakPeak/settings"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	//1.load config
	if err := settings.Init(); err != nil {
		fmt.Printf("Init Logger Failed,err:%v\n", err)
	}
	//2.init logger
	if err := logger.Init(); err != nil {
		fmt.Printf("Init Logger Failed,err:%v\n", err)
	}
	zap.L().Debug("logger init success...")
	//3.init Mysql connection
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init Mysql Failed,err:%v\n", err)
	}
	//4.init redis connection

	//5.init etcd connection

	//6.init server

}
