package route

import (
	"github.com/gin-gonic/gin"
	"go_test/app/controller/auth_key"
	"go_test/app/controller/common"
	"go_test/app/controller/login"
	"go_test/app/controller/user"
	"go_test/app/middleware"
)

func InitRouter(e *gin.Engine) {
	e.Any("/snowflake", common.Snowflake)
	e.Any("/sonyflake", common.Sonyflake)
	e.Any("/uuid1", common.Uuid1)
	e.Any("/uuid4", common.Uuid4)
	e.Any("/xid", common.Xid)
	e.Any("/ksuid", common.Ksuid)
	//登录相关
	e.POST("/login", login.Login)
	api := e.Group("/api", middleware.AuthMiddleware())
	api.GET("/getUserInfo", user.GetUserInfo)
	api.POST("/saveAuthKey", auth_key.SaveAuthKey)
	api.GET("/getAuthKeyList", auth_key.GetAuthKeyList)
}
