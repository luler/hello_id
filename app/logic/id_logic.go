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

var lock sync.Mutex

// 生成id
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

// 刷新ID到数据库
func FlushId() {
	var idRules []model.IdRule
	db_helper.Db().Select("type,current_id").Find(&idRules)
	if len(idRules) == 0 {
		return
	}
	var idRule model.IdRule
	for _, id_rule := range idRules {
		key := "IdRuleType:" + id_rule.Type
		if ir, found := cache_helper.GoCache().Get(key); found {
			idRule = ir.(model.IdRule)
		} else { //缓存不存在不处理
			return
		}

		if idRule.CurrentId <= id_rule.CurrentId { //没有增加则不处理
			return
		}

		//异步更新数据库
		if err := db_helper.Db().Where("id=?", idRule.Id).Updates(&model.IdRule{
			CurrentId: idRule.CurrentId,
		}).Error; err != nil {
			log_helper.Error("更新异常：", err)
		}
	}
}
