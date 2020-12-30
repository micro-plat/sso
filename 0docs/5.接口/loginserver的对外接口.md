#### 1 API接口说明

现将所有调用的api调用都包成了sdk, 项目在 sdk/sso

接口签名方式:业务参数,timestamp(Unix时间)按照asc码排序,不需要任何链接字符拼接成签名原串raw(例如:ident123456timestamp25425325125userID52425)    
sign=md5(raw+secrect)   

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
|FullName |string |⽤⼾全名称|
|SysIdent |string |系统编码|
|RoleID |int |角色id|
|Status |int |用户状态|
|RoleName |string |⻆⾊名称|
|SystemID |int |系统编号|
|IndexURL |string |登录回调地址|
|ExtParams |string |⽤⼾扩展参数(⼀个json对象)|
|Source |string |用户来源|
|SourceID |string |用户来源id|
|LastLoginTime |string |最后登录时间|

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
|UserID |string |⽤⼾标识|
|UserName |string |⽤⼾名称|
|FullName |string |⽤⼾全名称|
|WxOpID |string |微信openID|
|Mobile |string |联系电话|
|Email |string |邮箱|
|Status |string |联系电话|
|ExtParams |string| ⽤⼾扩展参数(⼀个json对象)|


###### 1.3 获取⽤⼾在某个⼦系统下的菜单数据

``` go
GetUserMenu(userID int)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userID| int |⽤⼾标识|

输出下面对象列表

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
|IndexUrl |string |登录回调地址
|Logo |string |系统图标地址

###### 1.5 获取当前用户可访问的其他子系统

``` go
GetUserOtherSystems(userID string)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userID| string| ⽤⼾编号|

输出下面对象列表

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|ID |string |系统标识
|Ident |string |系统ident(英⽂名称)
|Name |string |系统名称
|Theme |string |主题样式
|Layout |string |⻚⾯布局样式
|IndexUrl |string |登录回调地址
|Logo |string |系统图标地址



###### 1.6 获取某来源所有的用户列表

``` go
GetAllUser(source string, sourceID string)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|source| string| 用户来源|
|sourceID| string| 用户来源id|

输出下面对象列表


|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|UserID |string |⽤⼾标识|
|UserName |string |⽤⼾名称|
|FullName |string |⽤⼾全名称|
|WxOpID |string |微信openID|
|Mobile |string |联系电话|
|Email |string |邮箱|
|Status |string |联系电话|
|ExtParams |string| ⽤⼾扩展参数(⼀个json对象)|



###### 1.7 忘记并修改密码

``` go
ForgetPwd(source, sourceID, possword string)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|source| string| 用户来源|
|sourceID| string| 用户来源id|
|possword| string| 新密码|

无返回值



###### 1.8 获取用户有权限的Tags

``` go
GetUserTags(UserID int, tags string)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|UserID| string| 用户编号|
|tags| string| tags|

返回值:
```json
[
    {
        "tag":"tag1",
        "display":false
    },
    {
        "tag":"tag2",
        "display":true
    }
]
```


###### 1.9 添加用户

``` go
AddUser(userName, mobile, fullName, targetIdent, source, sourceSecrect, sourceID string)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userName| string| 用户名|
|mobile| string| 联系电话|
|fullName| string| 用户全名|
|targetIdent| string| 系统编号|
|source| string| 用户来源|
|sourceID| string| 用户来源|
|sourceSecrect| string| 签名密钥|

密钥是默认密码:请联系服务提供方获取

无返回值

###### 1.10 系统登录

``` go
Login(userName, password string)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userName| string| 用户名|
|password| string| 联系电话,通过md5|


输出

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|UserID |string |⽤⼾标识|
|UserName |string |⽤⼾名称|
|FullName |string |⽤⼾全名称|
|SysIdent |string |系统编码|
|RoleID |int |角色id|
|Status |int |用户状态|
|RoleName |string |⻆⾊名称|
|SystemID |int |系统编号|
|IndexURL |string |登录回调地址|
|ExtParams |string |⽤⼾扩展参数(⼀个json对象)|
|Source |string |用户来源|
|SourceID |string |用户来源id|
|LastLoginTime |string |最后登录时间|



###### 1.11 系统登录

``` go
ChangePwd(userID int64, expassword, newpassword string)
```

输⼊参数

|参数 |类型|说明|
| -------------|:--------------:|:--------------:|
|userID| string| 用户名|
|expassword| string| 原密码,通过md5|
|newpassword| string| 新密码,通过md5|


无返回值



