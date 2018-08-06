drop table sso_notify_user;

create table sso_notify_user(
	id number(20) not null,
  user_id number(10) not null,
  sys_id number(10) not null,
  level_id number(1) default 1 not null,
	title  varchar2(32) not null,
  keywords varchar2(32) not null,
	content  varchar2(32) not null,
	status number(1) default 1 not null,
  create_time date default sysdate not null,
  finish_time date
  );

comment on table sso_notify_user is '报警消息表';
comment on column sso_notify_user.user_id is '用户编号';
comment on column sso_notify_user.id is '报警编号';
comment on column sso_notify_user.sys_id is '系统编号';
comment on column sso_notify_user.level_id is '等级';
comment on column sso_notify_user.title is '标题';
comment on column sso_notify_user.keywords is '关键字';
comment on column sso_notify_user.content is '内容';
comment on column sso_notify_user.status is '状态[1:等待 2:正在 0:成功 9:失败]';
comment on column sso_notify_user.create_time is '创建时间';
comment on column sso_notify_user.finish_time is '完成时间';



alter table sso_notify_user
add constraint pk_notify_user primary key(id);


drop sequence seq_notify_user;

create sequence seq_notify_user_id
minvalue 100
maxvalue 9999999999999
start with 10000
cache 20;
