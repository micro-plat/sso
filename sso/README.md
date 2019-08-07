## sso api交互

SSOClient 就是在包装相关的api调用，简化用户调用

#### 1. 相关数据初始化
``` go

New(apiHost, ident, secret string)
apiHost: sso api地址
ident: 子系统标识(相当于英文名称)
secret: 子系统秘钥

返回ssoClient对象

```

#### 2. 登录

由于现在是跳转登录的方式,因此sso回调子系统地址后，子系统要验证登录的合法性 

``` go

CheckCodeLogin(code string)
code:这个code是回调子系统时带上的，此时要给sso验证，才能知道是登录的用户，如果验证成功
返回用户信息

```

#### 3. 根据用户名获取用户信息

``` go

GetUserInfoByName(userName string)
userName: 用户名称

```

#### 4. 获取用户在某个子系统下的菜单数据 

``` go

GetUserMenu(userID int)
userID: 用户标识

```

#### 5. 获取 

``` go

GetSystemInfo()

```




