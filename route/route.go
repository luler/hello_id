package route

import (
	"github.com/gin-gonic/gin"
	"go_test/app/controller"
	"go_test/app/controller/common"
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
	e.POST("/login", controller.Login)
	api := e.Group("/api", middleware.Auth())
	api.GET("/getUserInfo", controller.GetUserInfo)
	api.POST("/saveAuthKey", controller.SaveAuthKey)
	api.POST("/delAuthKey", controller.DelAuthKey)
	api.GET("/getAuthKeyList", controller.GetAuthKeyList)
	api.POST("/saveIdRule", controller.SaveIdRule)
	api.POST("/delIdRule", controller.DelIdRule)
	api.GET("/getIdRuleList", controller.GetIdRuleList)
	api_not_auth := e.Group("/api")
	api_not_auth.GET("/getId", middleware.AuthKey(), controller.GetId)
}
