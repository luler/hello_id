package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper/exception_helper"
	"go_test/app/helper/request_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/logic"
	"strconv"
)

// @Summary 生成自定义ID
// @Description  生成自定义ID
// @Tags ID生成相关接口
// @Accept  json
// @Produce  json
// @Param type query string true "ID标识"
// @Param authKey query string true "授权码"
// @Param length query int false "获取ID数量，默认1，最大500"
// @Success 200
// @Router /api/getId [get]
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

	if length > 500 {
		exception_helper.CommonException("单次生成数量不能大于500")
	}

	ids := []string{}
	ids = logic.GenerateId(param.Type, length)

	response_helper.Success(c, "获取成功", ids)
}
