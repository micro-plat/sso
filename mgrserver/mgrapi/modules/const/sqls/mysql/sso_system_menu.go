package mysql

const sso_system_menu = `CREATE TABLE  sso_system_menu (
	id number(20)  not null  comment '功能编号' ,
	name varchar2(64)  not null  comment '功能名称' ,
	parent number(20)  not null  comment '父级编号' ,
	sys_id number(20)  not null  comment '系统编号' ,
	level_id number(2)  not null  comment '等级' ,
	icon varchar2(64)    comment '图标' ,
	path varchar2(256)  not null  comment '地址' ,
	enable number(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
	create_time DATETIME default current_timestamp not null  comment '创建时间' ,
	sortrank number(20)  not null  comment '排序编号' ,
	is_open number(1) default 0   comment '是否展开' ,
	PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='功能表(菜单数据)';`
