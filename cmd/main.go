package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	config "github.com/kangyueyue/go-ai-ddd/conf"
	aihelper "github.com/kangyueyue/go-ai-ddd/infrastructure/common/aihepler"
	logger "github.com/kangyueyue/go-ai-ddd/infrastructure/common/log"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/container"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/mq"
	mysql "github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/db"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/message"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/redis"
	"github.com/kangyueyue/go-ai-ddd/interfaces/adapter"
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

	// init redis
	redis.Init()
	logger.Log.Infof("redis init success")

	// init mq
	mq.InitRabbitMq()
	logger.Log.Infof("mq init success")

	// port param
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}

	// 初始化容器,要在mysql之后
	container.LoadingDomain()

	// 初始化aihelper
	readDataFromDB()

	// start http server
	err := StartServer(host, port)
	if err != nil {
		panic(err)
	}
}

// StartServer 启动服务
func StartServer(addr string, port int) error {
	r := adapter.NewRouter()
	logger.Log.Infof("server start in port:%d", port)
	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}

// 从数据库加载消息并初始化 AIHelperManager
func readDataFromDB() error {
	manager := aihelper.GetGlobalManager()
	repo := message.NewMessageRepository(mysql.DB)
	msgs, err := repo.GetAllMessages()
	if err != nil {
		return err
	}
	// for
	for i := range msgs {
		m := &msgs[i]
		// 默认openai 模型
		modelType := "1"
		config := make(map[string]interface{})

		// 创建对应的AIHelper
		helper, err := manager.GetOrCreateAIHelper(m.UserName, m.SessionID, modelType, config)
		if err != nil {
			logger.Log.Infof("[readDataFromDB] failed to create helper for user=%s session=%s: %v", m.UserName, m.SessionID, err)
			continue
		}
		logger.Log.Info("readDataFromDB init:  ", helper.SessionID)
		helper.AddMessage(m.Content, m.UserName, m.IsUser, false)
	}
	logger.Log.Infof("AIHelperManager init success ")
	return nil
}
