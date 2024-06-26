package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/miaogu-go/bluebell/controller"
	"github.com/miaogu-go/bluebell/dao/mysql"
	"github.com/miaogu-go/bluebell/dao/redis"
	"github.com/miaogu-go/bluebell/logger"
	"github.com/miaogu-go/bluebell/pkg/snowflake"
	"github.com/miaogu-go/bluebell/routes"
	"github.com/miaogu-go/bluebell/settings"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	//初始化配置文件
	settings.Init()

	//初始化日志
	logger.Init(settings.Conf)
	defer zap.L().Sync()

	//初始化mysql
	mysql.Init(settings.Conf.DbConf)
	defer mysql.Close()

	// 初始化redis
	redis.Init(settings.Conf.RedisConf)
	defer redis.Close()
	if err := snowflake.Init(settings.Conf.AppConf.StartTime, settings.Conf.AppConf.MachineId); err != nil {
		fmt.Printf("init snowflake failed,err:%#v\n", err)
		return
	}

	//初始化验证器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init translation failed,err:%#v\n", err)
		return
	}

	//初始化路由
	r := routes.Setup()
	//运行模式
	gin.SetMode(settings.Conf.AppConf.Mode)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("shutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("server shutdown", zap.Error(err))
	}
	zap.L().Info("server exiting")
}
