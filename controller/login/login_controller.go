package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_test/helper"
	"go_test/helper/response_helper"
)

func Login(c *gin.Context) {
	type Param struct {
		Name     string `validate:"required,max=50" label:"账号"`
		Password string `validate:"required,max=50" label:"密码"`
		UserType int8   `validate:"required,oneof=1 2" label:"用户类型"`
	}
	var param Param

	helper.InputStruct(c, &param)
	fmt.Println(5555)
	response_helper.Success(c, "创建成功", param)
}
