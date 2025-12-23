package mysql

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	config "github.com/kangyueyue/go-ai-ddd/conf"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/message"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/session"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitMysql 初始化数据库
func InitMysql() error {
	host := config.GetConfig().MysqlHost
	port := config.GetConfig().MysqlPort
	dbname := config.GetConfig().MysqlDb
	username := config.GetConfig().MysqlUser
	password := config.GetConfig().MysqlPassword
	charset := config.GetConfig().MysqlCharset

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, dbname, charset)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, dbname, charset)

	var log logger.Interface
	if gin.Mode() == "debug" {
		log = logger.Default.LogMode(logger.Info)
	} else {
		log = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	return migration() // 自动建表
}

// migration 自动建表
func migration() error {
	return DB.AutoMigrate(
		new(user.UserPojo),
		new(message.MessagePojo),
		new(session.SessionPojo),
	)
}
