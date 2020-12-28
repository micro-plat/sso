
	CREATE TABLE  sso_system_info (
		id number(20)  not null  comment 'id' ,
		name varchar2(32)  not null  comment '系统名称' ,
		index_url varchar2(64)  not null  comment '首页地址' ,
		enable number(1) default 1 not null  comment '状态 1：启用 0: 禁用' ,
		login_timeout number(6) default 300 not null  comment '超时时长' ,
		logo varchar2(128)  not null  comment 'logo' ,
		theme varchar2(128) default 'bg-parimary' not null  comment '主题样式' ,
		layout varchar2(128) default 'app-header-fixed app-aside-fixed' not null  comment '页面布局样式' ,
		ident varchar2(16)  not null  comment '唯一标识' ,
		login_url varchar2(64)    comment '登录地址' ,
		wechat_status number(2) default 1 not null  comment '微信功能状态 1: 开启, 0: 关闭' ,
		secret varchar2(32)  not null  comment '签名密钥' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='系统信息';
