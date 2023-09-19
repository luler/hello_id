package login

import (
	"github.com/gin-gonic/gin"
	"go_test/helper"
	"go_test/helper/response_helper"
)

func Login(c *gin.Context) {
	response_helper.Success(c, "创建成功", helper.Input(c, "name", "password"))
}
