# hello_id

#### 介绍

高性能生成snowflake id、sonyflake id、uuid v1、uuid v4服务

#### 安装

```
docker-compose up -d
```

#### 使用

##### 生成雪花id（基于时间、机器id、随机数生成，有序，整数类型）

接口: http://ip:8080/snowflake

方法：GET、POST等

请求参数：

| 字段        | 是否必填 | 类型  | 描述                  |
|-----------|------|-----|---------------------|
| worker_id | 否    | int | 机器ID，取值范围0-1023，默认0 |

返回示例：

```json
{
  "code": 200,
  "data": {
    "id": "1697622854507429888",
    "type": "snowflake",
    "worker_id": 0
  },
  "message": "创建成功"
}
```

##### 生成sonyflake id（snowflake的改进版本，可用时间更长，有序，整数类型）

接口: http://ip:8080/sonyflake

方法：GET、POST等

请求参数：

| 字段 | 是否必填 | 类型 | 描述 |
|----|------|----|----|
|    |      |    |    |

返回示例：

```json
{
  "code": 200,
  "data": {
    "id": "476556854336421890",
    "type": "sonyflake"
  },
  "message": "创建成功"
}
```

##### 生成uuid V1 （根据时间、Mac地址等信息生成，有时间顺序）

接口: http://ip:8080/uuid1

方法：GET、POST等

请求参数：

| 字段 | 是否必填 | 类型 | 描述 |
|----|------|----|----|
|    |      |    |    |

返回示例：

```json
{
  "code": 200,
  "data": {
    "id": "ee13483c-48d6-11ee-ae17-00ff7507da7c",
    "type": "uuid1"
  },
  "message": "创建成功"
}
```

##### 生成uuid V4 （完全随机生成，无序）

接口: http://ip:8080/uuid4

方法：GET、POST等

请求参数：

| 字段 | 是否必填 | 类型 | 描述 |
|----|------|----|----|
|    |      |    |    |

返回示例：

```json
{
  "code": 200,
  "data": {
    "id": "ee13483c-48d6-11ee-ae17-00ff7507da7c",
    "type": "uuid1"
  },
  "message": "创建成功"
}
```

##### 生成xid（基于时间戳、随机数生成，有序）

接口: http://ip:8080/xid

方法：GET、POST等

请求参数：

| 字段 | 是否必填 | 类型 | 描述 |
|----|------|----|----|
|    |      |    |    |

返回示例：

```json
{
  "code": 200,
  "data": {
    "id": "cjpc168ip8e48m05a3l0",
    "type": "xid"
  },
  "message": "创建成功"
}
```

##### 生成ksuid（基于时间戳、随机数生成，随机数比xid更大，唯一性更好，有序）

接口: http://ip:8080/ksuid

方法：GET、POST等

请求参数：

| 字段 | 是否必填 | 类型 | 描述 |
|----|------|----|----|
|    |      |    |    |

返回示例：

```json
{
  "code": 200,
  "data": {
    "id": "2UpIxKEKXsXzb52yVxeEgif9WDJ",
    "type": "ksuid"
  },
  "message": "创建成功"
}
```