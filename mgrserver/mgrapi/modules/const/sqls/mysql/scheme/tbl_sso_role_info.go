package scheme

const sso_role_info = `
drop table if exists sso_role_info;


CREATE TABLE  sso_role_info (
	role_id bigint(20)  not null AUTO_INCREMENT comment '角色id' ,
	name varchar(64)  not null  comment '角色名称' ,
	status tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
	create_time datetime default current_timestamp not null  comment '创建时间' ,
	PRIMARY KEY (role_id),
	UNIQUE KEY unq_sso_role_info_name (name)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
`
