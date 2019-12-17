# 单点登录系统

### 1 大概结构
```
1 包含五个项目(apiserver, loginserver(lgapi,lgweb), mgrserver(mgrapi, mgrweb))
2 两个sdk(api调用包装、前端调用包装)
```

### 项目说明

##### 1 apiserver 
提供给子系统调用的api接口服务

##### 2 loginserver 跳转登录
###### 2.1 lgapi (跳转登录api)
###### 2.2 lgweb (跳转登录web)

##### 3 mgrserver 用户系统
###### 2.1 mgrapi (用户系统api)
###### 2.2 mgrweb (用户系统web)

### 使用方法
1 前端使用请看sdk/sso-js
2 后端使用请看sdk/sso

##### 发布使用中可能遇到的问题
1 jwt服务端不能接受到
###### 1.1 nginx 配置问题,拦截了__jwt__
###### 1.2 跨域问题(zookeeper-> header中设置(正常如下)))
``` json
{
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,PATCH,OPTIONS",
    "Access-Control-Allow-Headers": "X-Requested-With,Content-Type,__jwt__,X-Requested-Id",
    "Access-Control-Allow-Credentials": "true",
    "Access-Control-Expose-Headers":"__jwt__"
}
```
###### 1.3 登录验证接口没有加入例外 auth节点设置(正常如下)
``` json 
{
"jwt": {
    "name": "__jwt__",
    "expireAt": 36000,
    "mode": "HS512",
    "source": "HEADER",
    "secret": "396e03f316218ec948808cbc9cb5539d",
    "exclude": [
      "/sso/login/verify"
    ]
  }
}
```

2 登录回调失败(查看app配置是否正常),大概如下(必须包含如下三个字段)
``` json
{	
    "secret":"311124b57e468ff88e4f1c8743354314", //这个必须和数据库中的一样
    "sso_api_host":"http://api.sso.18jiayou1.com:6689", //apiserver线下地址
    "ident":"17sup" //创建系统时取的英文名
}
```

3 登录成功后不能会跳到子系统中
```
    系统中回调地址没有设置，在系统管理中加上地址：如: http://web.xx.100bm.cn/ssocallback
```
