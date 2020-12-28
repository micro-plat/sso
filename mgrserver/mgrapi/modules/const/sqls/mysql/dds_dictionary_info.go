package mysql

const dds_dictionary_info = `CREATE TABLE  dds_dictionary_info (
		id number(20)  not null  comment 'id' ,
		name varchar2(64)  not null  comment '名称' ,
		value varchar2(32)  not null  comment '值' ,
		type varchar2(32)  not null  comment '类型' ,
		sort_no number(20) default 0 not null  comment '排序值' ,
		status number(1)  not null  comment '状态 1: 禁用 0: 启用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='字典表';`
