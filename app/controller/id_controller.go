package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper"
	"go_test/app/helper/response_helper"
	"strconv"
)

// 获取ID
func GetId(c *gin.Context) {
	type Param struct {
		Type   string `validate:"required,alphanum" label:"规则标识"`
		Length any
	}
	var param Param
	helper.InputStruct(c, &param)
	length, _ := strconv.Atoi(param.Length.(string))
	if param.Length == nil || param.Length == 0 {
		length = 1
	}
	ids := []string{}
	for i := 0; i < length; i++ {

	}

	response_helper.Success(c, "获取成功", ids)
}
