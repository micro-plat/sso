create table sso_system_info(
  id number number(20) not null,
  name varchar2(32) not null,
	index_url varchar2(64) not null,
  enable number(1) default 1 not null
  );

comment on table sso_system_info is '系统信息表';
comment on column sso_system_info.id is '编号';
comment on column sso_system_info.name is '系统名称';
comment on column sso_system_info.index_url is '首页地址';
comment on column sso_system_info.enable is '状态 1：启用 0:禁用';


alter table sso_system_info
add constraint pk_user_info primary key(id);

alter table sso_system_info
add constraint pk_user_info unique(name);


create sequence seq_system_info_id
minvalue 100
maxvalue 999
start with 100
cache 20;
