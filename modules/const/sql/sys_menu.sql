drop table sso_system_menu;

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
comment on column sso_system_menu.enable is '状态 1:正常 0:禁用';
comment on column sso_system_menu.create_time is '创建时间';
comment on column sso_system_menu.sortrank is '排序编号';


alter table sso_system_menu
add constraint pk_system_menu primary key(id);

drop sequence seq_system_menu_id;

create sequence seq_system_menu_id
minvalue 10000
maxvalue 99999
start with 10000
cache 20;

drop sequence seq_system_menu_sort_id;

create sequence seq_system_menu_sort_id
minvalue 1
maxvalue 9999
start with 1
increment by 10
cache 20;



insert into sso_system_menu (ID, NAME, PARENT, SYS_ID, LEVEL_ID, ICON, PATH, ENABLE, CREATE_TIME, SORTRANK, IS_OPEN)
values (1, '用户权限', 0, 1, 1, '-', '-', 1, to_date('20180724', 'yyyymmdd'), 2, 1);

insert into sso_system_menu (ID, NAME, PARENT, SYS_ID, LEVEL_ID, ICON, PATH, ENABLE, CREATE_TIME, SORTRANK, IS_OPEN)
values (11, '用户角色', 1, 1, 2, 'fa fa-user-circle-o text-info', '-', 1, to_date('20180724', 'yyyymmdd'), 1, 1);

insert into sso_system_menu (ID, NAME, PARENT, SYS_ID, LEVEL_ID, ICON, PATH, ENABLE, CREATE_TIME, SORTRANK, IS_OPEN)
values (111, '用户管理', 11, 1, 3, 'fa fa-user-o text-info', '/user/index', 1, to_date('20180724', 'yyyymmdd'), 1, 1);

insert into sso_system_menu (ID, NAME, PARENT, SYS_ID, LEVEL_ID, ICON, PATH, ENABLE, CREATE_TIME, SORTRANK, IS_OPEN)
values (112, '角色权限', 11, 1, 3, 'fa fa-commenting-o text-info', '/user/role', 1, to_date('20180724', 'yyyymmdd'), 2, 1);

insert into sso_system_menu (ID, NAME, PARENT, SYS_ID, LEVEL_ID, ICON, PATH, ENABLE, CREATE_TIME, SORTRANK, IS_OPEN)
values (12, '系统功能', 1, 1, 2, 'fa fa-tasks text-info-lter', '-', 1, to_date('20180724', 'yyyymmdd'), 2, 1);

insert into sso_system_menu (ID, NAME, PARENT, SYS_ID, LEVEL_ID, ICON, PATH, ENABLE, CREATE_TIME, SORTRANK, IS_OPEN)
values (121, '系统管理', 12,1, 3, 'fa fa-folder text-success', '/sys/index', 1, to_date('20180724', 'yyyymmdd'), 1, 1);
