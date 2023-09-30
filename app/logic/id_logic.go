package logic

import (
	"go_test/app/helper/cache_helper"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/exception_helper"
	"go_test/app/helper/log_helper"
	"go_test/app/model"
	"strconv"
	"strings"
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
			exception_helper.CommonException("规则标识不存在")
		}
		//缓存24小时
		cache_helper.GoCache().Set(key, idRule, time.Hour*24)
	} else {
		idRule = ir.(model.IdRule)
	}
	ids := []string{}
	for i := 0; i < length; i++ {
		idRule.CurrentId += 1
		id := strconv.FormatInt(idRule.CurrentId, 10)
		tn := time.Now()
		//获取前缀
		prefix := ""
		if idRule.Prefix != "" {
			prefixArr := strings.Split(idRule.Prefix, ",")
			for _, pre := range prefixArr {
				switch pre {
				case "年":
					prefix = prefix + tn.Format("2006")
				case "月":
					prefix = prefix + tn.Format("01")
				case "日":
					prefix = prefix + tn.Format("02")
				case "时":
					prefix = prefix + tn.Format("15")
				case "分":
					prefix = prefix + tn.Format("04")
				case "秒":
					prefix = prefix + tn.Format("05")
				default:
					prefix = prefix + pre
				}
			}
		}
		//获取后缀
		suffix := ""
		if idRule.Suffix != "" {
			suffixArr := strings.Split(idRule.Suffix, ",")
			for _, su := range suffixArr {
				switch su {
				case "年":
					suffix = suffix + tn.Format("2006")
				case "月":
					suffix = suffix + tn.Format("01")
				case "日":
					suffix = suffix + tn.Format("02")
				case "时":
					suffix = suffix + tn.Format("15")
				case "分":
					suffix = suffix + tn.Format("04")
				case "秒":
					suffix = suffix + tn.Format("05")
				default:
					suffix = suffix + su
				}
			}
		}
		//限制固定长度
		if idRule.FixedLength > 0 {
			lll := len(prefix) + len(suffix) + len(id)
			if lll > idRule.FixedLength { //已达最大长度限制
				exception_helper.CommonException("已达最大长度限制")
			} else if lll < idRule.FixedLength { //小于长度需要补0
				id = strings.Repeat("0", idRule.FixedLength-lll) + id
			}
		}
		id = prefix + id + suffix
		ids = append(ids, id)
	}
	//缓存24小时
	cache_helper.GoCache().Set(key, idRule, time.Hour*24)
	return ids
}

// 刷新ID到数据库
func FlushId() {
	var idRule model.IdRule
	items := cache_helper.GoCache().Items()
	for k, v := range items {
		if strings.HasPrefix(k, "IdRuleType:") {
			idRule = v.Object.(model.IdRule)
			keyId := "keyId:" + idRule.Type
			if cacheCurrentId, exist := cache_helper.GoCache().Get(keyId); exist {
				if cacheCurrentId == idRule.CurrentId {
					continue
				}
			}
			//异步更新数据库
			if err := db_helper.Db().Where("id=?", idRule.Id).Updates(&model.IdRule{
				CurrentId: idRule.CurrentId,
			}).Error; err != nil {
				log_helper.Error("更新异常：", err)
			}

			cache_helper.GoCache().Set(keyId, idRule.CurrentId, time.Hour*24)
		}
	}
}
