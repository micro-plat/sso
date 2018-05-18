create table sso_menu_info(
	id number(20) not null,
  name varchar2(64) not null,
  parent number(20) not null,
  sys_id number(20) not null
  level number(2) not null
  icon varchar2(32) not null,
  path varchar2(32) not null
  status number(1) default 0 not null

  create_time date default sysdate not null
  sortrank number(20) not null
  );

comment on table sso_menu_info is '功能信息';
comment on column sso_menu_info.id is '功能编号';
comment on column sso_menu_info.name is '功能名称';
comment on column sso_menu_info.parent is '父级编号';
comment on column sso_menu_info.sys_id is '系统编号';
comment on column sso_menu_info.level is '等级';
comment on column sso_menu_info.icon is '图标';
comment on column sso_menu_info.path is '地址';
comment on column sso_menu_info.status is '状态 0:正常 2:禁用';
comment on column sso_menu_info.create_time is '创建时间';


alter table sso_menu_info
add constraint pk_menu_info primary key(id);


create sequence seq_menu_info_id
minvalue 10000
maxvalue 99999
start with 10000
cache 20;


create sequence seq_menu_info_sort_id
minvalue 1
maxvalue 9999
start with 1
increment by 10
cache 20;
