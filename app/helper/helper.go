package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"net/url"
	"reflect"
	"strings"
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
