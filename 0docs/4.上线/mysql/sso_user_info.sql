
	CREATE TABLE  sso_user_info (
		user_id number(10)  not null  comment 'id' ,
		full_name varchar2(10)  not null  comment '用户全名' ,
		user_name varchar2(64)  not null  comment '用户名' ,
		password varchar2(32)  not null  comment '密码' ,
		email varchar2(32)    comment 'email' ,
		status number(1) default 1 not null  comment '状态 0: 正常 1: 锁定 2: 禁用' ,
		mobile varchar2(12)  not null  comment '电话号码' ,
		wx_openid varchar2(64)    comment '微信openid' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		changepwd_times number(2) default 0 not null  comment '密码修改次数' ,
		source varchar2(36)  not null  comment '来源' ,
		source_id varchar2(120)  not null  comment '来源id' ,
		ext_params varchar2(1024)    comment '扩展参数' ,
		PRIMARY KEY (user_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户表';
