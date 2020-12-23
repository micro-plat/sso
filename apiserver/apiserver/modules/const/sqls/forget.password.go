package sqls

//ForgetPassword .
const ForgetPassword = `
update sso_user_info
set 
password =@password
where user_name=@user_name and source=@source
`
