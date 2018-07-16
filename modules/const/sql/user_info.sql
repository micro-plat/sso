create table sso_user_info(
	user_id number(20) not null,
  user_name varchar2(64) not null,
	password  varchar2(32) not null,
	email  varchar2(32) not null,
	status number(1) default 1 not null,
  mobile number(11) not null,
  wx_openid  varchar2(64) not null,
  create_time date default sysdate not null
  );

comment on table sso_user_info is '用户信息';
comment on column sso_user_info.user_id is '用户编号';
comment on column sso_user_info.user_name is '用户名';
comment on column sso_user_info.password is '密码';
comment on column sso_user_info.status is '状态 0:正常 1:锁定 2:禁用';
comment on column sso_user_info.email is '邮箱地址';
comment on column sso_user_info.mobile is '手机号';
comment on column sso_user_info.wx_openid is '微信openid';
comment on column sso_user_info.create_time is '创建时间';


alter table sso_user_info
add constraint pk_user_info primary key(user_id);

alter table sso_user_info
add constraint unq_user_info unique(user_name);


create sequence seq_user_info_id
minvalue 10000
maxvalue 99999
start with 10000
cache 20;
