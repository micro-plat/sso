
 /*drop table sso_data_permission;*/

    /*数据权限数据(来源于其他子系统)*/
	create table sso_data_permission(
		id bigint primary key auto_increment   not null    comment '功能编号' ,
        sys_id bigint  not null    comment '系统编号' ,
        ident varchar(32) not null DEFAULT ''  comment '系统标识',
		name varchar(128)  not null    comment '名称',
		type varchar(64)  not null    comment '业务类型',
		value varchar(64)  not null    comment '业务值',
		remark varchar(256)  not null    comment '说明'		
  ) COMMENT='数据权限数据(来源于其他子系统)';

 
 alter table sso_data_permission add unique(sys_id,type,value);
alter table sso_data_permission add index index_data_permission_ident_type(ident,type);