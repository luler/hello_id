package main

import (
	"go_test/app"
	"os"
)

func init() {
	//项目初始化
	app.InitApp("base")
}

func main() {

	// 自动创建表
	app.InitApp("migrate")

	os.Exit(0)
}
