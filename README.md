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
