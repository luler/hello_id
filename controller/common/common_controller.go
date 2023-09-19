package common

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"go_test/helper/response_helper"
	"strconv"
)

// Snowflake 创建雪花ID
func Snowflake(c *gin.Context) {
	worker_id, _ := strconv.Atoi(c.Query("worker_id"))
	if worker_id < 0 || worker_id > 1023 {
		response_helper.Fail(c, "机器ID只能设置在0-1023之间")
		return
	}

	node, _ := snowflake.NewNode(int64(worker_id))

	response_helper.Success(c, "创建成功", gin.H{
		"type":      "snowflake",
		"worker_id": worker_id,
		"id":        node.Generate(),
	})
}

// Sonyflake 创建sonyflake ID
func Sonyflake(c *gin.Context) {
	sf, _ := sonyflake.New(sonyflake.Settings{})
	id, _ := sf.NextID()
	response_helper.Success(c, "创建成功", gin.H{
		"type": "sonyflake",
		"id":   strconv.FormatUint(id, 10),
	})
}

// Uuid1 生成uuid版本1类型的id
func Uuid1(c *gin.Context) {
	uuid1 := uuid.NewV1()
	response_helper.Success(c, "创建成功", gin.H{
		"type": "uuid1",
		"id":   uuid1.String(),
	})
}

// Uuid4 生成uuid版本4类型的id
func Uuid4(c *gin.Context) {
	uuid4 := uuid.NewV4()
	response_helper.Success(c, "创建成功", gin.H{
		"type": "uuid4",
		"id":   uuid4.String(),
	})
}

// Xid 生成Xid类型的id
func Xid(c *gin.Context) {
	guid := xid.New()
	response_helper.Success(c, "创建成功", gin.H{
		"type": "xid",
		"id":   guid.String(),
	})
}

// Ksuid 生成Ksuid类型的id
func Ksuid(c *gin.Context) {
	guid := ksuid.New()
	response_helper.Success(c, "创建成功", gin.H{
		"type": "ksuid",
		"id":   guid.String(),
	})
}
