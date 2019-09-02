## sso sdk使用说明

在跳转登录返回后,子系统需要验证登录用户、获取菜单、获取系统信息等，为了降低使用接口的复杂度,将这些接口调用包装成sdk(请看完所有说明)

现在支持两种方案对接
1: 直接使用bind，sdk做api的调用以及服务的生成(写代码少，缺少灵活)
2: 自己管理ssoclient,以及相应sdk(api)的调用


### 方案一
``` go
在init中 增加
import "github.com/micro-plat/sso/sdk/sso"

if err := sso.Bind(r.MicroApp,conf.SsoApiHost, conf.Ident, conf.Secret); err != nil {
	return err
}

sdk生成了四个接口(前端对接接口)
"/sso/login/verify": 登录回调后验证
"/sso/member/menus/get"：获取用户菜单
"/sso/member/systems/get":获取登录用户有权限的其他系统信息
"/sso/system/info/get": 获取当前系统信息


handling.go中修改
将原来的jwt验证，和登录用户状态的保存都放到了sdk中, 现在只需要加上
import "github.com/micro-plat/sso/sdk/sso"

if err := sso.CheckAndSetMember(ctx); err != nil {
    return err
}

获取登录信息统一用sso sdk这边的方法(原来的就不用了)
import "github.com/micro-plat/sso/sdk/sso"

member := sso.GetMember(ctx)
```


### 方案二

####1 子系统服务端修改点
###### 1.1 在init()⽅法中注⼊sso client


``` go
import "github.com/micro-plat/sso/sdk/sso"
ssoCleint, err := sso.New(conf.SsoApiHost, conf.Ident, conf.Secret)
if err != nil {
return err
}
model.SaveSSOClient(c, ssoCleint) //将sso client 保存起来
```
参数|类型|说明
--|:--:|--:
SsoApiHost |string| 线下:http://api.sso.18jiayou1.com:6689, 线上：http://api.sso.18jiayou.com
ident|string|子系统ident
secret|string|子子系统秘钥, 在【用户系统】-> 【系统管理】中能找到

###### 1.2 保存和获取 sso client 实例
``` go
//SaveSSOClient 保存sso client
func SaveSSOClient(c component.IContainer, m *sso.Client) {
c.Set("__SsoClient__", m)
}
//GetSSOClient 获取sso client
func GetSSOClient(c component.IContainer) *sso.Client {
return c.Get("__SsoClient__").(*sso.Client)
}
```

###### 1.3 使⽤sso client 实例
``` go
//此代码就是前端ssocallback页面调用的api
//验证用户登录的真实性并获取用户信息
//验证用户登录的合法性, 传入回调给子系统的code
//如果验证通过就返回用户信息(子系统要生成自己的jwt)
ctx.Log.Info("-------sso登录后去取登录用户---------")
if err := ctx.Request.Check("code"); err != nil {
    return context.NewError(context.ERR_NOT_ACCEPTABLE, fmt.Errorf("code不能为空"))
}
ctx.Log.Info("调用sso api 用code取用户信息")
data, err := model.GetSSOClient(u.c).CheckCodeLogin(ctx.Request.GetString("code"))
if err != nil {
    return err
}
ctx.Log.Info("设置用户的登录jwt")
ctx.Response.SetJWT(data)
return data


//获取用户的菜单数据(将原来自己的http调用改成下面的方式)
data, err := model.GetSSOClient(u.c).GetUserMenu(userID)

//获取子系统 系统信息(将原来自己的http调用改成下面的方式)
data, err := model.GetSSOClient(u.c).GetSystemInfo()

```
---

####2 相关⽅法的说明(输⼊输出)
###### 2.1  初始化
``` go
New(apiHost, ident, secret string)
```
输⼊参数   

| 参数  | 类型 | 　说明         |    
| :----: | :---: | :--------: |     
|apiHost |string |sso api host |      
|ident | string | ⼦系统标识(相当于英⽂名称)|   
|secret | string | ⼦系统秘钥|   

输出    
```
返回ssoClient 对象
```

###### 2.1  ⽤⼾登录验证以及返回⽤⼾信息(跳转登录后)
``` go
CheckCodeLogin(code string)
```
```
由于现在是跳转登录的⽅式,因此sso回调⼦系统地址后，⼦系统要验证登录的合性
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|code| string| 调转登录返回给⼦系统的code|


输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|UserID |string |⽤⼾标识|
|UserName |string |⽤⼾名称|
|RoleName |string |⻆⾊名称|
|SystemID |string |系统编号|
|ExtParams |string |⽤⼾扩展参数(⼀个json对象)|

###### 2.3  根据⽤⼾名获取⽤⼾信息
``` go
GetUserInfoByName(userName string)
```
输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userName| string| ⽤⼾名称|

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userName |string |⽤⼾名称|
|WxOpID |string |微信openID|
|ExtParams |string| ⽤⼾扩展参数(⼀个json对象)|
|UserID |string| ⽤⼾标识|

######2.4 获取⽤⼾在某个⼦系统下的菜单数据
```go
GetUserMenu(userID int)
```
输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userID| int |⽤⼾标识|

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|ID |string |菜单标识
|Name |string |菜单名称
|Level |string |级次
|IsOpen |string |是否展开
|Icon |string |图标
|SystemID |string |系统标识
|Parent |string |⽗级编号
|Path |string |路由地址
|Sortrank |string |排序编号
|Children |对象数组 |⼦菜单

###### 2.5 获取⼦系统信息
``` go
GetSystemInfo()
```
输⼊参数(无)

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|ID |string |系统标识
|Ident |string |系统ident(英⽂名称)
|Name |string |系统名称
|Theme |string |主题样式
|Layout |string |⻚⾯布局样式
|Logo |string |系统图标地址


###### 2.6 获取当前用户可访问的其他子系统
``` go
GetUserOtherSystems()
```
输⼊参数(无)

输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|ID |string |系统标识
|Ident |string |系统ident(英⽂名称)
|Name |string |系统名称
|IndexUrl |string |⼦系统地址 host


##### 3 需要注意的地方
###### 3.1 去掉原来不用的代码
###### 3.2 新增sso回调的api接口,代码已在上面
###### 3.3 将这个回调的api地址要配置在auth中，它相当于登录，不需要验证(403)
###### 3.4 修改原来app的配置, sso_api_host改成新的地址
```
线下:http://api.sso.18jiayou1.com:6689
线上：http://api.sso.18jiayou.com
```
###### 3.5 将原来调用 sso apiserver的菜单，系统信息,都改成新的方式