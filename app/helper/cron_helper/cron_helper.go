package cron_helper

import (
	"github.com/gogits/cron"
	"go_test/app/middleware"
	"sync"
)

var lock sync.Mutex

func InitCron() {
	c := cron.New()
	c.AddFunc("定时清理ip限制缓存", "0 */1 * * * ?", func() {
		middleware.ClearIpRateLimit()
	})

	c.Start()
}
