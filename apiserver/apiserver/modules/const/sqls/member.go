package sqls

//QueryUserByLogin 获取用户登录信息
const QueryUserByLogin = `
select 
	user_id,user_name,password,status,wx_openid,ext_params 
from 
	sso_user_info 
where 
	user_name=@user_name`

//QueryUserInfoByUID 查询用户信息
const QueryUserInfoByUID = `
select 
	u.user_id,u.user_name,u.mobile,u.wx_openid,u.status
from sso_user_info u
where u.user_id=@user_id 
limit 1
`

//QueryUserSystem 查询用户可用的子系统
const QueryUserSystem = `
select
	sys.id,
	sys.name,
	sys.index_url,
	sys.logo
from sso_system_info sys
inner join sso_user_role ur on ur.sys_id = sys.id
inner join sso_role_info role on role.role_id = ur.role_id
where role.status = 0 AND
  ur.enable=1 and 
  sys.index_url is not null and
  sys.index_url <> '' and
  sys.ident <> @ident and
  sys.enable=1 and
  ur.user_id = @user_id; 
`

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
from sso_user_role r
inner join sso_system_info s on r.sys_id=s.id
inner join sso_role_info i on i.role_id=r.role_id 
where 
	r.user_id=@user_id 
	and s.ident=@ident 
	and r.enable=1	  
	and s.enable=1	
	and i.status=0`

//QueryUserByUserName 根据用户名获取用户信息
const QueryUserByUserName = `
select 
	user_id,
	user_name,
	status,
	wx_openid,
	ext_params 
from sso_user_info
where user_name=@user_name 
limit 1
`

//QuerySysAuth .
const QuerySysAuth = `
select 
	t.* from sso_user_role t 
inner join 
	sso_user_info i on i.user_id=t.user_id 
where 
	t.sys_id=@sys_id 
	and t.user_id=@user_id`

const QueryAllUserInfo = `
select 
	user_id,
	user_name,
	full_name,
	status
from sso_user_info
where (@source = '' or @source = source) and
	  (@source_id = 0 or @source_id = source_id)
`
