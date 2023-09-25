package middleware

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
)

// 创建速率限制器
type IpRateLimitStruct struct {
	Limiter   *rate.Limiter
	UpdatedAt time.Time
}

var limiters = make(map[string]*IpRateLimitStruct)

// 限制速度
func IpRateLimit(r int, b int) gin.HandlerFunc {
	return func(context *gin.Context) {
		ip := context.ClientIP()
		key := ip + "_" + strconv.Itoa(r) + "_" + strconv.Itoa(b)
		limiter, exist := limiters[key]
		if !exist {
			limiter = &IpRateLimitStruct{
				Limiter:   rate.NewLimiter(rate.Limit(r), b),
				UpdatedAt: time.Now(),
			}
			limiters[key] = limiter
		}
		if limiter.Limiter.Allow() {
			limiter.UpdatedAt = time.Now()
			context.Next()
		} else {
			helper.CommonException("请求过于频繁，请稍后重试", http.StatusTooManyRequests)
		}
	}
}

// 定时清理ip限制缓存
func ClearIpRateLimit() {
	for key, limiter := range limiters {
		if time.Since(limiter.UpdatedAt) > time.Minute {
			delete(limiters, key)
		}
	}
}
