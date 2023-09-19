package login

import (
	"github.com/gin-gonic/gin"
	"go_test/helper"
	"go_test/helper/response_helper"
)

func Login(c *gin.Context) {
	//param := helper.Input(c, "name", "password")
	type Param struct {
		Name     string      `validate:"required,max=50" label:"账号"`
		Password string      `validate:"required,max=50" label:"密码"`
		UserType interface{} `validate:"required,oneof=1 2" label:"用户类型"`
	}
	var param Param
	//mapstructure.Decode(param, &data)

	err := helper.InputStruct(c, &param)
	if err != nil {
		response_helper.Fail(c, err.(string))
		return
	}
	response_helper.Success(c, "创建成功", param)
}
