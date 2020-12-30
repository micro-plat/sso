#相关数据表说明

### 1. 系统信息[sso_system_info]

| 字段名        | 类型         |              默认值              | 为空  |    约束    | 描述                          |
| ------------- | ------------ | :------------------------------: | :---: | :--------: | :---------------------------- |
| id            | bigint(20)   |                                  |  否   | PK, SEQ | id                            |
| name          | varchar(32)  |                                  |  否   |    UNQ, IDX(idx_sso_sysinfo_urlident, 1)     | 系统名称                      |
| index_url     | varchar(64)  |                                  |  是   |        | 首页地址                      |
| enable        | tinyint(1)   |                1                 |  否   |        | 状态 1：启用 0: 禁用          |
| login_timeout | int(11)      |               300                |  否   |        | 超时时长                      |
| logo          | varchar(128) |                                  |  是   |        | logo                          |
| theme         | varchar(128) |           bg-parimary            |  是   |        | 主题样式                      |
| layout        | varchar(128) | app-header-fixed app-aside-fixed |  是   |        | 页面布局样式                  |
| ident         | varchar(32)  |                                  |  否   |    IDX(idx_sso_sysinfo_urlident, 2)     | 唯一标识                      |
| login_url     | varchar(64)  |                                  |  是   |        | 登录地址                      |
| wechat_status | tinyint(2)   |                1                 |  否   |        | 微信功能状态 1: 开启, 0: 关闭 |
| secret        | varchar(32)  |                                  |  是   |        | 签名密钥                      |

### 2. 功能表(菜单数据)[sso_system_menu]

| 字段名      | 类型         |      默认值       | 为空  |  约束  | 描述                  |
| ----------- | ------------ | :---------------: | :---: | :----: | :-------------------- |
| id          | bigint(20)   |                   |  否   | PK, SEQ | 功能编号            |
| name        | varchar(64)  |                   |  否   |     | 功能名称              |
| parent      | bigint(20)   |                   |  否   |     | 父级编号              |
| sys_id      | bigint(20)   |                   |  否   |     | 系统编号              |
| level_id    | tinyint(2)   |                   |  否   |     | 等级                  |
| icon        | varchar(64)  |                   |  是   |     | 图标                  |
| path        | varchar(256) |                   |  否   |     | 地址                  |
| enable      | tinyint(1)   |         0         |  否   |     | 状态 1: 禁用 0: 正常  |
| create_time | datetime     | current_timestamp |  否   |     | 创建时间              |
| sortrank    | bigint(20)   |                   |  否   |     | 排序编号              |
| is_open     | tinyint(1)   |         0         |  是   |     | 是否展开              |

### 3. 用户表[sso_user_info]

| 字段名          | 类型          |      默认值       | 为空  |  约束  | 描述                         |
| --------------- | ------------- | :---------------: | :---: | :----: | :--------------------------- |
| user_id         | bigint(20)    |                   |  否   | PK, SEQ | id                           |
| full_name       | varchar(32)   |                   |  否   |  | 用户全名                     |
| user_name       | varchar(64)   |                   |  否   |   UNQ   | 用户名                       |
| password        | varchar(32)   |                   |  否   |     | 密码                         |
| email           | varchar(32)   |                   |  是   |     | email                        |
| status          | tinyint(1)    |         1         |  否   |      | 状态 0: 正常 1: 锁定 2: 禁用 |
| mobile          | varchar(12)   |                   |  否   |     | 电话号码                     |
| wx_openid       | varchar(64)   |                   |  是   |     | 微信openid                   |
| create_time     | datetime      | current_timestamp |  否   |     | 创建时间                     |
| changepwd_times | bigint(20)    |         0         |  否   |     | 密码修改次数                 |
| ext_params      | varchar(1024) |                   |  是   |     | 扩展参数                     |
| last_login_time | datetime      |                   |  是   |     | 最后登录时间                 |
| source_id       | varchar(128)  |         0         |  否   |   IDX(idx_sso_userinfo_source, 1)   | 来源id                       |
| source          | varchar(36)   |        ''         |  否   |   IDX(idx_sso_userinfo_source, 2)   | 来源                         |

### 4. 角色表[sso_role_info]

| 字段名      | 类型        |      默认值       | 为空  |    约束    | 描述                 |
| ----------- | ----------- | :---------------: | :---: | :--------: | :------------------- |
| role_id     | bigint(20)  |                   |  否   | PK, SEQ | 角色id               |
| name        | varchar(64) |                   |  否   |     UNQ     | 角色名称             |
| status      | tinyint(1)  |         0         |  否   |          | 状态 1: 禁用 0: 正常 |
| create_time | datetime    | current_timestamp |  否   |         | 创建时间             |

### 5. 角色表[sso_role_menu]

| 字段名      | 类型       |      默认值       | 为空  |    约束    | 描述                 |
| ----------- | ---------- | :---------------: | :---: | :--------: | :------------------- |
| id          | bigint(20) |                   |  否   | PK, SEQ | id                   |
| sys_id      | bigint(20) |                   |  否   |     UNQ(unq_sso_rolem_srmenu, 3)     | 系统id               |
| role_id     | bigint(20) |         0         |  否   |      UNQ(unq_sso_rolem_srmenu, 2)     | 角色id               |
| menu_id     | bigint(20) |         0         |  否   |      UNQ(unq_sso_rolem_srmenu, 1)     | 菜单id               |
| enable      | tinyint(1) |         0         |  否   |          | 状态 1: 禁用 0: 正常 |
| create_time | datetime   | current_timestamp |  否   |         | 创建时间             |
| sortrank    | bigint(20) |         0         |  否   |          | 排序编号             |

### 6. 角色与规则关联信息表[sso_role_datapermission]

| 字段名               | 类型       |      默认值       | 为空  |    约束    | 描述                 |
| -------------------- | ---------- | :---------------: | :---: | :--------: | :------------------- |
| id                   | bigint(20) |                   |  否   | PK, SEQ | 功能编号             |
| sys_id               | bigint(20) |                   |  否   |  IDX(idx_sso_dataprem_sysrole, 1)     | 系统编号             |
| role_id              | bigint(20) |                   |  否   |  IDX(idx_sso_dataprem_sysrole, 2)     | 角色编号             |
| permission_config_id | bigint(20) |                   |  否   |         | 规则id               |
| status               | tinyint(1) |         0         |  否   |          | 状态 1: 禁用 0: 正常 |
| create_time          | datetime   | current_timestamp |  否   |          | 创建时间             |

### 7. 用户角色关联表[sso_user_role]

| 字段名  | 类型       | 默认值 | 为空  |    约束    | 描述                 |
| ------- | ---------- | :----: | :---: | :--------: | :------------------- |
| id      | bigint(20) |        |  否   | PK, SEQ | id                   |
| user_id | bigint(20) |        |  否   |      UNQ(unq_sso_userrole_usr, 1)     | 用户编号             |
| sys_id  | bigint(20) |        |  否   |      UNQ(unq_sso_userrole_usr, 2)     | 系统编号             |
| role_id | bigint(20) |        |  否   |      UNQ(unq_sso_userrole_usr, 3)     | 角色编号             |
| enable  | tinyint(1) |   1    |  否   |          | 状态 0: 启用 1: 禁用 |

### 8. 数据权限规则表[sso_data_permission]

| 字段名         | 类型         | 默认值 | 为空  |    约束    | 描述                 |
| -------------- | ------------ | :----: | :---: | :--------: | :------------------- |
| id             | bigint(20)   |        |  否   | PK, SEQ | id                   |
| sys_id         | bigint(20)   |        |  否   |     UNQ(unq_sso_dataprem_identn, 1)     | 系统编号             |
| ident          | varchar(32)  |        |  否   |         | 系统标识             |
| name           | varchar(128) |        |  否   |      UNQ(unq_sso_dataprem_identn, 2)     | 名称                 |
| table_name     | varchar(128) |        |  否   |          | 表名                 |
| operate_action | varchar(64)  |        |  否   |          | 操作动作             |
| rules          | text         |        |  是   |          | 规则json             |
| remark         | varchar(256) |        |  否   |          | 说明                 |
| status         | tinyint(4)   |   0    |  否   |          | 状态 0: 启用 1: 禁用 |

### 9. 操作日志[sso_operate_log]

| 字段名      | 类型         |      默认值       | 为空  |    约束    | 描述                                                                                |
| ----------- | ------------ | :---------------: | :---: | :--------: | :---------------------------------------------------------------------------------- |
| id          | bigint(20)   |                   |  否   | PK, SEQ | id                                                                                  |
| type        | tinyint(2)   |                   |  否   |         | 类型, 10. 登录操作 20. 系统数据操作  30. 角色数据操作  40. 菜单数据操作  50. 用户数据操作 |
| sys_id      | bigint(20)   |                   |  否   |         | 系统编号                                                                            |
| user_id     | bigint(20)   |                   |  否   |      IDX     | 操作人id                                                                            |
| create_time | datetime     | current_timestamp |  否   |          | 创建时间                                                                            |
| content     | varchar(512) |                   |  否   |          | 内容I                                                                               |

### 10. 字典表[dds_dictionary_info]

| 字段名  | 类型        | 默认值 | 为空  |  约束  | 描述                 |
| ------- | ----------- | :----: | :---: | :----: | :------------------- |
| id      | bigint(20)  |        |  否   | PK, SEQ | id                   |
| name    | varchar(64) |        |  否   |     | 名称                 |
| value   | varchar(32) |        |  否   |     | 值                   |
| type    | varchar(32) |        |  否   |     | 类型                 |
| sort_no | bigint(20)  |   0    |  否   |      | 排序值               |
| status  | tinyint(1)  |        |  否   |      | 状态 1: 禁用 0: 启用 |

### 11. 地区表[dds_area_info]

| 字段名       | 类型         | 默认值 | 为空  |  约束  | 描述                 |
| ------------ | ------------ | :----: | :---: | :----: | :------------------- |
| canton_code  | varchar(32)  |        |  否   | PK, | 区域编号             |
| chinese_name | varchar(128) |        |  否   |     | 中文名称             |
| parent_code  | varchar(32)  |        |  是   |   IDX   | 父级编号             |
| grade        | tinyint(1)   |        |  否   |     | 行政级别             |
| full_spell   | varchar(20)  |        |  是   |      | 英文/全拼            |
| simple_spell | varchar(20)  |        |  是   |      | 简拼                 |
| sort_no      | bigint(20)   |   0    |  否   |      | 排序值               |
| status       | tinyint(1)   |   0    |  否   |      | 状态 1: 禁用 0: 启用 |
