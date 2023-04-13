package CommonController

import (
	"github.com/gin-gonic/gin"
	"go_test/helper/ResponseTool"
	"go_test/helper/SnowFlake"
	"strconv"
)

// 创建雪花ID
func CreateSnowflakeId(c *gin.Context, workers [1024]*SnowFlake.Worker) {
	worker_id, _ := strconv.Atoi(c.Query("worker_id"))
	if worker_id < 0 || worker_id > 1023 {
		ResponseTool.Fail(c, "机器ID只能设置在0-1023之间", new(struct{}))
		return
	}
	worker := workers[worker_id]
	ResponseTool.Success(c, "创建成功", map[string]interface{}{
		"worker_id": worker_id,
		"id":        worker.GetId(),
	})
}
