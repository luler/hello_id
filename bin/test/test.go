package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"go_test/app"
	"go_test/app/helper/cache_helper"
)

func init() {
	//项目初始化
	app.InitApp("base")
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

	//logic.FlushId()

	//log_helper.Error("你是谁")
	cache_helper.GoCache().Set("test", 1, cache.NoExpiration)
	cache_helper.GoCache().Increment("test", 1)
	cache_helper.GoCache().Increment("test", 1)
	cache_helper.GoCache().Increment("test", 1)

	fmt.Println(cache_helper.GoCache().Get("test"))

}
