
 /*drop table sso_notify_records;*/

	create table sso_notify_records(
		id bigint primary key auto_increment   not null    comment '报警编号' ,
		sys_id bigint  not null    comment '系统编号' ,
		level_id bigint  not null    comment '等级' ,
		title varchar(32)  not null    comment '标题' ,
		content varchar(32)  not null    comment '内容' ,
		keywords varchar(32)  not null    comment '关键字' ,
		status  tinyint(1) DEFAULT 1 not null    comment '状态' ,
		create_time datetime DEFAULT CURRENT_TIMESTAMP not null    comment '创建时间' ,
		finish_time datetime      comment '完成时间' 
  )COMMENT='报警消息表';

 




