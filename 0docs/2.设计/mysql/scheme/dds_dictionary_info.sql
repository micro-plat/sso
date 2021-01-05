
	CREATE TABLE  dds_dictionary_info (
		id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		name varchar(64)  not null  comment '名称' ,
		value varchar(32)  not null  comment '值' ,
		type varchar(32)  not null  comment '类型' ,
		sort_no bigint(20) default 0 not null  comment '排序值' ,
		group_code varchar(32) default '*' not null  comment '分组编号' ,
		status tinyint(1)  not null  comment '状态 1: 禁用 0: 启用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='字典表';
