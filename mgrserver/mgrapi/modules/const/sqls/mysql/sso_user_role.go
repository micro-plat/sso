package mysql

const sso_user_role = `CREATE TABLE  sso_user_role (
		id bigint(20)  not null  comment 'id' ,
		user_id bigint(20)  not null  comment '用户编号' ,
		sys_id bigint(20)  not null  comment '系统编号' ,
		role_id bigint(20)  not null  comment '角色编号' ,
		enable tinyint(1) default 1 not null  comment '状态 0: 启用 1: 禁用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户角色关联表';
`
