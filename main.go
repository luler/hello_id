package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go_test/route"
	"io/ioutil"
	"os"
)

func init() {
	//加载.env配置
	godotenv.Load()
}

func main() {
	gin.SetMode(os.Getenv(gin.EnvGinMode))
	//不输出请求日志
	gin.DefaultWriter = ioutil.Discard

	engine := gin.Default()

	route.InitRouter(engine)

	engine.Run() // listen and serve on 0.0.0.0:8080
}
