
 /*drop table sso_system_info;*/

	create table sso_system_info(
		id bigint  primary key auto_increment  not null    comment '编号' ,
		name varchar(32)  not null  unique  comment '系统名称' ,
		index_url varchar(64)  not null    comment '首页地址' ,
		enable tinyint(1) DEFAULT 1  not null    comment '状态 1：启用 0:禁用' ,
		login_timeout int DEFAULT 300  not null    comment '超时时长' ,
		logo varchar(128)  DEFAULT null    comment 'logp' ,	
		theme varchar(128)  DEFAULT 'bg-primary|bg-primary|bg-dark' comment '主题样式' ,
		layout varchar(128)  DEFAULT 'app-header-fixed app-aside-fixed' comment '页面布局样式' ,
		ident varchar(16) not null    comment '唯一标识' ,
		login_url varchar(64)  DEFAULT null    comment '登录地址' ,
		wechat_status tinyint(2)  DEFAULT 1 not null    comment '微信功能状态 1:开启,0:关闭' ,
		secret varchar(32)      comment '签名密钥' 		
  ) AUTO_INCREMENT = 700,COMMENT='系统信息表';

alter table sso_system_info add index index_sso_system_info_name_ident(name,ident);

