package main

import (
	"SpeakPeak/controller"
	"SpeakPeak/dao/mysql"
	"SpeakPeak/logger"
	"SpeakPeak/pkg/snowflake"
	"SpeakPeak/routers"
	"SpeakPeak/settings"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

func main() {
	//1.load config

	if err := settings.Init(); err != nil {
		fmt.Printf("Init Logger Failed,err:%v\n", err)
	}
	//2.init logger
	if err := logger.Init(settings.Conf.LogConfig, viper.GetString("app.mode")); err != nil {
		fmt.Printf("Init Logger Failed,err:%v\n", err)
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3.init Mysql connection
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("Init Mysql Failed,err:%v\n", err)
	}
	defer mysql.Close()

	//4.init redis connection
	//if err := redis.Init(); err != nil {
	//	fmt.Printf("Init Redis Failed,err:%v\n", err)
	//	return
	//}

	//初始化内置的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err%v\n", err)
		return
	}

	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Printf("Init Snowflake Failed,err:%v\n", err)
		return
	}
	//5.register router
	r := routers.Setup()

	//run
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	//开启一个goroutine启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	//优雅关机
	quit := make(chan os.Signal, 1)                      //创建一个接受信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //接收信号
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
