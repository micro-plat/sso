create table sso_role_info(
	role_id number(20) not null,
  name varchar2(64) not null,
  status number(1) not null,  
  create_time date default sysdate not null
  );

comment on table sso_role_info is '角色信息';
comment on column sso_role_info.role_id is '角色编号';
comment on column sso_role_info.name is '角色名称';
comment on column sso_role_info.status is '状态 0:正常 1:锁定 2:禁用';
comment on column sso_role_info.create_time is '创建时间';


alter table sso_role_info
add constraint pk_role_info primary key(role_id);

alter table sso_role_info
add constraint unq_role_info unique(name);


create sequence seq_role_info_id
minvalue 10000
maxvalue 99999
start with 10000
cache 20;
