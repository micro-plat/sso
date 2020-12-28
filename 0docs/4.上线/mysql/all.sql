

	CREATE TABLE  sso_system_menu (
		id number(20)  not null  comment '功能编号' ,
		name varchar2(64)  not null  comment '功能名称' ,
		parent number(20)  not null  comment '父级编号' ,
		sys_id number(20)  not null  comment '系统编号' ,
		level_id number(2)  not null  comment '等级' ,
		icon varchar2(64)    comment '图标' ,
		path varchar2(256)  not null  comment '地址' ,
		enable number(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		sortrank number(20)  not null  comment '排序编号' ,
		is_open number(1) default 0   comment '是否展开' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='功能表(菜单数据)';


	CREATE TABLE  sso_user_info (
		user_id number(10)  not null  comment 'id' ,
		full_name varchar2(10)  not null  comment '用户全名' ,
		user_name varchar2(64)  not null  comment '用户名' ,
		password varchar2(32)  not null  comment '密码' ,
		email varchar2(32)    comment 'email' ,
		status number(1) default 1 not null  comment '状态 0: 正常 1: 锁定 2: 禁用' ,
		mobile varchar2(12)  not null  comment '电话号码' ,
		wx_openid varchar2(64)    comment '微信openid' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		changepwd_times number(2) default 0 not null  comment '密码修改次数' ,
		source varchar2(36)  not null  comment '来源' ,
		source_id varchar2(120)  not null  comment '来源id' ,
		ext_params varchar2(1024)    comment '扩展参数' ,
		PRIMARY KEY (user_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户表';


	CREATE TABLE  sso_data_permission (
		id number(20)  not null  comment 'id' ,
		sys_id number(20)  not null  comment '系统编号' ,
		ident varchar2(32)  not null  comment '系统标识' ,
		name varchar2(128)  not null  comment '名称' ,
		table_name varchar2(128)  not null  comment '表名' ,
		operate_action varchar2(64)  not null  comment '操作动作' ,
		rules LONGTEXT    comment '规则json' ,
		remark varchar2(256)  not null  comment '说明' ,
		status number(4) default 0 not null  comment '状态 0: 启用 1: 禁用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='数据权限规则表';


	CREATE TABLE  dds_area_info (
		canton_code varchar2(32)  not null  comment '区域编号' ,
		chinese_name varchar2(128)  not null  comment '中文名称' ,
		parent_code varchar2(32)    comment '父级编号' ,
		grade number(1)  not null  comment '行政级别' ,
		full_spell varchar2(20)    comment '英文/全拼' ,
		simple_spell varchar2(20)    comment '简拼' ,
		sort_no number(20) default 0 not null  comment '排序值' ,
		status number(1) default 0 not null  comment '状态 1: 禁用 0: 启用' ,
		PRIMARY KEY (canton_code)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='地区表';


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


	CREATE TABLE  sso_role_info (
		role_id number(20)  not null  comment '角色id' ,
		name varchar2(64)  not null  comment '角色名称' ,
		status number(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (role_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表';


	CREATE TABLE  sso_role_menu (
		id number(20)  not null  comment 'id' ,
		sys_id number(20)  not null  comment '系统id' ,
		role_id number(20) default 0 not null  comment '角色id' ,
		menu_id number(20) default 0 not null  comment '菜单id' ,
		enable number(1) default 0 not null  comment '状态 1: 禁用 0: 正常' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		sortrank number(20) default 0 not null  comment '排序编号' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色表';


	CREATE TABLE  sso_role_datapermission (
		id number(20)  not null  comment '功能编号' ,
		sys_id number(20)  not null  comment '系统编号' ,
		role_id number(20)  not null  comment '角色编号' ,
		permission_config_id number(20)  not null  comment '规则id' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='角色与规则关联信息表';


	CREATE TABLE  sso_user_role (
		id number(20)  not null  comment 'id' ,
		user_id number(20)  not null  comment '用户编号' ,
		sys_id number(20)  not null  comment '系统编号' ,
		role_id number(20)  not null  comment '角色编号' ,
		enable number(1) default 1 not null  comment '状态 0: 启用 1: 禁用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户角色关联表';


	CREATE TABLE  sso_operate_log (
		id number(20)  not null  comment 'id' ,
		type number(2)  not null  comment '类型,10.登录操作 20.系统数据操作  30.角色数据操作  40.菜单数据操作  50.用户数据操作' ,
		sys_id number(20)  not null  comment '系统编号' ,
		user_id number(20)  not null  comment '操作人id' ,
		create_time DATETIME  not null  comment '创建时间' ,
		content varchar2(512)  not null  comment '内容I' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='操作日志';


	CREATE TABLE  dds_dictionary_info (
		id number(20)  not null  comment 'id' ,
		name varchar2(64)  not null  comment '名称' ,
		value varchar2(32)  not null  comment '值' ,
		type varchar2(32)  not null  comment '类型' ,
		sort_no number(20) default 0 not null  comment '排序值' ,
		status number(1)  not null  comment '状态 1: 禁用 0: 启用' ,
		PRIMARY KEY (id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='字典表';
