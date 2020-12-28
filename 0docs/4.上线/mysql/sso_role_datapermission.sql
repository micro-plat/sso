
	CREATE TABLE  sso_role_datapermission (
		id number(20)  not null  comment '功能编号' ,
		sys_id number(20)  not null  comment '系统编号' ,
		role_id number(20)  not null  comment '角色编号' ,
		permission_config_id number(20)  not null  comment '规则id' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色与规则关联信息表';
