package scheme

const sso_data_permission = `
drop table if exists sso_data_permission;

CREATE TABLE  sso_data_permission (
	id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
	sys_id bigint(20)  not null  comment '系统编号' ,
	ident varchar(32)  not null  comment '系统标识' ,
	name varchar(128)  not null  comment '名称' ,
	table_name varchar(128)  not null  comment '表名' ,
	operate_action varchar(64)  not null  comment '操作动作' ,
	rules text    comment '规则json' ,
	remark varchar(256)  not null  comment '说明' ,
	status tinyint(4) default 0 not null  comment '状态 0: 启用 1: 禁用' ,
	PRIMARY KEY (id),
	UNIQUE KEY unq_sso_dataprem_identn (sys_id,name)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='数据权限规则表';
`
