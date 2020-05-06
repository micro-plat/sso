package sqls

//ForgetPassword .
const ForgetPassword = `
update sso_user_info
set 
password =@password
where source_id=@source_id and source=@source
`
