package main

import (
	"fmt"
	"go_test/helper"
	"go_test/helper/jwt_helper"
	"time"
)

func init() {
	//项目初始化
	helper.InitApp()
}

func main() {
	token := jwt_helper.GenerateToken(map[string]any{
		"uid": 1,
	})

	time.Sleep(time.Second * 10)
	fmt.Println(jwt_helper.ParseToken(token))
}
