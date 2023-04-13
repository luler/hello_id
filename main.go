package main

import (
	"github.com/gin-gonic/gin"
	"go_test/controller/CommonController"
	"go_test/helper/SnowFlake"
)

func main() {
	r := gin.Default()
	var workers [1024]*SnowFlake.Worker
	for i := 0; i < 1024; i++ {
		w, _ := SnowFlake.NewWorker(int64(i))
		workers[i] = w
	}

	r.GET("/createSnowflakeId", func(c *gin.Context) {
		CommonController.CreateSnowflakeId(c, workers)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
