package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/model"
	"time"
)

// 新增ID规则
func SaveIdRule(c *gin.Context) {
	type Param struct {
		Type      string `validate:"required,alphanum" label:"规则标识"`
		CurrentId any
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
			CurrentId: int(param.CurrentId.(float64)),
		}).Error; err != nil {
			helper.CommonException("保存失败")
		}
	} else { //不存在
		idrule := model.IdRule{
			Type: param.Type,
		}
		if param.CurrentId != nil {
			idrule.CurrentId = param.CurrentId.(int)
		}
		if err := db_helper.Db().Create(&idrule).Error; err != nil {
			helper.CommonException("新增失败")
		}
	}

	response_helper.Success(c, "保存成功")
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
