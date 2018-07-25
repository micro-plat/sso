drop table sso_role_menu;

create table sso_role_menu(
	id number(20) not null,
  sys_id number(20) not null,
  role_id number(20) not null,
  menu_id number(20) not null,
  enable number(1) default 0 not null,
  create_time date default sysdate not null,
  sortrank number(20) not null
  );

comment on table sso_role_menu is '功能信息';
comment on column sso_role_menu.id is '功能编号';
comment on column sso_role_menu.sys_id is '系统编号';
comment on column sso_role_menu.role_id is '角色编号';
comment on column sso_role_menu.menu_id is '菜单编号';
comment on column sso_role_menu.enable is '状态 1:正常 0:禁用';
comment on column sso_role_menu.create_time is '创建时间';
comment on column sso_role_menu.sortrank is '排序编号';


alter table sso_role_menu
add constraint pk_role_menu primary key(id);

alter table sso_role_menu
add constraint unq_role_menu unique(sys_id,role_id,menu_id);

drop sequence seq_role_menu_id;

create sequence seq_role_menu_id
minvalue 10000
maxvalue 99999
start with 10000
cache 20;

drop sequence seq_role_menu_sort_id;

create sequence seq_role_menu_sort_id
minvalue 1
maxvalue 9999
start with 1
increment by 10
cache 20;



insert into sso_role_menu (ID, SYS_ID, ROLE_ID, MENU_ID, ENABLE, CREATE_TIME, SORTRANK)
values (1, 1, 1, 1, 1, to_date('20180724', 'yyyymmdd'), 1);

insert into sso_role_menu (ID, SYS_ID, ROLE_ID, MENU_ID, ENABLE, CREATE_TIME, SORTRANK)
values (2, 1, 1, 11, 1, to_date('20180724', 'yyyymmdd'), 2);

insert into sso_role_menu (ID, SYS_ID, ROLE_ID, MENU_ID, ENABLE, CREATE_TIME, SORTRANK)
values (3, 1, 1, 12, 1, to_date('20180724', 'yyyymmdd'), 5);

insert into sso_role_menu (ID, SYS_ID, ROLE_ID, MENU_ID, ENABLE, CREATE_TIME, SORTRANK)
values (4, 1, 1, 111, 1, to_date('20180724', 'yyyymmdd'), 3);

insert into sso_role_menu (ID, SYS_ID, ROLE_ID, MENU_ID, ENABLE, CREATE_TIME, SORTRANK)
values (5, 1, 1, 112, 1, to_date('20180724', 'yyyymmdd'), 4);

insert into sso_role_menu (ID, SYS_ID, ROLE_ID, MENU_ID, ENABLE, CREATE_TIME, SORTRANK)
values (6, 1, 1, 121, 1, to_date('20180724', 'yyyymmdd'), 6);
