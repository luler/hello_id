package login

import (
	"github.com/gin-gonic/gin"
	"go_test/helper"
	"go_test/helper/response_helper"
	"go_test/model"
)

func Login(c *gin.Context) {
	type Param struct {
		Name     string `validate:"required,max=50" label:"账号"`
		Password string `validate:"required,max=50" label:"密码"`
	}
	var param Param
	helper.InputStruct(c, &param)
	var user model.User
	helper.Db().Where("name=?", param.Name).First(&user)
	if user.ID == 0 {
		helper.CommonException("用户不存在")
	}
	response_helper.Success(c, "登录成功", user)
}
