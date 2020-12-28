
	CREATE TABLE  sso_operate_log (
		id number(20)  not null  comment 'id' ,
		type number(2)  not null  comment '类型,10.登录操作 20.系统数据操作  30.角色数据操作  40.菜单数据操作  50.用户数据操作' ,
		sys_id number(20)  not null  comment '系统编号' ,
		user_id number(20)  not null  comment '操作人id' ,
		create_time DATETIME  not null  comment '创建时间' ,
		content varchar2(512)  not null  comment '内容I' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='操作日志';
