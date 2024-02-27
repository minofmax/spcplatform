package dao

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DBClient *gorm.DB

func InitDataBase() {
	rotationLogger := &lumberjack.Logger{
		Filename:   "logs/db.log",
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
		LocalTime:  true, // 启用本地时区
	}

	newLogger := logger.New(log.New(rotationLogger, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             2 * time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	})

	var err error
	DBClient, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:shopee_openscan@tcp(127.0.0.1:3306)/tools_platform?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
}
