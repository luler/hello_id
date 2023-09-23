package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 获取请求参数-返回map类型
func Input(c *gin.Context, fields ...string) map[string]interface{} {
	param := make(map[string]interface{})
	extractParam := func(queryParams url.Values, t string) map[string]interface{} {
		for key, values := range queryParams {
			if strings.HasSuffix(key, "]") {
				//存在类似param[1]的参数，需要特殊解析下
				key = key[0:strings.Index(key, "[")]
				switch t {
				case "get":
					qm := c.QueryMap(key)
					param[key] = qm
				case "post":
					qm := c.PostFormMap(key)
					param[key] = qm
				}

			} else {
				for _, value := range values {
					//判断是否存在多个相同参数，是就组成数组
					if _, ok := param[key]; ok {
						if reflect.TypeOf(param[key]).Kind() == reflect.Array {
							param[key] = append(param[key].([]interface{}), value)
						} else {
							param[key] = []interface{}{
								param[key],
								value,
							}
						}
					} else {
						param[key] = value
					}
				}
			}
		}

		return param
	}
	// 获取所有的 GET 参数
	queryParams := c.Request.URL.Query()
	for key, value := range extractParam(queryParams, "get") {
		param[key] = value
	}
	// 获取所有的 POST 参数
	c.Request.ParseForm()
	queryParams = c.Request.PostForm
	for key, value := range extractParam(queryParams, "post") {
		param[key] = value
	}
	// 获取multi-form类型所有的 POST 参数
	err := c.Request.ParseMultipartForm(32 << 24)
	if err == nil {
		queryParams = c.Request.MultipartForm.Value
		for key, value := range extractParam(queryParams, "post") {
			param[key] = value
		}
	}
	//获取json类型参数
	raw, _ := c.GetRawData()
	jsonData := make(map[string]interface{})
	json.Unmarshal(raw, &jsonData)
	for key, value := range jsonData {
		param[key] = value
	}
	//过滤参数
	filteredParam := make(map[string]interface{})
	if len(fields) > 0 {
		for _, field := range fields {
			if _, ok := param[field]; ok {
				filteredParam[field] = param[field]
			}
		}
	} else {
		filteredParam = param
	}

	return filteredParam
}

// 获取参数并验证
func InputStruct(c *gin.Context, param interface{}) {
	data := Input(c)
	mapstructure.Decode(data, param)
	Check(param)
}

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
	InputStruct(c, &param)

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
