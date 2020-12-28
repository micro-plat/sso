package mysql

const sso_role_menu = `CREATE TABLE  sso_role_menu (
	id number(20)  not null  comment 'id' ,
	sys_id number(20)  not null  comment '系统id' ,
	role_id number(20) default 0 not null  comment '角色id' ,
	menu_id number(20) default 0 not null  comment '菜单id' ,
	enable number(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
	create_time DATETIME default current_timestamp not null  comment '创建时间' ,
	sortrank number(20) default 0 not null  comment '排序编号' ,
	PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表';`
