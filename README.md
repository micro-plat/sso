# 单点登录系统

### 1 大概结构
```
1 包含两个项目(loginserver(包含api接口和web统一登录页面), mgrserver(用户权限管理系统))
2 两个sdk(api调用包装、前端调用包装)
```

### 项目说明

##### 1 loginserver 跳转登录
###### 1.1 api服务 (提供子系统调用的用户和系统操作接口)
###### 1.2 web服务 (跳转登录web)

##### 2 mgrserver 用户权限管理系统
###### 2.1 mgrweb (用户权限管理系统)

### 使用方法
1 前端使用请看sdk/sso-js
2 后端使用请看sdk/sso

##### 发布使用中可能遇到的问题
1 jwt服务端不能接受到
###### 1.1 nginx 配置问题,拦截了__sso_jwt__
###### 1.2 跨域问题(zookeeper-> header中设置(正常如下)))
``` json
{
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
    "Access-Control-Allow-Headers": "X-Requested-With,Content-Type,__sso_jwt__,X-Requested-Id",
    "Access-Control-Allow-Credentials": "true",
    "Access-Control-Expose-Headers":"__sso_jwt__"
}
```

2 登录回调失败(查看注册中心app配置是否正常),大概如下(必须包含如下三个字段)
``` json
{	
    "secret":"311124b57e468ff88e4f1c8743354314", //这个必须和数据库中的一样
    "sso_api_host":"http://api.sso.18jiayou1.com:6689", //loginserver的api接口地址
    "ident":"17sup" //创建系统时取的英文名
}
```

3 登录成功后不能回跳到子系统中
    系统中回调地址没有设置，在用户权限管理系统中加上地址：如: http://web.xx.100bm.cn/ssocallback
```

