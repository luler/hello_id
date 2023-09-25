package cron_helper

import (
	"github.com/gogits/cron"
	"go_test/app/logic"
	"go_test/app/middleware"
	"sync"
)

var lock sync.Mutex

func InitCron() {
	c := cron.New()
	c.AddFunc("定时刷新ID到数据库", "*/1 * * * * ?", func() {
		if lock.TryLock() { //只能允许一个任务在跑
			defer lock.Unlock()
		} else {
			return
		}

		logic.FlushId()

	})
	c.AddFunc("定时清理ip限制缓存", "0 */1 * * * ?", func() {
		middleware.ClearIpRateLimit()
	})

	c.Start()
}
