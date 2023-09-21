package main

import (
	"go_test/helper"
	"go_test/helper/db_helper"
	"go_test/model"
	"os"
)

func init() {
	//项目初始化
	helper.InitApp()
}

func main() {

	// 自动创建表
	db_helper.Db().AutoMigrate(&model.User{})

	os.Exit(0)
}
