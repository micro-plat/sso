## sso sdk使用说明

在跳转登录返回后,子系统需要验证登录用户、获取菜单、获取系统信息等，为了降低使用接口的复杂度,将这些接口调用包装成sdk

####1 子系统服务端修改点
###### 1.1 在init()⽅法中注⼊sso client

``` go
import "github.com/micro-plat/sso/sso"
ssoCleint, err := sso.New(conf.SsoApiHost, conf.Ident, conf.Secret)
if err != nil {
return err
}
model.SaveSSOClient(c, ssoCleint) //将sso client 保存起来
```

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
//验证用户登录的真实性并获取用户信息
data, err := model.GetSSOClient(u.c).CheckCodeLogin(code)

//获取用户的菜单数据
//获取子系统信息
```
---

####2 相关⽅法的说明(输⼊输出)
###### 2.1  初始化
``` go
New(apiHost, ident, secret string)
```
输⼊参数  
|参数 |类型|说明|  
| -------------|:--------------:|:--------------:|  
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
|IndexUrl |string |sso登录后⼦系统回调地址
|Layout |string |⻚⾯布局样式
|Logo |string |系统图标地址
