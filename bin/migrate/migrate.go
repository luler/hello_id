package main

import (
	"go_test/app"
	"go_test/app/helper/db_helper"
	"go_test/app/model"
	"os"
)

func init() {
	//项目初始化
	app.InitApp()
}

func main() {

	// 自动创建表
	db_helper.Db().AutoMigrate(&model.User{}, &model.AuthKey{}, &model.IdRule{})

	os.Exit(0)
}
