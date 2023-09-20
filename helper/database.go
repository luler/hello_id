package helper

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB

// 初始化数据库
func InitDb() {
	dsn := os.Getenv("DB_NAME") // 连接数据库
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.Getenv("TABLE_PREFIX"),
			SingularTable: true,
		},
	})

	DB = db
}

// 获取Db对象
func Db() *gorm.DB {
	return DB
}
