package main

import (
	"github.com/joho/godotenv"
	"go_test/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

func init() {
	//加载.env配置
	godotenv.Load()
}

func main() {
	dsn := os.Getenv("DB_NAME") // 连接数据库
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.Getenv("TABLE_PREFIX"),
			SingularTable: true,
		},
	})

	// 自动创建表
	db.AutoMigrate(&model.User{})

	os.Exit(0)
}
