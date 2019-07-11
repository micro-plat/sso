#相关数据表说明

### 1. 微信公众号配置[sso_system_info]

| 字段名             | 类型          | 默认值   | 为空 | 约束   | 描述                   |
| ----------------- | ------------ | :-----: | :--: | :---: | :--------------------- |
| id                | number(10) |           |  否  | PK,IS | id                     |
| name              | varchar2(32) |         |  否  |  IUS  | 系统名称                |
| index_url         | varchar2(64) |         |  否  |  IUS  | 首页地址                  |
| enable            | number(1)    |   1     |  否  |  IUS  | 状态 1：启用 0:禁用      |
| login_timeout     | number(6)    | 300     |  否  |  IS   | 超时时长           |
| logo              | varchar2(128)|         |  否  |  IUS  | 支付服务商编号         |
| pay_key           | varchar2(64) |         |  否  |  IUS  | 支付服务商 key         |
| sub_appid         | varchar2(64) |         |  否  |  IUS  | 子商户 app id          |
| sub_mchid         | varchar2(64) |         |  否  |  IUS  | 子商户编户号           |
| wechat_host       | varchar2(64) |         |  是  |  IUS  | 微信授权域名           |
| wcode_template_id | varchar2(64) |         |  是  |  IUS  | 微信验证码模板消息编号 |


### 1.2 用户表
sso_user_info

### 1.3 功能表
sso_system_menu

### 1.4 角色表
sso_role_info

### 1.5 用户角色关联表
sso_user_role