package main

import (
	"go_test/app"
	"go_test/app/logic"
)

func init() {
	//项目初始化
	app.InitApp()
}

func main() {
	//token := jwt_helper.GenerateToken(map[string]any{
	//	"uid": 1,
	//})
	//
	//time.Sleep(time.Second * 10)
	//fmt.Println(jwt_helper.ParseToken(token))

	//for i := 0; i < 1000000; i++ {
	//	log_helper.Info(i, " 1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	//}

	logic.FlushId()
}
