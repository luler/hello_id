package logic

import (
	"go_test/app/helper"
	"go_test/app/helper/cache_helper"
	"go_test/app/helper/db_helper"
	"go_test/app/model"
	"strconv"
	"sync"
	"time"
)

var lock sync.Mutex

func GenerateId(idRuleType string, length int) []string {
	lock.Lock()
	defer lock.Unlock()

	key := "IdRuleType:" + idRuleType
	var idRule model.IdRule
	if ir, found := cache_helper.GoCache().Get(key); !found {
		if err := db_helper.Db().Where("type=?", idRuleType).First(&idRule).Error; err != nil {
			helper.CommonException("规则标识不存在")
		}
		//缓存24小时
		cache_helper.GoCache().Set(key, idRule, time.Hour*24)
	} else {
		idRule = ir.(model.IdRule)
	}
	ids := []string{}
	for i := 0; i < length; i++ {
		idRule.CurrentId += 1
		ids = append(ids, strconv.Itoa(idRule.CurrentId))
	}
	//缓存24小时
	cache_helper.GoCache().Set(key, idRule, time.Hour*24)
	return ids
}
