package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go_test/app/helper/cache_helper"
	"go_test/app/helper/db_helper"
	"go_test/app/helper/exception_helper"
	"go_test/app/helper/helper"
	"go_test/app/helper/page_helper"
	"go_test/app/helper/request_helper"
	"go_test/app/helper/response_helper"
	"go_test/app/model"
	"time"
)

// 新增授权码
func SaveAuthKey(c *gin.Context) {
	type Param struct {
		Id     int    `validate:"number"`
		Remark string `validate:"required,max=255" label:"用途描述"`
	}
	var param Param
	request_helper.InputStruct(c, &param)

	auth_key := model.AuthKey{}
	auth_key.Id = uint(param.Id)
	auth_key.Remark = param.Remark
	if param.Id > 0 {
		if err := db_helper.Db().Model(auth_key).Updates(map[string]interface{}{
			"Remark": auth_key.Remark,
		}).Error; err != nil {
			exception_helper.CommonException("保存失败")
		}
	} else {
		uuid4 := uuid.NewV4().String()
		auth_key.AuthKey = uuid4
		if err := db_helper.Db().Create(&auth_key).Error; err != nil {
			fmt.Println(err)
			exception_helper.CommonException("新增失败")
		}
	}
	response_helper.Success(c, "保存成功")
}

// 获取授权码列表
func GetAuthKeyList(c *gin.Context) {
	type Param struct {
		AuthKey string
		Remark  string
	}
	var param Param
	request_helper.InputStruct(c, &param)
	db := db_helper.Db()
	if param.AuthKey != "" {
		db = db.Where("auth_key like ?", "%"+param.AuthKey+"%")
	}
	if param.Remark != "" {
		db = db.Where("remark like ?", "%"+param.Remark+"%")
	}
	res := page_helper.AutoPage(c, db.Order("id desc").Model(model.AuthKey{}))

	for _, a := range res["list"].([]map[string]interface{}) {
		a["createdAt"] = helper.LocalTimeFormat(a["createdAt"].(time.Time))
		a["updatedAt"] = helper.LocalTimeFormat(a["updatedAt"].(time.Time))
	}
	response_helper.Success(c, "获取成功", res)
}

// 获取授权码列表
func DelAuthKey(c *gin.Context) {
	type Param struct {
		Ids []interface{} `validate:"required,min=1"`
	}
	var param Param
	request_helper.InputStruct(c, &param)

	var authKeys []string
	db_helper.Db().Model(&model.AuthKey{}).Where("id in ?", param.Ids).Pluck("auth_key", &authKeys)
	if len(authKeys) == 0 {
		exception_helper.CommonException("授权码不存在")
	}

	if db_helper.Db().Where("id in ?", param.Ids).Delete(&model.AuthKey{}).Error != nil {
		exception_helper.CommonException("删除失败")
	}
	//删除缓存
	for _, authKey := range authKeys {
		key := "AuthKey:" + authKey
		cache_helper.GoCache().Delete(key)
	}
	response_helper.Success(c, "删除成功")
}
