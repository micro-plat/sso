# 对外接口明细

### 1. 子系统远程登录
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

### 2. 子系统获取用户菜单数据
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


### 3. 子系统,获取用户信息
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

### 3. 子系统用户修改密码
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

### 3. 子系统获取系统信息
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