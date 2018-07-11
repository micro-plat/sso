create table sso_system_info(
  id number(20) not null,
  name varchar2(32) not null,
	index_url varchar2(64) not null,
  logo varchar2(64) not null,
  login_timeout number(20) default 86400 not null,
  theme varchar2(128) default 'bg-primary|bg-primary|bg-dark' not null,
  layout varchar2(128) default 'app-header-fixed app-aside-fixed' not null,  
  enable number(1) default 1 not null
  );

comment on table sso_system_info is '系统信息表';
comment on column sso_system_info.id is '编号';
comment on column sso_system_info.name is '系统名称';
comment on column sso_system_info.index_url is '首页地址';
comment on column sso_system_info.logo is 'logo地址';
comment on column sso_system_info.login_timeout is '登录超时时长';
comment on column sso_system_info.theme is '主题样式';
comment on column sso_system_info.layout is '页面布局样式';
comment on column sso_system_info.enable is '状态 1：启用 0:禁用';


alter table sso_system_info
add constraint pk_system_info primary key(id);

alter table sso_system_info
add constraint unq_system_info unique(name);


create sequence seq_system_info_id
minvalue 100
maxvalue 999
start with 100
cache 20;
