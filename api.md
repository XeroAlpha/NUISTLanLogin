# NUIST LAN API

1. 获取 IP 地址
2. 尝试登录 / 获取可用渠道
3. 选择渠道后再次进行登录

## 获取 IP

```json
GET /api/v1/ip
-----
Content-Type: application/json
{
    "code": 200,
    "data": "<IP地址>"
}
```

## 尝试登录 / 获取可用渠道

```json
POST /api/v1/login
Access-Control-Allow-Origin: *
Content-Type: application/json;charset=UTF-8
{
    "username": "<用户名>",
    "password": "<密码>",
    "ifautologin": "<允许自动登录>", // true = 1, false = 0
    "channel": "_GET",
    "pagesign": "firstauth",
    "usripadd": "<IP地址>"
}
-----
// 需要选择渠道
Content-Type: application/json;charset=gbk
{
    "code": 200,
    "message": "ok",
    "data": {
        "channels": [
            {
                "id": "<渠道ID>",
                "name": "<渠道名>"
            },
            // ...
        ]
    }
}
-----
// 直接登录
Content-Type: application/json;charset=gbk
{
    "code": 200,
    "message": "ok",
    "data": {
        "reauth": false, 
        "username": "<用户名>",
        "balance": "<余额/元>",
        "duration": "<在线时长/秒>",				
        "outport": "<渠道名>",
        "totaltimespan": "<累计使用时长/秒>",
        "usripadd": "<IP地址>"
    }
}
-----
// 登录失败
Content-Type: application/json;charset=gbk
{
    "code": 201, // code != 200
    "message": "ok",
    "data": {
        "text": "<错误信息>",
        "url": null // 有时可能是一个 URL，此时需要引导跳转到对应的 URL
    }
}
```

## 再次登录

```json
POST /api/v1/login
Access-Control-Allow-Origin: *
Content-Type: application/json;charset=UTF-8
{
    "username": "<用户名>",
    "password": "<密码>",
    "ifautologin": "<允许自动登录>", // true = 1, false = 0
    "channel": "<渠道ID>", // 渠道ID 不可为 "0"
    "pagesign": "secondauth",
    "usripadd": "<IP地址>"
}
-----
// 响应参见【直接登录】与【登录失败】
```

## 请求在线统计信息
```json
POST /api/v1/login
Access-Control-Allow-Origin: *
Content-Type: application/json;charset=UTF-8
{
    "username": "<用户名>",
    "password": "<密码>",
    "ifautologin": "1",
    "channel": "_ONELINEINFO",
    "pagesign": "thirdauth",
    "usripadd": "<IP地址>"
}
-----
// 响应参见【直接登录】
```

## 注销
```json
POST /api/v1/login
Access-Control-Allow-Origin: *
Content-Type: application/json;charset=UTF-8
{
    "username": "<用户名>",
    "password": "<密码>",
    "ifautologin": "1",
    "channel": "0",
    "pagesign": "thirdauth",
    "usripadd": "<IP地址>"
}
-----
Content-Type: application/json;charset=gbk
{
    "code": 200,
    "message": "ok",
    "data": {
        "text": "<信息>",
        "url": null
    }
}
```

## 