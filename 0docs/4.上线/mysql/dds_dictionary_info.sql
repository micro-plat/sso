create table dds_dictionary_info(
	id bigint  not null PRIMARY KEY AUTO_INCREMENT  comment '编号',
	name varchar(64)  not null    comment '名称' ,
	value varchar(32) not null comment '值',
	type varchar(32)  not null   comment '类型',
	sort_no bigint not null default 0 comment '排序值',
	status tinyint(1) not null comment '状态(0:启用,1:禁用)'			
)COMMENT='字典表';

ALTER TABLE dds_dictionary_info ADD INDEX idx_dictionary_info_type (type);


insert into dds_dictionary_info(name, value, type, status)VALUES('新增','新增', 'operate_action', 0);
insert into dds_dictionary_info(name, value, type, status)VALUES('修改','修改', 'operate_action', 0);
insert into dds_dictionary_info(name, value, type, status)VALUES('启用','启用', 'operate_action', 0);
insert into dds_dictionary_info(name, value, type, status)VALUES('禁用','禁用', 'operate_action', 0);
insert into dds_dictionary_info(name, value, type, status)VALUES('锁定','锁定', 'operate_action', 0);