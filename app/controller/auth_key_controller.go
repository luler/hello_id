package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go_test/app/helper"
	"go_test/app/helper/cache_helper"
	"go_test/app/helper/db_helper"
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
	helper.InputStruct(c, &param)

	auth_key := model.AuthKey{}
	auth_key.Id = uint(param.Id)
	auth_key.Remark = param.Remark
	if param.Id > 0 {
		if err := db_helper.Db().Model(auth_key).Updates(map[string]interface{}{
			"Remark": auth_key.Remark,
		}).Error; err != nil {
			helper.CommonException("保存失败")
		}
	} else {
		uuid4 := uuid.NewV4().String()
		auth_key.AuthKey = uuid4
		if err := db_helper.Db().Create(&auth_key).Error; err != nil {
			fmt.Println(err)
			helper.CommonException("新增失败")
		}
	}
	response_helper.Success(c, "保存成功")
}

// 获取授权码列表
func GetAuthKeyList(c *gin.Context) {
	res := helper.AutoPage(c, db_helper.Db().Model(model.AuthKey{}))

	for _, a := range res["List"].([]map[string]interface{}) {
		a["CreatedAt"] = helper.LocalTimeFormat(a["CreatedAt"].(time.Time))
		a["UpdatedAt"] = helper.LocalTimeFormat(a["UpdatedAt"].(time.Time))
	}
	response_helper.Success(c, "获取成功", res)
}

// 获取授权码列表
func DelAuthKey(c *gin.Context) {
	type Param struct {
		Ids []interface{} `validate:"required,min=1"`
	}
	var param Param
	helper.InputStruct(c, &param)

	var authKeys []string
	db_helper.Db().Model(&model.AuthKey{}).Where("id in ?", param.Ids).Pluck("auth_key", &authKeys)
	if len(authKeys) == 0 {
		helper.CommonException("授权码不存在")
	}

	if db_helper.Db().Where("id in ?", param.Ids).Delete(&model.AuthKey{}).Error != nil {
		helper.CommonException("删除失败")
	}
	//删除缓存
	for _, authKey := range authKeys {
		key := "AuthKey:" + authKey
		cache_helper.GoCache().Delete(key)
	}
	response_helper.Success(c, "删除成功")
}
