#### 1 API接口说明
现将所有调用的api调用都包成了sdk,项目在 sdk/sso

###### 1.1  ⽤⼾登录验证以及返回⽤⼾信息(跳转登录后)
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

###### 1.2  根据⽤⼾名获取⽤⼾信息
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

######1.3 获取⽤⼾在某个⼦系统下的菜单数据
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

###### 1.4 获取⼦系统信息
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


###### 1.5 获取当前用户可访问的其他子系统
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
