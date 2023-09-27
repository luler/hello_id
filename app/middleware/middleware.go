package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_test/app/helper/cache_helper"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/exception_helper"
	"go_test/app/helper/jwt_helper"
	"go_test/app/helper/request_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/model"
	"net/http"
	"reflect"
	"strings"
	"time"
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
				if _, ok := r.(exception_helper.MyException); ok {
					r := r.(exception_helper.MyException)
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
			param := request_helper.Input(context, "token")
			if value, exists := param["token"]; exists {
				token = value.(string)
			}
		}
		token = strings.Replace(token, "Bearer ", "", -1)
		if token == "" {
			exception_helper.CommonException("请登录", http.StatusUnauthorized)
		}
		claims := jwt_helper.ParseToken(token)

		data := claims["data"].(map[string]any)
		context.Set("uid", data["uid"])

		context.Next()
	}
}

// 授权码中间件
func AuthKey() gin.HandlerFunc {
	return func(context *gin.Context) {
		type Param struct {
			AuthKey string `validate:"required,uuid4" label:"授权码"`
		}
		var param Param
		request_helper.InputStruct(context, &param)

		key := "AuthKey:" + param.AuthKey
		if _, found := cache_helper.GoCache().Get(key); !found {
			var authKey model.AuthKey
			if err := db_helper.Db().Where("auth_key=?", param.AuthKey).First(&authKey).Error; err != nil {
				exception_helper.CommonException("授权码不存在")
			}
			//缓存24小时
			cache_helper.GoCache().Set(key, authKey, time.Hour*24)
		}

		context.Next()
	}
}
