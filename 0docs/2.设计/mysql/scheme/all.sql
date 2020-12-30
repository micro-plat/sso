

	CREATE TABLE  sso_data_permission (
		id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		sys_id bigint(20)  not null  comment '系统编号' ,
		ident varchar(32)  not null  comment '系统标识' ,
		name varchar(128)  not null  comment '名称' ,
		table_name varchar(128)  not null  comment '表名' ,
		operate_action varchar(64)  not null  comment '操作动作' ,
		rules text    comment '规则json' ,
		remark varchar(256)  not null  comment '说明' ,
		status tinyint(4) default 0 not null  comment '状态 0: 启用 1: 禁用' ,
		PRIMARY KEY (id),
		UNIQUE KEY unq_sso_dataprem_identn (sys_id,name)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='数据权限规则表';


	CREATE TABLE  dds_dictionary_info (
		id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		name varchar(64)  not null  comment '名称' ,
		value varchar(32)  not null  comment '值' ,
		type varchar(32)  not null  comment '类型' ,
		sort_no bigint(20) default 0 not null  comment '排序值' ,
		status tinyint(1)  not null  comment '状态 1: 禁用 0: 启用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='字典表';


	CREATE TABLE  dds_area_info (
		canton_code varchar(32)  not null  comment '区域编号' ,
		chinese_name varchar(128)  not null  comment '中文名称' ,
		parent_code varchar(32)    comment '父级编号' ,
		grade tinyint(1)  not null  comment '行政级别' ,
		full_spell varchar(20)    comment '英文/全拼' ,
		simple_spell varchar(20)    comment '简拼' ,
		sort_no bigint(20) default 0 not null  comment '排序值' ,
		status tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 启用' ,
		PRIMARY KEY (canton_code),
		KEY key_dds_area_info_parent_code (parent_code)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='地区表';


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


	CREATE TABLE  sso_user_info (
		user_id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		full_name varchar(32)  not null  comment '用户全名' ,
		user_name varchar(64)  not null  comment '用户名' ,
		password varchar(32)  not null  comment '密码' ,
		email varchar(32)    comment 'email' ,
		status tinyint(1) default 1 not null  comment '状态 0: 正常 1: 锁定 2: 禁用' ,
		mobile varchar(12)  not null  comment '电话号码' ,
		wx_openid varchar(64)    comment '微信openid' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		changepwd_times bigint(20) default 0 not null  comment '密码修改次数' ,
		ext_params varchar(1024)    comment '扩展参数' ,
		last_login_time datetime    comment '最后登录时间' ,
		source_id varchar(128) default '0' not null  comment '来源id' ,
		source varchar(36) default '''' not null  comment '来源' ,
		PRIMARY KEY (user_id),
		UNIQUE KEY unq_sso_user_info_user_name (user_name),
		
		KEY idx_sso_userinfo_source (source_id,source)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户表';


	CREATE TABLE  sso_role_menu (
		id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		sys_id bigint(20)  not null  comment '系统id' ,
		role_id bigint(20) default 0 not null  comment '角色id' ,
		menu_id bigint(20) default 0 not null  comment '菜单id' ,
		enable tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		sortrank bigint(20) default 0 not null  comment '排序编号' ,
		PRIMARY KEY (id),
		UNIQUE KEY unq_sso_rolem_srmenu (menu_id,role_id,sys_id)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表';


	CREATE TABLE  sso_user_role (
		id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		user_id bigint(20)  not null  comment '用户编号' ,
		sys_id bigint(20)  not null  comment '系统编号' ,
		role_id bigint(20)  not null  comment '角色编号' ,
		enable tinyint(1) default 1 not null  comment '状态 0: 启用 1: 禁用' ,
		PRIMARY KEY (id),
		UNIQUE KEY unq_sso_userrole_usr (user_id,sys_id,role_id)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户角色关联表';


	CREATE TABLE  sso_operate_log (
		id bigint(20)  not null AUTO_INCREMENT comment 'id' ,
		type tinyint(2)  not null  comment '类型, 10. 登录操作 20. 系统数据操作  30. 角色数据操作  40. 菜单数据操作  50. 用户数据操作' ,
		sys_id bigint(20)  not null  comment '系统编号' ,
		user_id bigint(20)  not null  comment '操作人id' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		content varchar(512)  not null  comment '内容I' ,
		PRIMARY KEY (id),
		KEY key_sso_operate_log_user_id (user_id)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='操作日志';


	CREATE TABLE  sso_system_menu (
		id bigint(20)  not null AUTO_INCREMENT comment '功能编号' ,
		name varchar(64)  not null  comment '功能名称' ,
		parent bigint(20)  not null  comment '父级编号' ,
		sys_id bigint(20)  not null  comment '系统编号' ,
		level_id tinyint(2)  not null  comment '等级' ,
		icon varchar(64)    comment '图标' ,
		path varchar(256)  not null  comment '地址' ,
		enable tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		sortrank bigint(20)  not null  comment '排序编号' ,
		is_open tinyint(1) default 0   comment '是否展开' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='功能表(菜单数据)';


	CREATE TABLE  sso_role_info (
		role_id bigint(20)  not null AUTO_INCREMENT comment '角色id' ,
		name varchar(64)  not null  comment '角色名称' ,
		status tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (role_id),
		UNIQUE KEY unq_sso_role_info_name (name)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表';


	CREATE TABLE  sso_role_datapermission (
		id bigint(20)  not null AUTO_INCREMENT comment '功能编号' ,
		sys_id bigint(20)  not null  comment '系统编号' ,
		role_id bigint(20)  not null  comment '角色编号' ,
		permission_config_id bigint(20)  not null  comment '规则id' ,
		status tinyint(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (id),
		KEY idx_sso_dataprem_sysrole (sys_id,role_id)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色与规则关联信息表';
