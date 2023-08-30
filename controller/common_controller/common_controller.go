package common_controller

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go_test/helper/response_tool"
	"go_test/helper/snowflake"
	"strconv"
)

// CreateSnowflakeId 创建雪花ID
func CreateSnowflakeId(c *gin.Context, workers [1024]*snowflake.Worker) {
	worker_id, _ := strconv.Atoi(c.Query("worker_id"))
	if worker_id < 0 || worker_id > 1023 {
		response_tool.Fail(c, "机器ID只能设置在0-1023之间")
		return
	}
	worker := workers[worker_id]
	response_tool.Success(c, "创建成功", gin.H{
		"type":      "snowflake",
		"worker_id": worker_id,
		"id":        worker.GetId(),
	})
}

// Uuid1 生成uuid版本1类型的id
func Uuid1(c *gin.Context) {
	uuid1 := uuid.NewV1()
	response_tool.Success(c, "创建成功", gin.H{
		"type": "uuid1",
		"id":   uuid1.String(),
	})
}

// Uuid4 生成uuid版本4类型的id
func Uuid4(c *gin.Context) {
	uuid4 := uuid.NewV4()
	response_tool.Success(c, "创建成功", gin.H{
		"type": "uuid4",
		"id":   uuid4.String(),
	})
}
