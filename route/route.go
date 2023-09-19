package route

import (
	"github.com/gin-gonic/gin"
	"go_test/controller/common"
	"go_test/controller/login"
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
}
