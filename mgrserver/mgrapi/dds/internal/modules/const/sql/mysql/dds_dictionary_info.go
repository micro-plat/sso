package mysql

const dds_dictionary_info = `
DROP TABLE if exists dds_dictionary_info;
CREATE TABLE  dds_dictionary_info (
	id BIGINT(10)  not null AUTO_INCREMENT comment '编号' ,
	name VARCHAR(128)  not null  comment '名称' ,
	value VARCHAR(4000)  not null  comment '值' ,
	type VARCHAR(32)  not null  comment '类型' ,
	sort_no TINYINT(2) default 0 not null  comment '排序值' ,
	status TINYINT(1)  not null  comment '状态' ,
	group_code VARCHAR(32) default '*' not null  comment '分组编号' ,
	PRIMARY KEY (id),
	KEY IDX_DICTIONARY_INFO_TYPE (type)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='字典表';`
