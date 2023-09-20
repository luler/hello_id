package helper

import (
	"github.com/joho/godotenv"
)

// 项目启动初始化
func InitApp() {
	//加载.env配置
	godotenv.Load()
	//初始化DB
	InitDb()
}
