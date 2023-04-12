package main

import (
	"github.com/gin-gonic/gin"
	"go_test/controller"
	"go_test/helper/SnowFlake"
)

func main() {
	r := gin.Default()
	cc := new(controller.CommonController)
	worker, _ := SnowFlake.NewWorker(1)
	r.GET("/ping", func(c *gin.Context) {
		cc.CreateSnowflakeId(c, worker)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
