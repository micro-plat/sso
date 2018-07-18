create table sso_system_menu(
	id number(20) not null,
  name varchar2(64) not null,
  parent number(20) not null,
  sys_id number(20) not null,
  level_id number(2) not null,
  icon varchar2(32) not null,
  path varchar2(32) not null,
  is_open number(1) not null,
  enable number(1) default 0 not null,
  create_time date default sysdate not null,
  sortrank number(20) not null
  );

comment on table sso_system_menu is '功能信息';
comment on column sso_system_menu.id is '功能编号';
comment on column sso_system_menu.name is '功能名称';
comment on column sso_system_menu.parent is '父级编号';
comment on column sso_system_menu.sys_id is '系统编号';
comment on column sso_system_menu.level_id is '等级';
comment on column sso_system_menu.icon is '图标';
comment on column sso_system_menu.path is '地址';
comment on column sso_system_menu.is_open is '是否展开';
comment on column sso_system_menu.status is '状态 1:正常 0:禁用';
comment on column sso_system_menu.create_time is '创建时间';
comment on column sso_system_menu.sortrank is '排序编号';


alter table sso_system_menu
add constraint pk_system_menu primary key(id);


create sequence seq_system_menu_id
minvalue 10000
maxvalue 99999
start with 10000
cache 20;


create sequence seq_system_menu_sort_id
minvalue 1
maxvalue 9999
start with 1
increment by 10
cache 20;
