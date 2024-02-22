package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DBClient *gorm.DB
var dbLogFile, _ = os.Create("logs/db.log")

func InitDataBase() {
	newLogger := logger.New(log.New(dbLogFile, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             2 * time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	})

	var err error
	DBClient, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123@tcp(127.0.0.1:3306)/tools_platform?charset=utf8mb4&parseTime=True&loc=Local",
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
