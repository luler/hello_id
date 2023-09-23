package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/model"
)

// 获取登录用户信息
func GetUserInfo(c *gin.Context) {

	var user model.User

	uid, _ := c.Get("uid")
	err := db_helper.Db().Where("id=?", uid).Omit("password").First(&user)
	if err.Error != nil {
		helper.CommonException("用户不存在")
	}

	response_helper.Success(c, "获取成功", user)
}
