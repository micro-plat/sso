-- 增加表字段 这个是在原来的库(sso_v3)中运行
alter table sso_user_info add source varchar(36) NOT NULL DEFAULT '' COMMENT '来源';
alter table sso_user_info add source_id varchar(120) NOT NULL DEFAULT '' COMMENT '来源id';
alter table sso_user_info add last_login_time datetime DEFAULT NULL COMMENT '最近登陆时间';