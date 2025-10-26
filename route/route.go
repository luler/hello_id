package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_test/app/controller"
	"go_test/app/controller/common"
	"go_test/app/middleware"
	"net/http"
)

func InitRouter(e *gin.Engine) {
	//前端路由
	e.Static("/helloId", "./web/dist/helloId")
	e.NoRoute(func(context *gin.Context) {
		context.File("./web/dist/helloId/index.html")
	})
	e.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/helloId/idRule")
	})
	e.GET("/api/README.md", func(context *gin.Context) {
		context.File("./README.md")
	})
	api := e.Group("/api")
	//swagger页面
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.GET("/snowflake", middleware.AuthKey(), common.Snowflake)
	api.GET("/sonyflake", middleware.AuthKey(), common.Sonyflake)
	api.GET("/uuid1", middleware.AuthKey(), common.Uuid1)
	api.GET("/uuid4", middleware.AuthKey(), common.Uuid4)
	api.GET("/uuid7", middleware.AuthKey(), common.Uuid7)
	api.GET("/xid", middleware.AuthKey(), common.Xid)
	api.GET("/ksuid", middleware.AuthKey(), common.Ksuid)
	api.GET("/ulid", middleware.AuthKey(), common.Ulid)
	api.GET("/nanoid", middleware.AuthKey(), common.Nanoid)
	api.GET("/getId", middleware.AuthKey(), controller.GetId)
	//登录相关
	api.POST("/login", middleware.IpRateLimit(1, 1), controller.Login)
	api.GET("/casLogin", controller.CasLogin)
	auth := api.Group("", middleware.Auth())
	auth.POST("/resetPassword", controller.ResetPassword)
	auth.GET("/getUserInfo", controller.GetUserInfo)
	auth.POST("/saveAuthKey", controller.SaveAuthKey)
	auth.POST("/delAuthKey", controller.DelAuthKey)
	auth.GET("/getAuthKeyList", controller.GetAuthKeyList)
	auth.POST("/saveIdRule", controller.SaveIdRule)
	auth.POST("/delIdRule", controller.DelIdRule)
	auth.GET("/getIdRuleList", controller.GetIdRuleList)
	auth.GET("/createId", controller.GetId)
}
