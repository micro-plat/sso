drop table sso_notify_records;

create table sso_notify_records(
	id number(20) not null,
  sys_id number(10) not null,
  level_id number(1) default 1 not null,
	title  varchar2(32) not null,
  keywords varchar2(32) not null,
	content  varchar2(32) not null,
	status number(1) default 1 not null,
  create_time date default sysdate not null,
  finish_time date
  );

comment on table sso_notify_records is '报警消息表';
comment on column sso_notify_records.id is '报警编号';
comment on column sso_notify_records.sys_id is '系统编号';
comment on column sso_notify_records.level_id is '等级';
comment on column sso_notify_records.title is '标题';
comment on column sso_notify_records.keywords is '关键字';
comment on column sso_notify_records.content is '内容';
comment on column sso_notify_records.status is '状态';
comment on column sso_notify_records.create_time is '创建时间';
comment on column sso_notify_records.finish_time is '完成时间';



alter table sso_notify_records
add constraint pk_notify_records primary key(id);


drop sequence seq_notify_records;

create sequence seq_notify_records_id
minvalue 100
maxvalue 9999999999999
start with 10000
cache 20;
