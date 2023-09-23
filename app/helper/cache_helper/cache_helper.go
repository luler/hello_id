package cache_helper

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var c *cache.Cache
var once sync.Once

// 获取单例
func GoCache() *cache.Cache {
	once.Do(func() {
		c = cache.New(time.Minute*5, time.Minute*10)
	})
	return c
}
