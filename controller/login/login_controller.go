package login

import (
	"github.com/gin-gonic/gin"
	"go_test/helper"
	"go_test/helper/db_helper"
	"go_test/helper/jwt_helper"
	"go_test/helper/response_helper"
	"go_test/model"
	"golang.org/x/crypto/bcrypt"
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
	err := db_helper.Db().Where("name=?", param.Name).First(&user)
	if err.Error != nil {
		if param.Name == os.Getenv("ADMIN_NAME") {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
			user.Name = param.Name
			user.Password = string(hashedPassword)
			user.Status = 1
			res := db_helper.Db().Save(&user)
			if res.Error != nil {
				helper.CommonException("系统异常")
			}
			db_helper.Db().Where("name=?", param.Name).First(&user)
		} else {
			helper.CommonException("用户不存在")
		}
	}

	//判断密码是否一致
	passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if passwordErr != nil {
		helper.CommonException("密码输入有误")
	}
	//是否启用
	if user.Status != 1 {
		helper.CommonException("用户已被禁用")
	}

	res := jwt_helper.IssueToken(gin.H{
		"uid": user.ID,
	})

	response_helper.Success(c, "登录成功", res)
}
