package middleware

import (
	"github.com/gin-gonic/gin"
	"go_test/helper"
	"go_test/helper/response_helper"
	"net/http"
)

// 初始化全局中间件
func InitMiddleware(e *gin.Engine) {
	//异常捕获中间件
	e.Use(ExceptionMiddleware())
}

// 异常捕获中间件
func ExceptionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(helper.MyException); ok {
					r := r.(helper.MyException)
					response_helper.Common(context, r.Code, r.Message, r.Data)
				} else {
					response_helper.Common(context, http.StatusInternalServerError, r.(string))
				}
			}
		}()

		context.Next()
	}
}
