package sql

//QueryUserByLogin 获取用户登录信息
<<<<<<< HEAD:apiserver/modules/const/sql/member.go
const QueryUserByLogin = `select user_id,user_name,password,status,wx_openid,ext_params from sso_user_info where user_name=@user_name and rownum<=1`
=======
const QueryUserByLogin = `
select 
	user_id,user_name,password,status,wx_openid,ext_params 
from 
	sso_user_info 
where 
	user_name=@user_name`
>>>>>>> 750f5c63baeb3b4a71bc53caecd154a8e0ed6969:flowserver/modules/const/sql/member.go

//QueryUserInfoByUID 查询用户信息
const QueryUserInfoByUID = `
select 
	u.user_id,u.user_name,u.mobile,u.wx_openid,u.status,u.login_url 
from 
	sso_user_info u
where 
	u.user_id=@user_id `

//QueryUserInfoByOpenID 查询用户信息
const QueryUserInfoByOpenID = `
select 
	u.user_id,u.user_name,u.password,u.status 
from 
	sso_user_info u
where 
	u.wx_openid=@open_id `

//QueryUserRole 查询系统角色列表
const QueryUserRole = `
select 
	r.role_id,i.name role_name,s.index_url,s.login_url,s.login_timeout,r.sys_id 
from 
	sso_user_role r
inner join 
	sso_system_info s on r.sys_id=s.id
inner join 
	sso_role_info i on i.role_id=r.role_id 
where 
	r.user_id=@user_id 
	and s.ident=@ident 
	and r.enable=1 
	and s.enable=1`

<<<<<<< HEAD:apiserver/modules/const/sql/member.go
const QuerySysAuth = `select t.* from sso_user_role t 
inner join sso_user_info i on i.user_id=t.user_id 
where t.sys_id=@sys_id and t.user_id=@user_id`

//QueryUserByUserName 根据用户名获取用户信息
const QueryUserByUserName = `
	select 
		user_id,
		user_name,
		status,
		wx_openid,
		ext_params 
	from 
		sso_user_info
	where 
		user_name=@user_name 
	and 
		rownum<=1
`
=======
//QuerySysAuth .
const QuerySysAuth = `
select 
	t.* from sso_user_role t 
inner join 
	sso_user_info i on i.user_id=t.user_id 
where 
	t.sys_id=@sys_id 
	and t.user_id=@user_id`
>>>>>>> 750f5c63baeb3b4a71bc53caecd154a8e0ed6969:flowserver/modules/const/sql/member.go
