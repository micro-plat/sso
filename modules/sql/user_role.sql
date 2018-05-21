create table sso_user_role(
  id number number(20) not null,
  user_id number(20) not null,
	sys_id number(20) not null,
	role_id number(20) not null,
  enable number(1) default 1 not null
  );

comment on table sso_user_role is '用户角色信息';
comment on column sso_user_role.id is '编号';
comment on column sso_user_role.user_id is '用户编号';
comment on column sso_user_role.sys_id is '系统编号';
comment on column sso_user_role.role_id is '角色编号';
comment on column sso_user_role.enable is '状态 1：启用 0:禁用';


alter table sso_user_role
add constraint pk_user_role primary key(id);

alter table sso_user_role
add constraint unq_user_role unique(user_id,sys_id,role_id);


create sequence seq_user_role_id
minvalue 100
maxvalue 999
start with 100
cache 20;
