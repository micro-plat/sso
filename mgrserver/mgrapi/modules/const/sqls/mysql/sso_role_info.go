package mysql

const sso_role_info = `CREATE TABLE  sso_role_info (
	role_id number(20)  not null  comment '角色id' ,
	name varchar2(64)  not null  comment '角色名称' ,
	status number(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
	create_time DATETIME default current_timestamp not null  comment '创建时间' ,
	PRIMARY KEY (role_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表';`
