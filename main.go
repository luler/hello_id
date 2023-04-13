package main

import (
	"github.com/gin-gonic/gin"
	"go_test/controller/common_controller"
	"go_test/helper/snowflake"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	var workers [1024]*snowflake.Worker
	for i := 0; i < 1024; i++ {
		w, _ := snowflake.NewWorker(int64(i))
		workers[i] = w
	}

	r.GET("/create", func(c *gin.Context) {
		common_controller.CreateSnowflakeId(c, workers)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
