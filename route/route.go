package route

import (
	"github.com/gin-gonic/gin"
	"go_test/app/controller"
	"go_test/app/controller/common"
	"go_test/app/middleware"
)

func InitRouter(e *gin.Engine) {
	api := e.Group("/api")
	api.Any("/snowflake", common.Snowflake)
	api.Any("/sonyflake", common.Sonyflake)
	api.Any("/uuid1", common.Uuid1)
	api.Any("/uuid4", common.Uuid4)
	api.Any("/xid", common.Xid)
	api.Any("/ksuid", common.Ksuid)
	//登录相关
	api.POST("/login", middleware.IpRateLimit(1, 1), controller.Login)
	auth := api.Group("", middleware.Auth())
	auth.POST("/resetPassword", controller.ResetPassword)
	auth.GET("/getUserInfo", controller.GetUserInfo)
	auth.POST("/saveAuthKey", controller.SaveAuthKey)
	auth.POST("/delAuthKey", controller.DelAuthKey)
	auth.GET("/getAuthKeyList", controller.GetAuthKeyList)
	auth.POST("/saveIdRule", controller.SaveIdRule)
	auth.POST("/delIdRule", controller.DelIdRule)
	auth.GET("/getIdRuleList", controller.GetIdRuleList)
	api.GET("/getId", middleware.AuthKey(), controller.GetId)
}
