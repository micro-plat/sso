
 /*drop table sso_role_menu;*/

	create table sso_role_menu(
		id bigint primary key auto_increment   not null    comment '功能编号' ,
		sys_id bigint  not null    comment '系统编号' ,
		role_id bigint  not null    comment '角色编号' ,
		menu_id bigint  not null    comment '菜单编号' ,
		enable tinyint(1) DEFAULT 0 not null    comment '状态 1:正常 0:禁用' ,
		create_time datetime DEFAULT CURRENT_TIMESTAMP not null    comment '创建时间' ,
		sortrank bigint  not null    comment '排序编号'
				
  )COMMENT='功能信息';
    /*auto_increment=10000*/

alter table sso_role_menu add unique(sys_id,role_id,menu_id);
alter table sso_role_menu add index index_role_menu_roleid_menuid(role_id,menu_id)
 




