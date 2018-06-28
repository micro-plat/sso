create table sso_user_popular(
  id number(20) not null,
  user_id number(20) not null,
	sys_id number(20) not null,
  parent_id number(20) not null,
	menu_id number(20) not null,
  used_cnt number(1) default 1 not null,
  create_time date default sysdate not null
  );

comment on table sso_user_popular is '用户常用功能表';
comment on column sso_user_popular.id is '编号';
comment on column sso_user_popular.user_id is '用户编号';
comment on column sso_user_popular.sys_id is '系统编号';
comment on column sso_user_popular.parent_id is '父级编号';
comment on column sso_user_popular.menu_id is '菜单编号';
comment on column sso_user_popular.used_cnt is '使用频率';
comment on column sso_user_popular.create_time is '添加时间';

alter table sso_user_popular
add constraint pk_user_popular primary key(id);

alter table sso_user_popular
add constraint unq_user_popular unique(user_id,menu_id);


create sequence seq_user_popular_id
minvalue 100
maxvalue 999
start with 100
cache 20;
