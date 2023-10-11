package main

import (
	"github.com/gin-gonic/gin"
	"go_test/app"
	"go_test/app/middleware"
	_ "go_test/docs"
	"go_test/route"
	"io/ioutil"
	"os"
)

func init() {
	//项目初始化
	app.InitApp("base", "migrate", "cron")
}

// @title 接口文档
// @version 1.0
// @description 当前页面用于展示项目一些开放的接口
// @termsOfService http://swagger.io/terms/
// @contact.name 开发人员
// @contact.url https://cas.luler.top/
// @contact.email 1207032539@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath
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
