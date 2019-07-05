
 /*drop table sso_notify_subscribe;*/

	create table sso_notify_subscribe(
		id bigint  primary key auto_increment  not null    comment '报警编号' ,
		user_id int  not null    comment '用户编号' ,
		sys_id int  not null    comment '系统编号' ,
		level_id tinyint(1)  not null    comment '等级' ,
		keywords varchar(32)  not null    comment '关键字' ,
		status tinyint(1) DEFAULT 1 not null    comment '状态' ,
		create_time datetime DEFAULT CURRENT_TIMESTAMP  not null    comment '创建时间' 
				
  ) AUTO_INCREMENT = 11000, COMMENT='报警消息表';

 




