
 /*drop table sso_user_popular;*/

	create table sso_user_popular(
		id bigint primary key auto_increment   not null    comment '编号' ,
		user_id bigint  not null    comment '用户编号',
		sys_id bigint  not null    comment '系统编号' ,
		parent_id bigint  not null    comment '父级编号' ,
		menu_id bigint  not null    comment '菜单编号' ,
		used_cnt bigint DEFAULT 1 not null    comment '使用频率' ,
		create_time datetime DEFAULT CURRENT_TIMESTAMP not null    comment '添加时间' 
				
  )COMMENT='用户常用功能表';

alter table sso_user_popular add unique(user_id,menu_id)
alter table sso_user_popular add index index_user_popular_sysid_user_id(sys_id, user_id) 
 




