package route

import (
	"github.com/gin-gonic/gin"
	"go_test/controller/common_controller"
)

func InitRouter(e *gin.Engine) {
	e.Any("/snowflake", common_controller.Snowflake)
	e.Any("/sonyflake", common_controller.Sonyflake)
	e.Any("/uuid1", common_controller.Uuid1)
	e.Any("/uuid4", common_controller.Uuid4)
}
