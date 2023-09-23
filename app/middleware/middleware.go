package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_test/app/helper"
	"go_test/app/helper/jwt_helper"
	"go_test/app/helper/response_helper"
	"net/http"
	"reflect"
	"strings"
)

// 初始化全局中间件
func InitMiddleware(e *gin.Engine) {
	//异常捕获中间件
	e.Use(Exception())
}

// 异常捕获中间件
func Exception() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(helper.MyException); ok {
					r := r.(helper.MyException)
					response_helper.Common(context, r.Code, r.Message, r.Data)
				} else {
					response_helper.Common(context, http.StatusInternalServerError, fmt.Sprintf("%v", reflect.ValueOf(r)))
				}
				context.Abort() //终止请求
			}
		}()

		context.Next()
	}
}

// 授权校验中间件
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			param := helper.Input(context, "token")
			if value, exists := param["token"]; exists {
				token = value.(string)
			}
		}
		token = strings.Replace(token, "Bearer ", "", -1)
		if token == "" {
			helper.CommonException("请登录", http.StatusUnauthorized)
		}
		claims := jwt_helper.ParseToken(token)

		data := claims["data"].(map[string]any)
		context.Set("uid", data["uid"])

		context.Next()
	}
}
