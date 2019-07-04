package sql

//AddOperate .
const AddOperate = `
insert into 
	sso_operate_log(type,sys_id,user_id,content) 
values
	(@type,@sys_id,@user_id,@content)`
