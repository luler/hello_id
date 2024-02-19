# hello_id

#### 介绍

高性能生成snowflake id、sonyflake id、uuid v1、uuid v4、xid、ksuid以及自定义ID的服务

#### 安装

```
docker-compose up -d
```

#### 使用

##### 生成雪花id（基于时间、机器id、随机数生成，有序，整数类型）

接口: /api/snowflake

方法：GET

请求参数：

| 字段       | 是否必填 | 类型     | 描述                  |
|----------|------|--------|---------------------|
| workerId | 否    | int    | 机器ID，取值范围0-1023，默认0 |
| length   | 否    | int    | 	获取ID数量，默认1，最大500   |
| authKey  | 是    | string | 	授权码                |

返回示例：

```json
{
  "code": 200,
  "data": {
    "ids": [
      "1711927875550052352"
    ],
    "type": "snowflake",
    "workerId": 0
  },
  "message": "创建成功"
}
```

##### 生成sonyflake id（snowflake的改进版本，可用时间更长，有序，整数类型）

接口: /api/sonyflake

方法：GET

请求参数：

| 字段      | 是否必填 | 类型     | 描述                |
|---------|------|--------|-------------------|
| length  | 否    | int    | 	获取ID数量，默认1，最大500 |
| authKey | 是    | string | 	授权码              |

返回示例：

```json
{
  "code": 200,
  "data": {
    "ids": [
      "482279196769390593"
    ],
    "type": "sonyflake"
  },
  "message": "创建成功"
}
```

##### 生成uuid V1 （根据时间、Mac地址等信息生成，有时间顺序）

接口: /api/uuid1

方法：GET

请求参数：

| 字段      | 是否必填 | 类型     | 描述                |
|---------|------|--------|-------------------|
| length  | 否    | int    | 	获取ID数量，默认1，最大500 |
| authKey | 是    | string | 	授权码              |

返回示例：

```json
{
  "code": 200,
  "data": {
    "ids": [
      "c9973b96-67db-11ee-89f2-00ff0c3e16f1"
    ],
    "type": "uuid1"
  },
  "message": "创建成功"
}
```

##### 生成uuid V4 （完全随机生成，无序）

接口: /api/uuid4

方法：GET

请求参数：

| 字段      | 是否必填 | 类型     | 描述                |
|---------|------|--------|-------------------|
| length  | 否    | int    | 	获取ID数量，默认1，最大500 |
| authKey | 是    | string | 	授权码              |

返回示例：

```json
{
  "code": 200,
  "data": {
    "ids": [
      "f26cb76e-574b-4797-949e-85f449c77ee3"
    ],
    "type": "uuid4"
  },
  "message": "创建成功"
}
```

##### 生成xid（基于时间戳、随机数生成，有序）

接口: /api/xid

方法：GET

请求参数：

| 字段      | 是否必填 | 类型     | 描述                |
|---------|------|--------|-------------------|
| length  | 否    | int    | 	获取ID数量，默认1，最大500 |
| authKey | 是    | string | 	授权码              |

返回示例：

```json
{
  "code": 200,
  "data": {
    "ids": [
      "ckj09svfe3qiv1150kr0"
    ],
    "type": "xid"
  },
  "message": "创建成功"
}
```

##### 生成ksuid（基于时间戳、随机数生成，随机数比xid更大，唯一性更好，有序）

接口: /api/ksuid

方法：GET

请求参数：

| 字段      | 是否必填 | 类型     | 描述                |
|---------|------|--------|-------------------|
| length  | 否    | int    | 	获取ID数量，默认1，最大500 |
| authKey | 是    | string | 	授权码              |

返回示例：

```json
{
  "code": 200,
  "data": {
    "ids": [
      "2Wb9ZXX2OWPOpubXAKRLwQGn7cP"
    ],
    "type": "ksuid"
  },
  "message": "创建成功"
}
```

##### 生成自定义ID（简单定制化，原子有序）

接口: /api/getId

方法：GET

请求参数：

| 字段      | 是否必填 | 类型     | 描述                |
|---------|------|--------|-------------------|
| type    | 是    | string | ID标识              |
| length  | 否    | int    | 	获取ID数量，默认1，最大500 |
| authKey | 是    | string | 	授权码              |

返回示例：

```json
{
  "code": 200,
  "data": [
    "20230930000000114.go"
  ],
  "message": "获取成功"
}
```