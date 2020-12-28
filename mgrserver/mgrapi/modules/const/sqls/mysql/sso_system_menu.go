package mysql

const sso_system_menu = `CREATE TABLE  sso_system_menu (
	id bigint(20)  not null  comment '功能编号' ,
	name varchar(64)  not null  comment '功能名称' ,
	parent bigint(20)  not null  comment '父级编号' ,
	sys_id bigint(20)  not null  comment '系统编号' ,
	level_id tinyint(2)  not null  comment '等级' ,
	icon varchar(64)    comment '图标' ,
	path varchar(256)  not null  comment '地址' ,
	enable tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
	create_time DATETIME default current_timestamp not null  comment '创建时间' ,
	sortrank bigint(20)  not null  comment '排序编号' ,
	is_open tinyint(1) default 0   comment '是否展开' ,
	PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='功能表(菜单数据)';`
