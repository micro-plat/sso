package sql

//QueryMenuInfo 获取用户菜单信息
const QueryMenuInfo = `select user_id,user_name,status from sso_user_info where user_name=@user_name and password=@password and rownum<=1`
