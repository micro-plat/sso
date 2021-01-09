package scheme

const sso_operate_log = `
drop table if exists sso_operate_log;

CREATE TABLE  sso_operate_log (
	id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
	type tinyint(2)  not null  comment '类型, 10. 登录操作 20. 系统数据操作  30. 角色数据操作  40. 菜单数据操作  50. 用户数据操作' ,
	sys_id bigint(20)  not null  comment '系统编号' ,
	user_id bigint(20)  not null  comment '操作人id' ,
	create_time datetime default current_timestamp not null  comment '创建时间' ,
	content varchar(512)  not null  comment '内容I' ,
	PRIMARY KEY (id),
	KEY key_sso_operate_log_user_id (user_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='操作日志';
`
