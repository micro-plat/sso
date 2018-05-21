create table sso_role_menu_info(
	id number(20) not null,
  sys_id number(20) not null
  role_id number(20) not null,
  menu_id number(20) not null,
  status number(1) default 0 not null
  create_time date default sysdate not null
  sortrank number(20) not null
  );

comment on table sso_role_menu_info is '功能信息';
comment on column sso_role_menu_info.id is '功能编号';
comment on column sso_role_menu_info.sys_id is '系统编号';
comment on column sso_role_menu_info.role_id is '角色编号';
comment on column sso_role_menu_info.menu_id is '菜单编号';
comment on column sso_role_menu_info.status is '状态 0:正常 2:禁用';
comment on column sso_role_menu_info.create_time is '创建时间';
comment on column sso_menu_info.sortrank is '排序编号';


alter table sso_role_menu_info
add constraint pk_menu_info primary key(id);

alter table sso_user_info
add constraint pk_user_info unique(sys_id,role_id,menu_id);


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
