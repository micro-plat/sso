
	CREATE TABLE  sso_system_info (
		id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		name varchar(32)  not null  comment '系统名称' ,
		index_url varchar(64)    comment '首页地址' ,
		enable tinyint(1) default 1 not null  comment '状态 1：启用 0: 禁用' ,
		login_timeout int(11) default 300 not null  comment '超时时长' ,
		logo varchar(128)    comment 'logo' ,
		theme varchar(128) default 'bg-parimary'   comment '主题样式' ,
		layout varchar(128) default 'app-header-fixed app-aside-fixed'   comment '页面布局样式' ,
		ident varchar(32)  not null  comment '唯一标识' ,
		login_url varchar(64)    comment '登录地址' ,
		wechat_status tinyint(2) default 1 not null  comment '微信功能状态 1: 开启, 0: 关闭' ,
		secret varchar(32)    comment '签名密钥' ,
		PRIMARY KEY (id),
		UNIQUE KEY unq_sso_system_info_name (name),
		
		KEY idx_sso_sysinfo_urlident (name,ident)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='系统信息';
