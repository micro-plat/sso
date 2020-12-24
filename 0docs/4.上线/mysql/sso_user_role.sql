
 /*drop table sso_user_role;*/

	create table sso_user_role(
		id bigint primary key auto_increment  not null   comment '编号' ,
		user_id bigint   not null    comment '用户编号' ,
		sys_id bigint  not null    comment '系统编号' ,
		role_id bigint  not null    comment '角色编号' ,
		enable tinyint(1) not null default 1      comment '状态 1：启用 0:禁用' 
				
  )comment='用户角色信息';

    /*auto_increment=10000*/

alter table sso_user_role add unique(user_id,sys_id,role_id);
alter table sso_user_role add index index_user_role_sysid_roleid(sys_id, user_id, role_id); 