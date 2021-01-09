package scheme

const sso_role_menu = `
drop table if exists sso_role_menu;


CREATE TABLE  sso_role_menu (
	id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
	sys_id bigint(20)  not null  comment '系统id' ,
	role_id bigint(20) default 0 not null  comment '角色id' ,
	menu_id bigint(20) default 0 not null  comment '菜单id' ,
	enable tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
	create_time datetime default current_timestamp not null  comment '创建时间' ,
	sortrank bigint(20) default 0 not null  comment '排序编号' ,
	PRIMARY KEY (id),
	UNIQUE KEY unq_sso_rolem_srmenu (menu_id,role_id,sys_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
`
