package request_helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"go_test/app/helper/helper"
	"go_test/app/helper/valid_helper"
	"net/url"
	"reflect"
	"strings"
)

// 获取请求参数-返回map类型
func Input(c *gin.Context, fields ...string) map[string]interface{} {
	param := make(map[string]interface{})
	// 获取所有的 GET 参数
	queryParams := c.Request.URL.Query()
	for key, value := range extractParam(c, queryParams, "get") {
		param[key] = value
	}
	// 获取所有的 POST 参数
	c.Request.ParseForm()
	queryParams = c.Request.PostForm
	for key, value := range extractParam(c, queryParams, "post") {
		param[key] = value
	}
	// 获取multi-form类型所有的 POST 参数
	err := c.Request.ParseMultipartForm(32 << 24)
	if err == nil {
		queryParams = c.Request.MultipartForm.Value
		for key, value := range extractParam(c, queryParams, "post") {
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
	//参数过滤
	param = helper.FilterMap(param, fields)
	return param
}

// 获取GET请求参数
func ParamGet(c *gin.Context, fields ...string) map[string]interface{} {
	param := make(map[string]interface{})
	// 获取所有的 GET 参数
	queryParams := c.Request.URL.Query()
	for key, value := range extractParam(c, queryParams, "get") {
		param[key] = value
	}
	//参数过滤
	param = helper.FilterMap(param, fields)
	return param
}

// 获取PostForm请求参数
func ParamPostForm(c *gin.Context, fields ...string) map[string]interface{} {
	param := make(map[string]interface{})
	// 获取所有的PostForm参数
	c.Request.ParseForm()
	queryParams := c.Request.PostForm
	for key, value := range extractParam(c, queryParams, "post") {
		param[key] = value
	}

	return param
}

// 获取MultipartForm请求参数
func ParamMultipartForm(c *gin.Context, fields ...string) map[string]interface{} {
	param := make(map[string]interface{})
	// 获取multi-form类型所有的 POST 参数
	err := c.Request.ParseMultipartForm(32 << 24)
	if err == nil {
		queryParams := c.Request.MultipartForm.Value
		for key, value := range extractParam(c, queryParams, "post") {
			param[key] = value
		}
	}
	//参数过滤
	param = helper.FilterMap(param, fields)
	return param
}

// 获取json格式的请求参数
func ParamRawJson(c *gin.Context, fields ...string) map[string]interface{} {
	param := make(map[string]interface{})
	//获取json类型参数
	raw, _ := c.GetRawData()
	jsonData := make(map[string]interface{})
	json.Unmarshal(raw, &jsonData)
	for key, value := range jsonData {
		param[key] = value
	}

	//参数过滤
	param = helper.FilterMap(param, fields)
	return param
}

// 获取参数并验证
func InputStruct(c *gin.Context, param interface{}) {
	data := Input(c)
	mapstructure.Decode(data, param)
	valid_helper.Check(param)
}

// 解析请求参数
func extractParam(c *gin.Context, queryParams url.Values, t string) map[string]interface{} {
	param := make(map[string]interface{})
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
