
 /*drop table sso_notify_user;*/

	create table sso_notify_user(
		id bigint primary key auto_increment  not null    comment '报警编号' ,
		user_id int  not null    comment '用户编号' ,
		sys_id int  not null    comment '系统编号' ,
		level_id tinyint(1) DEFAULT 1 not null    comment '等级' ,
		keywords varchar(32)  not null    comment '关键字' ,
		title varchar(32)  not null    comment '标题' ,
		content varchar(32)  not null    comment '内容' ,
		status tinyint(1)  not null    comment '状态[1:等待 2:正在 0:成功 9:失败]' ,
		create_time datetime DEFAULT CURRENT_TIMESTAMP  not null    comment '创建时间' ,
		scan_batch_id varchar(32)      comment '扫描批次' ,
		send_count int DEFAULT 0  not null    comment '发送次数' ,
		flow_timeout datetime      comment '流程超时时间' ,
		finish_time datetime      comment '完成时间' 
				
  ) AUTO_INCREMENT = 11000, COMMENT='报警消息表';

 




