package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper/request_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/logic"
	"strconv"
)

// 获取ID
func GetId(c *gin.Context) {
	type Param struct {
		Type   string `validate:"required,alphanum" label:"规则标识"`
		Length any
	}
	var param Param
	request_helper.ParamGetStruct(c, &param)

	var length int
	if param.Length == nil {
		length = 1
	} else {
		length, _ = strconv.Atoi(param.Length.(string))
	}

	ids := []string{}
	ids = logic.GenerateId(param.Type, length)

	response_helper.Success(c, "获取成功", ids)
}
