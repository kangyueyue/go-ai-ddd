package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	config "github.com/kangyueyue/go-ai-ddd/conf"
	logger "github.com/kangyueyue/go-ai-ddd/infrastructure/common/log"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/container"
	mysql "github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/db"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/redis"
	"github.com/kangyueyue/go-ai-ddd/interfaces/adapter/initialize"
)

// main is the entry point for the application
func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	level := conf.MainConfig.Level
	appName := conf.MainConfig.AppName
	// init logger
	logger.InitLog(level, appName)

	// init gin mode
	switch level {
	case "test":
		gin.SetMode(gin.TestMode) // 测试环境
	case "info":
		gin.SetMode(gin.ReleaseMode) // 线上环境
	default:
		gin.SetMode(gin.DebugMode) // 开发环境
	}

	// init mysql
	if err := mysql.InitMysql(); err != nil {
		logger.Log.Errorf("InitMysql error %s", err.Error())
		return
	}
	// TODO:初始化aihelper
	// readDataFromDB()

	// init redis
	redis.Init()
	logger.Log.Infof("redis init success")

	// TODO:init mq
	// mq.InitRabbitMq()
	logger.Log.Infof("mq init success")

	// port param
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}

	// 初始化容器,要在mysql之后
	container.LoadingDomain()

	// start http server
	err := StartServer(host, port)
	if err != nil {
		panic(err)
	}
}

// StartServer 启动服务
func StartServer(addr string, port int) error {
	r := initialize.NewRouter()
	logger.Log.Infof("server start in port:%d", port)
	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}
