package route

import (
	"github.com/gin-gonic/gin"
	"go_test/controller/common_controller"
	"go_test/helper/snowflake"
)

func InitRouter(e *gin.Engine) {
	var workers [1024]*snowflake.Worker
	for i := 0; i < 1024; i++ {
		w, _ := snowflake.NewWorker(int64(i))
		workers[i] = w
	}
	e.Any("/snowflake", func(c *gin.Context) {
		common_controller.CreateSnowflakeId(c, workers)
	})

	e.Any("/uuid1", common_controller.Uuid1)
	e.Any("/uuid4", common_controller.Uuid4)
}
