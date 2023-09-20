package login

import (
	"github.com/gin-gonic/gin"
	"go_test/helper"
	"go_test/helper/response_helper"
	"go_test/model"
	"os"
)

func Login(c *gin.Context) {
	type Param struct {
		Name     string `validate:"required,max=50" label:"账号"`
		Password string `validate:"required,max=50" label:"密码"`
	}
	var param Param
	helper.InputStruct(c, &param)
	var user model.User
	err := helper.Db().Where("name=?", param.Name).First(&user)
	if err.Error != nil {
		if param.Name == os.Getenv("ADMIN_NAME") {
			user.Name = param.Name
			user.Password = os.Getenv("ADMIN_PASSWORD")
			user.Status = 1
			res := helper.Db().Save(&user)
			if res.Error != nil {
				helper.CommonException("系统异常")
			}
			helper.Db().Where("name=?", param.Name).First(&user)
		} else {
			helper.CommonException("用户不存在")
		}
	}

	response_helper.Success(c, "登录成功", user)
}
