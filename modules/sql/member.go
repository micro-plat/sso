package sql

//QueryUserByLogin 获取用户登录信息
const QueryUserByLogin = `select user_id,user_name,password,status from sso_user_info where user_name=@user_name and rownum<=1`

//QueryUserInfoByUID 查询用户信息
const QueryUserInfoByUID = `select u.user_id,u.user_name,u.mobile,u.wx_openid,u.status from sso_user_info u 
where u.user_id=@user_id `

//QueryUserRole 查询系统角色列表
const QueryUserRole = `select r.role_id,s.index_url,s.login_timeout from sso_user_role r inner join sso_system_info s
 on r.sys_id=s.id
 where r.user_id=@user_id and r.sys_id=@sys_id and r.enable=1 and s.enable=1`
