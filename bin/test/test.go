package main

import (
	"fmt"
	"go_test/helper"
	"time"
)

func init() {
	//项目初始化
	helper.InitApp()
}

func main() {
	token := helper.GenerateToken(map[string]any{
		"uid": 1,
	})

	time.Sleep(time.Second * 10)
	fmt.Println(helper.ParseToken(token))
}
