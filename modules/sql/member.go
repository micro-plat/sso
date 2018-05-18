package sql

//QueryUserInfo 查询用户信息
const QueryUserInfo = `select user_id,user_name,status from sso_user_info where user_name=@user_name and password=@password and rownum<=1`

//QueryRoles 查询系统角色列表
const QueryRoles = `select r.role_id,r.sys_id,s.index_url from sso_user_role r inner join sso_system_info s
 on r.sys_id=s.id
 where r.user_id=@user_id and r.sys_id@sys_id and r.enable=1 and s.enable=1`
