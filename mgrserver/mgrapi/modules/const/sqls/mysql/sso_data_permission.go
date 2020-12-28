package mysql

const sso_data_permission = `CREATE TABLE  sso_data_permission (
		id number(20)  not null  comment 'id' ,
		sys_id number(20)  not null  comment '系统编号' ,
		ident varchar2(32)  not null  comment '系统标识' ,
		name varchar2(128)  not null  comment '名称' ,
		table_name varchar2(128)  not null  comment '表名' ,
		operate_action varchar2(64)  not null  comment '操作动作' ,
		rules LONGTEXT    comment '规则json' ,
		remark varchar2(256)  not null  comment '说明' ,
		status number(4) default 0 not null  comment '状态 0: 启用 1: 禁用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='数据权限规则表';`
