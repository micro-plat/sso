
 /*drop table sso_user_info;*/

	create table sso_user_info(
		user_id bigint primary key auto_increment   not null    comment '用户编号' ,
		user_name varchar(64)  not null unique  comment '用户名' ,
		password varchar(32)  not null    comment '密码' ,
		email varchar(32)      comment '邮箱地址' ,
		status tinyint(1) DEFAULT 1  not null    comment '状态 0:正常 1:锁定 2:禁用' ,
		mobile char(12)  not null    comment '手机号,座机号可能12位' ,
		wx_openid varchar(64)      comment '微信openid' ,
		create_time datetime DEFAULT CURRENT_TIMESTAMP  not null  comment '创建时间' ,
		changepwd_times bigint  DEFAULT 0  comment '密码修改次数' ,
		ext_params varchar(1024)      comment '扩展字段' 
				
  ) COMMENT='用户信息表';

 alter table sso_user_info add index index_user_info_username(user_name);




