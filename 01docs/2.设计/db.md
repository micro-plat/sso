#相关数据表说明

### 1. 系统信息[sso_system_info]

| 字段名             | 类型          | 默认值   | 为空 | 约束   | 描述                   |
| ----------------- | ------------ | :-----: | :--: | :---: | :--------------------- |
| id                | number(10) |           |  否  | PK,IS | id                     |
| name              | varchar2(32) |         |  否  |  IUS  | 系统名称                |
| index_url         | varchar2(64) |         |  否  |  IUS  | 首页地址                |
| enable            | number(1)    |   1     |  否  |  IUS  | 状态 1：启用 0:禁用      |
| login_timeout     | number(6)    | 300     |  否  |  IS   | 超时时长                |
| logo              | varchar2(128)|         |  否  |  IUS  | logo                   |
| theme             | varchar2(128) |        |  否  |  IUS  | 主题样式                |
| layout            | varchar2(128) |        |  否  |  IUS  | 页面布局样式             |
| ident             | varchar2(16) |         |  否  |  IUS  | 唯一标识                 |
| login_url         | varchar2(64) |         |  是  |  IUS  | 登录地址                 |
| wechat_status     | number(1)    |    1    |  否  |  IUS  | 微信功能状态 1:开启,0:关闭 |
| secret            | varchar2(32) |         |  否  |  IUS  | 签名密钥                 |


### 1.2 用户表[sso_user_info]

| 字段名             | 类型          | 默认值   | 为空 | 约束   | 描述                   |
| ----------------- | ------------ | :-----: | :--: | :---: | :--------------------- |
| user_id           | number(10) |           |  否  | PK,IS | id                     |
| user_name         | varchar2(64) |         |  否  |  IUS  | 用户名                  |
| password          | varchar2(32) |         |  否  |  IUS  | 密码                    |
| email             | varchar2(32) |         |  否  |  IUS  | email      |
| status            | number(1)    |   1     |  否  |  IS   | 状态 0:正常 1:锁定 2:禁用 |
| mobile            | varchar2(12)|          |  否  |  IUS  | 电话号码                 |
| wx_openid         | varchar2(64) |         |  否  |  IUS  | 微信openid              |
| create_time       | date         |         |  否  |  IUS  | 创建时间                  |
| changepwd_times   | number(2）   |         |  否  |  IUS  | 密码修改次数               |
| ext_params        | varchar2(1024)|        |  是  |  IUS  | 扩展参数                 |


### 1.3 功能表(菜单数据)[sso_system_menu]

| 字段名             | 类型          | 默认值   | 为空 | 约束   | 描述                   |
| ----------------- | ------------ | :-----: | :--: | :---: | :--------------------- |
| id                | number(10) |           |  否  | PK,IS | id                     |
| sys_id            | number(10) |           |  否  |  IUS  | 系统编号                 |
| role_id           | number(10) |           |  否  |  IUS  | 角色编号                 |
| menu_id           | number(10) |           |  否  |  IUS  | 菜单编号                 |
| enable            | number(1)  |   1     |  否  |  IS   | 状态 1:正常 0:禁用        |
| create_time       | date         |         |  否  |  IUS  | 创建时间                  |
| changepwd_times   | number(2）   |         |  否  |  IUS  | 密码修改次数               |
| ext_params        | varchar2(1024)|        |  是  |  IUS  | 扩展参数                 |

### 1.4 角色表
sso_role_info

### 1.5 用户角色关联表
sso_user_role

# 2 表之间的关联关系
系统下面有菜单
角色关联菜单
用户关联角色

用户没有挂在系统下面