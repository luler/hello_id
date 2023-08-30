package response_tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, message string, data ...interface{}) {
	res := make(map[string]interface{})
	res["code"] = http.StatusOK
	res["message"] = message
	res["data"] = []int{}
	if len(data) > 0 {
		res["data"] = data[0]
	}
	c.JSON(res["code"].(int), res)
}

func Fail(c *gin.Context, message string, data ...interface{}) {
	res := make(map[string]interface{})
	res["code"] = http.StatusBadRequest
	res["message"] = message
	res["data"] = []int{}
	if len(data) > 0 {
		res["data"] = data[0]
	}
	c.JSON(res["code"].(int), res)
}
