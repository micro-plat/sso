create table sso_role_menu(
	id number(20) not null,
  sys_id number(20) not null,
  role_id number(20) not null,
  menu_id number(20) not null,
  status number(1) default 0 not null,
  create_time date default sysdate not null,
  sortrank number(20) not null
  );

comment on table sso_role_menu is '功能信息';
comment on column sso_role_menu.id is '功能编号';
comment on column sso_role_menu.sys_id is '系统编号';
comment on column sso_role_menu.role_id is '角色编号';
comment on column sso_role_menu.menu_id is '菜单编号';
comment on column sso_role_menu.status is '状态 0:正常 2:禁用';
comment on column sso_role_menu.create_time is '创建时间';
comment on column sso_role_menu.sortrank is '排序编号';


alter table sso_role_menu
add constraint pk_role_menu primary key(id);

alter table sso_role_menu
add constraint unq_role_menu unique(sys_id,role_id,menu_id);


create sequence seq_role_menu_id
minvalue 10000
maxvalue 99999
start with 10000
cache 20;


create sequence seq_role_menu_sort_id
minvalue 1
maxvalue 9999
start with 1
increment by 10
cache 20;
