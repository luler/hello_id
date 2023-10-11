package common

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"go_test/app/helper/exception_helper"
	"go_test/app/helper/response_helper"
	"strconv"
)

// @Summary 创建雪花ID
// @Description  创建雪花ID
// @Tags ID生成相关接口
// @Accept  json
// @Produce  json
// @Param workerId query int false "机器ID，取值范围0-1023，默认0"
// @Param length query int false "获取ID数量，默认1，最大500"
// @Success 200
// @Router /api/snowflake [get]
func Snowflake(c *gin.Context) {
	workerId, _ := strconv.Atoi(c.Query("workerId"))
	length, _ := strconv.Atoi(c.Query("length"))
	if workerId < 0 || workerId > 1023 {
		response_helper.Fail(c, "机器ID只能设置在0-1023之间")
		return
	}

	node, _ := snowflake.NewNode(int64(workerId))
	if length <= 0 {
		length = 1
	} else if length > 500 {
		exception_helper.CommonException("单次生成数量不能大于500")
	}
	var ids []string
	for i := 0; i < length; i++ {
		ids = append(ids, node.Generate().String())
	}

	response_helper.Success(c, "创建成功", gin.H{
		"type":     "snowflake",
		"workerId": workerId,
		"ids":      ids,
	})
}

// @Summary "创建sonyflake ID"
// @Description  "创建sonyflake ID"
// @Tags ID生成相关接口
// @Accept  json
// @Produce  json
// @Param length query int false "获取ID数量，默认1，最大500"
// @Success 200
// @Router /api/sonyflake [get]
func Sonyflake(c *gin.Context) {
	length, _ := strconv.Atoi(c.Query("length"))
	if length <= 0 {
		length = 1
	} else if length > 500 {
		exception_helper.CommonException("单次生成数量不能大于500")
	}
	sf, _ := sonyflake.New(sonyflake.Settings{})
	var ids []string
	var id uint64
	for i := 0; i < length; i++ {
		id, _ = sf.NextID()
		ids = append(ids, strconv.FormatUint(id, 10))
	}
	response_helper.Success(c, "创建成功", gin.H{
		"type": "sonyflake",
		"ids":  ids,
	})
}

// @Summary 生成uuid版本1类型的id
// @Description  生成uuid版本1类型的id
// @Tags ID生成相关接口
// @Accept  json
// @Produce  json
// @Param length query int false "获取ID数量，默认1，最大500"
// @Success 200
// @Router /api/uuid1 [get]
func Uuid1(c *gin.Context) {
	length, _ := strconv.Atoi(c.Query("length"))
	if length <= 0 {
		length = 1
	} else if length > 500 {
		exception_helper.CommonException("单次生成数量不能大于500")
	}
	var ids []string
	for i := 0; i < length; i++ {
		ids = append(ids, uuid.NewV1().String())
	}
	response_helper.Success(c, "创建成功", gin.H{
		"type": "uuid1",
		"ids":  ids,
	})
}

// @Summary 生成uuid版本4类型的id
// @Description  生成uuid版本4类型的id
// @Tags ID生成相关接口
// @Accept  json
// @Produce  json
// @Param length query int false "获取ID数量，默认1，最大500"
// @Success 200
// @Router /api/uuid4 [get]
func Uuid4(c *gin.Context) {
	length, _ := strconv.Atoi(c.Query("length"))
	if length <= 0 {
		length = 1
	} else if length > 500 {
		exception_helper.CommonException("单次生成数量不能大于500")
	}
	var ids []string
	for i := 0; i < length; i++ {
		ids = append(ids, uuid.NewV4().String())
	}
	response_helper.Success(c, "创建成功", gin.H{
		"type": "uuid4",
		"ids":  ids,
	})
}

// @Summary 生成Xid类型的id
// @Description  生成Xid类型的id
// @Tags ID生成相关接口
// @Accept  json
// @Produce  json
// @Param length query int false "获取ID数量，默认1，最大500"
// @Success 200
// @Router /api/xid [get]
func Xid(c *gin.Context) {
	length, _ := strconv.Atoi(c.Query("length"))
	if length <= 0 {
		length = 1
	} else if length > 500 {
		exception_helper.CommonException("单次生成数量不能大于500")
	}
	var ids []string
	for i := 0; i < length; i++ {
		ids = append(ids, xid.New().String())
	}
	response_helper.Success(c, "创建成功", gin.H{
		"type": "xid",
		"ids":  ids,
	})
}

// @Summary 生成Ksuid类型的id
// @Description  生成Ksuid类型的id
// @Tags ID生成相关接口
// @Accept  json
// @Produce  json
// @Param length query int false "获取ID数量，默认1，最大500"
// @Success 200
// @Router /api/ksuid [get]
func Ksuid(c *gin.Context) {
	length, _ := strconv.Atoi(c.Query("length"))
	if length <= 0 {
		length = 1
	} else if length > 500 {
		exception_helper.CommonException("单次生成数量不能大于500")
	}
	var ids []string
	for i := 0; i < length; i++ {
		ids = append(ids, ksuid.New().String())
	}
	response_helper.Success(c, "创建成功", gin.H{
		"type": "ksuid",
		"ids":  ids,
	})
}
