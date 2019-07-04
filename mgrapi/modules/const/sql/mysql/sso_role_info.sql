
 /*drop table sso_role_info;*/

	create table sso_role_info(
		`role_id` bigint primary key auto_increment  not null    comment '角色编号' ,
		`name` varchar(64) not null unique   comment '角色名称' ,
		`status` tinyint(1) DEFAULT 0 not null    comment '状态 0:正常 2:禁用' ,
		`create_time` datetime DEFAULT CURRENT_TIMESTAMP not null comment '创建时间' 
				
  )COMMENT='角色信息';

  alter table sso_role_info add index index_role_info_name(name)