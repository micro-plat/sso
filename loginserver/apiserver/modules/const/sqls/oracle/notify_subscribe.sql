drop table sso_notify_subscribe;

create table sso_notify_subscribe(
	id number(20) not null,
  user_id number(10) not null,
  sys_id number(10) not null,
  level_id number(1) default 1 not null,
  keywords varchar2(32) not null,
	status number(1) default 1 not null,
  create_time date default sysdate not null
  );

comment on table sso_notify_subscribe is '报警消息表';
comment on column sso_notify_subscribe.user_id is '用户编号';
comment on column sso_notify_subscribe.id is '报警编号';
comment on column sso_notify_subscribe.sys_id is '系统编号';
comment on column sso_notify_subscribe.level_id is '等级';
comment on column sso_notify_subscribe.keywords is '关键字';
comment on column sso_notify_subscribe.status is '状态';
comment on column sso_notify_subscribe.create_time is '创建时间';



alter table sso_notify_subscribe
add constraint pk_notify_subscribe primary key(id);


drop sequence seq_notify_subscribe;

create sequence seq_notify_subscribe_id
minvalue 100
maxvalue 9999999999999
start with 10000
cache 20;
