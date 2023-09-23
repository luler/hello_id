package logic

import (
	"go_test/app/helper"
	"go_test/app/helper/cache_helper"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/log_helper"
	"go_test/app/model"
	"strconv"
	"sync"
	"time"
)

var lock1 sync.Mutex
var lock2 sync.Mutex

func GenerateId(idRuleType string, length int) []string {
	lock1.Lock()
	defer lock1.Unlock()

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
	go func() {
		lock2.Lock()
		defer lock2.Unlock()
		var idRule2 model.IdRule
		if ir2, found2 := cache_helper.GoCache().Get(key); found2 {
			idRule2 = ir2.(model.IdRule)
		} else { //缓存不存在不处理
			return
			log_helper.Debug("缓存不存在不处理")
		}
		if idRule.CurrentId != idRule2.CurrentId { //乐观锁，不相同不处理
			return
			log_helper.Debug("乐观锁，不相同不处理")
		}

		//异步更新数据库
		if err := db_helper.Db().Where("id=?", idRule.Id).Updates(&model.IdRule{
			CurrentId: idRule.CurrentId,
		}).Error; err != nil {
			helper.CommonException("规则标识不存在")
		}
	}()
	return ids
}
