package scheme

const sso_role_datapermission = `
drop table if exists sso_role_datapermission;

CREATE TABLE  sso_role_datapermission (
	id bigint(20)  not null AUTO_INCREMENT comment '功能编号' ,
	sys_id bigint(20)  not null  comment '系统编号' ,
	role_id bigint(20)  not null  comment '角色编号' ,
	permission_config_id bigint(20)  not null  comment '规则id' ,
	status tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
	create_time datetime default current_timestamp not null  comment '创建时间' ,
	PRIMARY KEY (id),
	KEY idx_sso_dataprem_sysrole (sys_id,role_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='角色与规则关联信息表';

	`
