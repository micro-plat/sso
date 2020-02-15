
 /*drop table sso_data_permission;*/

    /*数据权限规则配置*/
	create table sso_data_permission(
		id bigint primary key auto_increment   not null    comment '功能编号' ,
        sys_id bigint  not null    comment '系统编号' ,
        ident varchar(32) not null DEFAULT ''  comment '系统标识',
		name varchar(128)  not null    comment '规则名称',
		rules VARCHAR(1000) not null comment '权限规则(json)',
		remark varchar(256)  not null    comment '说明'		
  ) COMMENT='数据权限规则配置';

 
alter table sso_data_permission add unique(sys_id, name);