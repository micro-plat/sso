# 单点登录系统

### 1 mgrweb 
是后台管理系统的前端vue项目

### 2 mgrapi 
后台管理系统的api接口项目

#### 2.1 功能介绍
1.子系统管理
2.子系统菜单管理
3.子系统角色管理
4.用户管理
5.用户子系统角色管理


### 3 apiserver 
对外子系统提供api接口  

#### 3.1 子系统远程登录
`
url /subsys/login
reqBody: 
     username #用户名 
     password #密码
     ident    #子系统标识 
     sign     #签名
     timestamp#时间戳
response:
     成功: 用户信息
     失败: 错误信息
     
` 

#### 3.2 子系统获取用户菜单数据
`
url /subsys/menu
reqBody: 
    system_id:子系统id标识
    user_id:用户标识
    ident:子系统标识
    timestamp:时间戳
    sign:签名
response:
     成功: 用户的菜单信息
     失败: 错误信息
` 

### 3.3 子系统,获取用户信息
`
url /subsys/user/info
reqBody: 
    username:用户名称
    ident:子系统标识
    timestamp:时间戳
    sign:签名
response:
     成功: 用户信息
     失败: 错误信息
`

### 3.4 子系统用户修改密码
`
url /subsys/pwd
reqBody: 
    user_id:用户标识
    password:新密码
    password_old:老密码
    ident:子系统标识
    timestamp:时间戳
    sign:签名
response:
     成功: success
     失败: 错误信息

`

### 3.5 子系统获取系统信息
`
url /subsys/info
reqBody: 
    ident:子系统标识
    timestamp:时间戳
    sign:签名
response:
     成功: 系统信息
     失败: 错误信息
`

### 3.6 子系统通过code来拿取登录的用户标识
`
url /subsys/user/code
reqBody: 
    ident:子系统标识
    timestamp:时间戳
    sign:签名
    code:单点登录还回去的code
response:
     成功: 还回用户id,用户名
     失败: 错误信息
`


### 4 lgapi
sso给子系统提供的单点登录的api


### 5 lgweb
sso给子系统提供的单点登录的web


### 6 说明
#### 6.1 现在sso系统支持两种登录方式
6.1.1 api调用登录(上面有相关的接口)
6.1.2 页面跳转登录
现在提供三个地址可用:
```
1 : 用户登录(登录只用跳转这个页面)
/jump?callback="http%3A%2F%2F192.168.5.79%2F"&sysid=12    
callback是登录成功后要回调的地址，callback地址最好 encodeURIComponent
sysid子系统标识(为了记录登录日志)

2： 刷新sso jwt
/reflesh 刷新sso的登录信息,让sso的过期和子系统保持一致

3: 子系统退出登录
子系统自己先退出，然后调用sso地址
/login?callback="http%3A%2F%2F192.168.5.79%2F  
callback是用户再次登录可跳转地址

```
6.2.2 关于callback地址
可不传，用户登陆后，sso会让用户自己去选择进入那个子系统，但建议还是传过来比较好
```

```
6.2.3 关于回调说明
```
回调地址中会带入code=147336589635，到子系统后应该拿这个code去apiserver取userid(此时子系统就知道是哪个用户登陆了),
后面的调用api和原来保持一致
```


### 7 误区
现在系统分的有点多，对子系统调用的有三个(apiserver纯api, lgapi是页面跳转登陆的后台,以及页面跳转登陆的前台lgweb)
如果不用页面跳转登陆方式就直接用apiserver接口, 否则就用lgweb中的 三个url(上面已提到)


