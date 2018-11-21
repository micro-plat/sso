package sql

const AddOperate = `insert into sso_operate_log(id,type,sys_id,user_id,content) 
values(seq_operate_log_id.nextval,@type,@sys_id,@user_id,@content)`
