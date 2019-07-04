
 /*drop table sso_operate_log;*/

	create table sso_operate_log(
		id bigint primary key auto_increment   not null    comment '编号' ,
		type tinyint(2)  not null    comment '类型:10=登录操作,20=系统数据操作,30=角色数据操作,40=菜单数据操作,50=用户数据操作' ,
		sys_id bigint  not null    comment '系统编号(用户当时发起操作的系统编号)' ,
		user_id bigint  not null    comment '操作人userid' ,
		create_time datetime DEFAULT CURRENT_TIMESTAMP  not null    comment '创建时间' ,
		content varchar(512)  DEFAULT '' not null    comment '内容' 
				
  )COMMENT='操作日志';

  alter table sso_operate_log add index index_operate_log_userid(user_id)
 




