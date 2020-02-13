
 /*drop table sso_data_permission;*/

    /*数据权限数据(来源于其他子系统)*/
	create table sso_data_permission(
		id bigint primary key auto_increment   not null    comment '功能编号' ,
        sys_id bigint  not null    comment '系统编号' ,
        ident varchar(32) not null DEFAULT ''  comment '系统标识',
		name varchar(128)  not null    comment '名称',
		table_name VARCHAR(128) not null comment '表名',
		operate_action varchar(64)  not null    comment '操作动作',
		rules VARCHAR(1000) not null comment '权限规则(json)',
		remark varchar(256)  not null    comment '说明'		
  ) COMMENT='数据权限数据(来源于其他子系统)';

 
alter table sso_data_permission add unique(sys_id, table_name, operate_action);