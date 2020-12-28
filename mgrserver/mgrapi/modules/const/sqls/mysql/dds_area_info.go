package mysql

const dds_area_info = `CREATE TABLE  dds_area_info (
		canton_code varchar(32)  not null  comment '区域编号' ,
		chinese_name varchar(128)  not null  comment '中文名称' ,
		parent_code varchar(32)    comment '父级编号' ,
		grade tinyint(1)  not null  comment '行政级别' ,
		full_spell varchar(20)    comment '英文/全拼' ,
		simple_spell varchar(20)    comment '简拼' ,
		sort_no bigint(20) default 0 not null  comment '排序值' ,
		status tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 启用' ,
		PRIMARY KEY (canton_code)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='地区表';
`
