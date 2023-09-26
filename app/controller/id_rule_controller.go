package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper"
	"go_test/app/helper/cache_helper"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/model"
	"time"
)

// 新增ID规则
func SaveIdRule(c *gin.Context) {
	type Param struct {
		Type        string `validate:"required,alphanum" label:"规则标识"`
		CurrentId   any
		Prefix      string
		Suffix      string
		FixedLength int16
	}
	var param Param
	helper.InputStruct(c, &param)

	var count int64
	db_helper.Db().Model(&model.IdRule{}).Where("type=?", param.Type).Count(&count)
	if count > 0 { //存在
		if param.CurrentId == nil {
			helper.CommonException("请设置当前ID")
		}
		if err := db_helper.Db().Model(&model.IdRule{}).Where("type=?", param.Type).Updates(model.IdRule{
			CurrentId:   int(param.CurrentId.(float64)),
			Prefix:      param.Prefix,
			Suffix:      param.Suffix,
			FixedLength: param.FixedLength,
		}).Error; err != nil {
			helper.CommonException("保存失败")
		}
	} else { //不存在
		idrule := model.IdRule{
			Type:        param.Type,
			CurrentId:   int(param.CurrentId.(float64)),
			Prefix:      param.Prefix,
			Suffix:      param.Suffix,
			FixedLength: param.FixedLength,
		}
		if err := db_helper.Db().Create(&idrule).Error; err != nil {
			helper.CommonException("新增失败")
		}
	}

	//删除缓存
	key := "IdRuleType:" + param.Type
	cache_helper.GoCache().Delete(key)

	response_helper.Success(c, "保存成功")
}

// 删除ID规则
func DelIdRule(c *gin.Context) {
	type Param struct {
		Ids []interface{} `validate:"required,min=1"`
	}
	var param Param
	helper.InputStruct(c, &param)

	var types []string
	db_helper.Db().Model(&model.IdRule{}).Where("id in ?", param.Ids).Pluck("type", &types)
	if len(types) == 0 {
		helper.CommonException("ID规则不存在")
	}

	if db_helper.Db().Where("id in ?", param.Ids).Delete(&model.IdRule{}).Error != nil {
		helper.CommonException("删除失败")
	}
	//删除缓存
	for _, t := range types {
		key := "IdRuleType:" + t
		cache_helper.GoCache().Delete(key)
	}
	response_helper.Success(c, "删除成功")
}

// 获取授权码列表
func GetIdRuleList(c *gin.Context) {
	res := helper.AutoPage(c, db_helper.Db().Model(model.IdRule{}))

	for _, a := range res["List"].([]map[string]interface{}) {
		a["CreatedAt"] = helper.LocalTimeFormat(a["CreatedAt"].(time.Time))
		a["UpdatedAt"] = helper.LocalTimeFormat(a["UpdatedAt"].(time.Time))
	}
	response_helper.Success(c, "获取成功", res)
}
