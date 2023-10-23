package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/exception_helper"
	"go_test/app/helper/page_helper"
	"go_test/app/helper/request_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/model"
)

// 新增ID规则
func SaveIdRule(c *gin.Context) {
	type Param struct {
		Id        int
		Type      string `validate:"required,alphanum" label:"规则标识"`
		CurrentId any
		Step      int32 `validate:"required,gte=1" label:"预占幅度"`
		Prefix    string
		Suffix    string
		MinLength int
	}
	var param Param
	request_helper.InputStruct(c, &param)

	var count int64
	db_helper.Db().Model(&model.IdRule{}).Where("type=?", param.Type).Count(&count)
	var cid int64
	if param.CurrentId != nil {
		cid = int64(param.CurrentId.(float64))
	}
	if count > 0 && param.Id == 0 {
		exception_helper.CommonException("ID标识已存在")
	}
	if count > 0 { //存在
		if err := db_helper.Db().Model(&model.IdRule{}).Where("type=?", param.Type).Updates(model.IdRule{
			CurrentId: cid,
			Step:      param.Step,
			Prefix:    param.Prefix,
			Suffix:    param.Suffix,
			MinLength: param.MinLength,
		}).Error; err != nil {
			exception_helper.CommonException("保存失败")
		}
	} else { //不存在
		idrule := model.IdRule{
			Type:      param.Type,
			CurrentId: cid,
			Step:      param.Step,
			Prefix:    param.Prefix,
			Suffix:    param.Suffix,
			MinLength: param.MinLength,
		}
		if err := db_helper.Db().Create(&idrule).Error; err != nil {
			exception_helper.CommonException("新增失败")
		}
	}

	response_helper.Success(c, "保存成功")
}

// 删除ID规则
func DelIdRule(c *gin.Context) {
	type Param struct {
		Ids []interface{} `validate:"required,min=1"`
	}
	var param Param
	request_helper.InputStruct(c, &param)

	var types []string
	db_helper.Db().Model(&model.IdRule{}).Where("id in ?", param.Ids).Pluck("type", &types)
	if len(types) == 0 {
		exception_helper.CommonException("ID规则不存在")
	}

	if db_helper.Db().Where("id in ?", param.Ids).Delete(&model.IdRule{}).Error != nil {
		exception_helper.CommonException("删除失败")
	}
	response_helper.Success(c, "删除成功")
}

// 获取授权码列表
func GetIdRuleList(c *gin.Context) {
	type Param struct {
		Type string
	}
	var param Param
	request_helper.InputStruct(c, &param)
	db := db_helper.Db()
	if param.Type != "" {
		db = db.Where("type like ?", "%"+param.Type+"%")
	}
	res := page_helper.AutoPage(c, db.Order("id desc").Model(model.IdRule{}))

	response_helper.Success(c, "获取成功", res)
}
