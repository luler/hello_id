package helper

import (
	"github.com/joho/godotenv"
	"go_test/helper/db_helper"
)

// 项目启动初始化
func InitApp() {
	//加载.env配置
	godotenv.Load()
	//初始化DB
	db_helper.InitDb()
}
