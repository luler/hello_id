package db_helper

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

var db *gorm.DB

// 初始化数据库
func InitDb() {
	dsn := os.Getenv("DB_NAME") // 连接数据库
	openDb, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.Getenv("TABLE_PREFIX"),
			SingularTable: true,
		},
	})

	db = openDb
}

// 获取Db对象
func Db() *gorm.DB {
	return db
}
