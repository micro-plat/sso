package mysql

const sso_user_info = `CREATE TABLE  sso_user_info (
	user_id bigint(20)  not null  comment 'id' ,
	full_name varchar(10)  not null  comment '用户全名' ,
	user_name varchar(64)  not null  comment '用户名' ,
	password varchar(32)  not null  comment '密码' ,
	email varchar(32)    comment 'email' ,
	status tinyint(1) default 1 not null  comment '状态 0: 正常 1: 锁定 2: 禁用' ,
	mobile varchar(12)  not null  comment '电话号码' ,
	wx_openid varchar(64)    comment '微信openid' ,
	create_time DATETIME default current_timestamp not null  comment '创建时间' ,
	changepwd_times tinyint(2) default 0 not null  comment '密码修改次数' ,
	source varchar(36)  not null  comment '来源' ,
	source_id varchar(120)  not null  comment '来源id' ,
	ext_params varchar(1024)    comment '扩展参数' ,
	PRIMARY KEY (user_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户表';`
