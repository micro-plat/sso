package sqls

//AddOperate .
const AddOperate = `
insert into 
	sso_operate_log(type,sys_id,user_id,content) 
values
	(@type,@sys_id,@user_id,@content)`

//ValidUserNameExist 验证用户是否存在
const ValidUserNameExist = `
select t.* from sso_user_info t 
where t.user_name = @user_name
`
