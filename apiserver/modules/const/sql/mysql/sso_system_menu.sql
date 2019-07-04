
 /*drop table sso_system_menu;*/

	create table sso_system_menu(
		id bigint primary key auto_increment   not null    comment '功能编号' ,
		name varchar(64)  not null    comment '功能名称' ,
		parent bigint  not null    comment '父级编号' ,
		sys_id bigint  not null    comment '系统编号' ,
		level_id tinyint(2)  not null    comment '等级' ,
		icon varchar(64)  DEFAULT null    comment '图标' ,
		path varchar(256)  not null    comment '地址' ,
		enable tinyint(1) DEFAULT 0 not null    comment '状态 1:正常 0:禁用' ,
		create_time datetime  DEFAULT CURRENT_TIMESTAMP not null    comment '创建时间' ,
		sortrank bigint  not null    comment '排序编号' ,
		is_open   tinyint(1) DEFAULT 0  comment '是否展开' 
					
  )COMMENT='功能信息';

 
alter table sso_system_menu add index index_system_menu_sysid(sys_id) 



