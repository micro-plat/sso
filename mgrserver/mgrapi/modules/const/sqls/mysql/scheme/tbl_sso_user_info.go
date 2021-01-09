package scheme

const sso_user_info = `
drop table if exists sso_user_info;


CREATE TABLE  sso_user_info (
	user_id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
	full_name varchar(32)  not null  comment '用户全名' ,
	user_name varchar(64)  not null  comment '用户名' ,
	password varchar(32)  not null  comment '密码' ,
	email varchar(32)    comment 'email' ,
	status tinyint(1) default 1 not null  comment '状态 0: 正常 1: 锁定 2: 禁用' ,
	mobile varchar(12)  not null  comment '电话号码' ,
	wx_openid varchar(64)    comment '微信openid' ,
	create_time datetime default current_timestamp not null  comment '创建时间' ,
	changepwd_times bigint(20) default 0 not null  comment '密码修改次数' ,
	ext_params varchar(1024)    comment '扩展参数' ,
	last_login_time datetime    comment '最后登录时间' ,
	source_id varchar(128) default '0' not null  comment '来源id' ,
	source varchar(36) default '''' not null  comment '来源' ,
	PRIMARY KEY (user_id),
	UNIQUE KEY unq_sso_user_info_user_name (user_name),
	
	KEY idx_sso_userinfo_source (source_id,source)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

	`
