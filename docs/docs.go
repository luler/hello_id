// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "开发人员",
            "url": "https://cas.luler.top/",
            "email": "1207032539@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/getId": {
            "get": {
                "description": "生成自定义ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ID生成相关接口"
                ],
                "summary": "生成自定义ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID标识",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "获取ID数量，默认1，最大500",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "授权码",
                        "name": "authKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/ksuid": {
            "get": {
                "description": "生成Ksuid类型的id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ID生成相关接口"
                ],
                "summary": "生成Ksuid类型的id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "获取ID数量，默认1，最大500",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "授权码",
                        "name": "authKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/snowflake": {
            "get": {
                "description": "创建雪花ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ID生成相关接口"
                ],
                "summary": "创建雪花ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "机器ID，取值范围0-1023，默认0",
                        "name": "workerId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "获取ID数量，默认1，最大500",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "授权码",
                        "name": "authKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/sonyflake": {
            "get": {
                "description": "\"创建sonyflake ID\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ID生成相关接口"
                ],
                "summary": "\"创建sonyflake ID\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "获取ID数量，默认1，最大500",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "授权码",
                        "name": "authKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/uuid1": {
            "get": {
                "description": "生成uuid版本1类型的id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ID生成相关接口"
                ],
                "summary": "生成uuid版本1类型的id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "获取ID数量，默认1，最大500",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "授权码",
                        "name": "authKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/uuid4": {
            "get": {
                "description": "生成uuid版本4类型的id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ID生成相关接口"
                ],
                "summary": "生成uuid版本4类型的id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "获取ID数量，默认1，最大500",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "授权码",
                        "name": "authKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/xid": {
            "get": {
                "description": "生成Xid类型的id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ID生成相关接口"
                ],
                "summary": "生成Xid类型的id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "获取ID数量，默认1，最大500",
                        "name": "length",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "授权码",
                        "name": "authKey",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "接口文档",
	Description:      "当前页面用于展示项目一些开放的接口",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
