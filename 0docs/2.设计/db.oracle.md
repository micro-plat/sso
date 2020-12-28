#相关数据表说明

### 1. 系统信息[sso_system_info]

| 字段名        | 类型          |              默认值              | 为空  |  约束  | 描述                          |
| ------------- | ------------- | :------------------------------: | :---: | :----: | :---------------------------- |
| id            | number(20)    |                                  |  否   | PK, IS | id                            |
| name          | varchar2(32)  |                                  |  否   |  IUS   | 系统名称                      |
| index_url     | varchar2(64)  |                                  |  否   |  IUS   | 首页地址                      |
| enable        | number(1)     |                1                 |  否   |  IUS   | 状态 1：启用 0: 禁用          |
| login_timeout | number(6)     |               300                |  否   |   IS   | 超时时长                      |
| logo          | varchar2(128) |                                  |  否   |  IUS   | logo                          |
| theme         | varchar2(128) |           bg-parimary            |  否   |  IUS   | 主题样式                      |
| layout        | varchar2(128) | app-header-fixed app-aside-fixed |  否   |  IUS   | 页面布局样式                  |
| ident         | varchar2(16)  |                                  |  否   |  IUS   | 唯一标识                      |
| login_url     | varchar2(64)  |                                  |  是   |  IUS   | 登录地址                      |
| wechat_status | number(2)     |                1                 |  否   |  IUS   | 微信功能状态 1: 开启, 0: 关闭 |
| secret        | varchar2(32)  |                                  |  否   |  IUS   | 签名密钥                      |

### 2. 功能表(菜单数据)[sso_system_menu]

| 字段名      | 类型          | 默认值  | 为空  |  约束  | 描述                 |
| ----------- | ------------- | :-----: | :---: | :----: | :------------------- |
| id          | number(20)    |         |  否   | PK, IS | 功能编号             |
| name        | varchar2(64)  |         |  否   |  IUS   | 功能名称             |
| parent      | number(20)    |         |  否   |  IUS   | 父级编号             |
| sys_id      | number(20)    |         |  否   |  IUS   | 系统编号             |
| level_id    | number(2)     |         |  否   |  IUS   | 等级                 |
| icon        | varchar2(64)  |         |  是   |  IUS   | 图标                 |
| path        | varchar2(256) |         |  否   |  IUS   | 地址                 |
| enable      | number(1)     |    0    |  否   |  IUS   | 状态 1: 禁用 0: 正常 |
| create_time | date          | sysdate |  否   |  IUS   | 创建时间             |
| sortrank    | number(20)    |         |  否   |  IUS   | 排序编号             |
| is_open     | number(1)     |    0    |  是   |  IUS   | 是否展开             |

### 3. 用户表[sso_user_info]

| 字段名          | 类型           | 默认值  | 为空  |  约束  | 描述                         |
| --------------- | -------------- | :-----: | :---: | :----: | :--------------------------- |
| user_id         | number(10)     |         |  否   | PK, IS | id                           |
| full_name       | varchar2(10)   |         |  否   | PK, IS | 用户全名                     |
| user_name       | varchar2(64)   |         |  否   |  IUS   | 用户名                       |
| password        | varchar2(32)   |         |  否   |  IUS   | 密码                         |
| email           | varchar2(32)   |         |  是   |  IUS   | email                        |
| status          | number(1)      |    1    |  否   |   IS   | 状态 0: 正常 1: 锁定 2: 禁用 |
| mobile          | varchar2(12)   |         |  否   |  IUS   | 电话号码                     |
| wx_openid       | varchar2(64)   |         |  是   |  IUS   | 微信openid                   |
| create_time     | date           | sysdate |  否   |  IUS   | 创建时间                     |
| changepwd_times | number(2）     |    0    |  否   |  IUS   | 密码修改次数                 |
| source          | varchar2(36）  |         |  否   |  IUS   | 来源                         |
| source_id       | varchar2(120） |         |  否   |  IUS   | 来源id                       |
| ext_params      | varchar2(1024) |         |  是   |  IUS   | 扩展参数                     |

### 4. 角色表[sso_role_info]

| 字段名      | 类型         | 默认值  | 为空  |  约束  | 描述                 |
| ----------- | ------------ | :-----: | :---: | :----: | :------------------- |
| role_id     | number(20)   |         |  否   | PK, IS | 角色id               |
| name        | varchar2(64) |         |  否   |  IUS   | 角色名称             |
| status      | number(1)    |    0    |  否   |   IS   | 状态 1: 禁用 0: 正常 |
| create_time | date         | sysdate |  否   |  IUS   | 创建时间             |

### 5. 角色表[sso_role_menu]

| 字段名      | 类型       | 默认值  | 为空  |  约束  | 描述                 |
| ----------- | ---------- | :-----: | :---: | :----: | :------------------- |
| id          | number(20) |         |  否   | PK, IS | id                   |
| sys_id      | number(20) |         |  否   |  IUS   | 系统id               |
| role_id     | number(20) |    0    |  否   |   IS   | 角色id               |
| menu_id     | number(20) |    0    |  否   |   IS   | 菜单id               |
| enable      | number(1)  |    0    |  否   |   IS   | 状态 1: 禁用 0: 正常 |
| create_time | date       | sysdate |  否   |  IUS   | 创建时间             |
| sortrank    | number(20) |    0    |  否   |   IS   | 排序编号             |


### 6. 角色与规则关联信息表[sso_role_datapermission]

| 字段名               | 类型       | 默认值  | 为空  |  约束  | 描述     |
| -------------------- | ---------- | :-----: | :---: | :----: | :------- |
| id                   | number(20) |         |  否   | PK, IS | 功能编号 |
| sys_id               | number(20) |         |  否   |  IUS   | 系统编号 |
| role_id              | number(20) |         |  否   |  IUS   | 角色编号 |
| permission_config_id | number(20) |         |  否   |  IUS   | 规则id   |
| create_time          | date       | sysdate |  否   |   IS   | 创建时间 |


### 7. 用户角色关联表[sso_user_role]

| 字段名  | 类型       | 默认值 | 为空  |  约束  | 描述                 |
| ------- | ---------- | :----: | :---: | :----: | :------------------- |
| id      | number(20) |        |  否   | PK, IS | id                   |
| user_id | number(20) |        |  否   |  IUS   | 用户编号             |
| sys_id  | number(20) |        |  否   |  IUS   | 系统编号             |
| role_id | number(20) |        |  否   |  IUS   | 角色编号             |
| enable  | number(1)  |   1    |  否   |   IS   | 状态 0: 启用 1: 禁用 |

### 8. 数据权限规则表[sso_data_permission]

| 字段名         | 类型          | 默认值 | 为空  |  约束  | 描述                 |
| -------------- | ------------- | :----: | :---: | :----: | :------------------- |
| id             | number(20)    |        |  否   | PK, IS | id                   |
| sys_id         | number(20)    |        |  否   |  IUS   | 系统编号             |
| ident          | varchar2(32)  |        |  否   |  IUS   | 系统标识             |
| name           | varchar2(128) |        |  否   |   IS   | 名称                 |
| table_name     | varchar2(128) |        |  否   |   IS   | 表名                 |
| operate_action | varchar2(64)  |        |  否   |   IS   | 操作动作             |
| rules          | CLOB          |        |  是   |   IS   | 规则json             |
| remark         | varchar2(256) |        |  否   |   IS   | 说明                 |
| status         | number(4)     |   0    |  否   |   IS   | 状态 0: 启用 1: 禁用 |


### 9. 操作日志[sso_operate_log]

| 字段名      | 类型          | 默认值 | 为空  |  约束  | 描述                                                                                |
| ----------- | ------------- | :----: | :---: | :----: | :---------------------------------------------------------------------------------- |
| id          | number(20)    |        |  否   | PK, IS | id                                                                                  |
| type        | number(2)     |        |  否   |  IUS   | 类型,10.登录操作 20.系统数据操作  30.角色数据操作  40.菜单数据操作  50.用户数据操作 |
| sys_id      | number(20)    |        |  否   |  IUS   | 系统编号                                                                            |
| user_id     | number(20)    |        |  否   |   IS   | 操作人id                                                                            |
| create_time | date          |        |  否   |   IS   | 创建时间                                                                            |
| content     | varchar2(512) |        |  否   |   IS   | 内容I                                                                               |

### 10. 字典表[dds_dictionary_info]

| 字段名  | 类型         | 默认值 | 为空  |  约束  | 描述                 |
| ------- | ------------ | :----: | :---: | :----: | :------------------- |
| id      | number(20)   |        |  否   | PK, IS | id                   |
| name    | varchar2(64) |        |  否   |  IUS   | 名称                 |
| value   | varchar2(32) |        |  否   |  IUS   | 值                   |
| type    | varchar2(32) |        |  否   |  IUS   | 类型                 |
| sort_no | number(20)   |   0    |  否   |   IS   | 排序值               |
| status  | number(1)    |        |  否   |   IS   | 状态 1: 禁用 0: 启用 |


### 11. 地区表[dds_area_info]

| 字段名       | 类型          | 默认值 | 为空  |  约束  | 描述                 |
| ------------ | ------------- | :----: | :---: | :----: | :------------------- |
| canton_code  | varchar2(32)  |        |  否   | PK, IS | 区域编号             |
| chinese_name | varchar2(128) |        |  否   |  IUS   | 中文名称             |
| parent_code  | varchar2(32)  |        |  是   |  IUS   | 父级编号             |
| grade        | number(1)     |        |  否   |  IUS   | 行政级别             |
| full_spell   | varchar2(20)  |        |  是   |   IS   | 英文/全拼            |
| simple_spell | varchar2(20)  |        |  是   |   IS   | 简拼                 |
| sort_no      | number(20)    |   0    |  否   |   IS   | 排序值               |
| status       | number(1)     |   0    |  否   |   IS   | 状态 1: 禁用 0: 启用 |
