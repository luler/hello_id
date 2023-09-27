package main

import (
	"github.com/gin-gonic/gin"
	"go_test/app"
	"go_test/app/middleware"
	"go_test/route"
	"io/ioutil"
	"os"
)

func init() {
	//项目初始化
	app.InitApp("base", "migrate", "cron")
}

func main() {
	gin.SetMode(os.Getenv(gin.EnvGinMode))
	//不输出请求日志
	gin.DefaultWriter = ioutil.Discard

	engine := gin.Default()
	//初始化中间件
	middleware.InitMiddleware(engine)
	//初始化路由
	route.InitRouter(engine)

	engine.Run() // listen and serve on 0.0.0.0:8080
}
