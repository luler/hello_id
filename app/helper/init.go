package helper

import (
	"github.com/joho/godotenv"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/log_helper"
)

// 项目启动初始化
func InitApp() {
	//加载.env配置
	godotenv.Load()
	//初始化DB
	db_helper.InitDb()
	//初始化日志记录方式
	log_helper.InitlogHelper()
}
