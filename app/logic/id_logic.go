package logic

import (
	"go_test/app/helper/db_helper"
	"go_test/app/helper/exception_helper"
	"go_test/app/model"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
	"sync"
	"time"
)

var lock sync.Mutex

var idRules = make(map[string]model.IdRule)

// 生成id
func GenerateId(idRuleType string, length int) []string {
	lock.Lock()
	defer lock.Unlock()

	var idRule model.IdRule
	var exists bool
	ids := []string{}
	for i := 0; i < length; i++ {

		idRule, exists = idRules[idRuleType]

		if !exists || idRule.Step <= 0 {
			tx := db_helper.Db().Begin()
			if tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("type=?", idRuleType).First(&idRule).Error != nil {
				exception_helper.CommonException("规则标识不存在")
			}
			if idRule.Step <= 0 {
				exception_helper.CommonException("预占幅度不能小于等于0")
			}

			idRuleCopy := idRule

			idRuleCopy.CurrentId = idRuleCopy.CurrentId + int64(idRuleCopy.Step)
			tx.Save(&idRuleCopy)

			tx.Commit()

			idRules[idRuleType] = idRule
		}

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
		//限制最小长度
		if idRule.MinLength > 0 {
			lll := len(id)
			if lll < idRule.MinLength { //已达最大长度限制
				id = strings.Repeat("0", idRule.MinLength-lll) + id
			}
		}
		id = prefix + id + suffix
		ids = append(ids, id)
		idRule.Step--
		idRules[idRuleType] = idRule
	}
	return ids
}
