package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"go_test/app/helper/request_helper"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// 格式化日期时间
func LocalTimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 自动分页获取
func AutoPage(c *gin.Context, db *gorm.DB) map[string]interface{} {
	type Param struct {
		Page     any
		PageSize any
	}
	var param Param
	request_helper.InputStruct(c, &param)

	switch param.Page.(type) {
	case string:
		param.Page, _ = strconv.Atoi(param.Page.(string))
	case int, int8, int16, int32, int64:
		param.Page = param.Page
	default:
		param.Page = 1
	}
	switch param.PageSize.(type) {
	case string:
		param.PageSize, _ = strconv.Atoi(param.PageSize.(string))
	case int, int8, int16, int32, int64:
		param.PageSize = param.PageSize
	default:
		param.PageSize = 10
	}

	var total int64
	db.Count(&total)
	var data []map[string]interface{}
	if total > 0 {
		offset := (param.Page.(int) - 1) * param.PageSize.(int)
		db.Limit(param.PageSize.(int)).Offset(offset).Find(&data)
	}
	if data == nil {
		data = []map[string]interface{}{}
	}
	//字段处理
	for key, item := range data {
		// 创建新的 map 存储转换后的字段
		converted := make(map[string]interface{})
		for key, value := range item {
			// 将字段名转换为大驼峰
			convertedKey := strcase.ToCamel(key)
			converted[convertedKey] = value
		}
		data[key] = converted
	}
	res := make(map[string]interface{})
	res["List"] = data
	res["Page"] = param.Page
	res["PageSize"] = param.PageSize
	res["Total"] = total

	return res
}
