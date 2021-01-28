package mysql

const dds_area_info = `	
DROP TABLE if exists dds_area_info;
CREATE TABLE  dds_area_info (
	canton_code VARCHAR(32)  not null  comment '区域编号' ,
	chinese_name VARCHAR(128)  not null  comment '中文名称' ,
	parent_code VARCHAR(32)    comment '父级编号' ,
	grade TINYINT(2)    comment '行政级别' ,
	full_spell VARCHAR(64)    comment '英文/全拼' ,
	simple_spell VARCHAR(16)    comment '简拼' ,
	sort_id INT(11) default 0 not null  comment '排序' ,
	status SMALLINT(2)  default 0 not null  comment '状态' ,
	PRIMARY KEY (canton_code),
	KEY IDX_AREA_PARENT_CODE (parent_code)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='地区表';
`
