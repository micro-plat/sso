package sqls

//QueryUserByLogin 获取用户登录信息
const QueryUserByLogin = `
select 
	user_id,user_name,password,status,wx_openid,ext_params 
from 
	sso_user_info 
where 
	user_name=@user_name`

//QueryUserSystem 查询用户可用的子系统
const QueryUserSystem = `
select
	sys.id,
	sys.name,
	sys.index_url,
	sys.index_url path,
	'blank' type,
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
sso_user_infoselect 
	user_id,
	user_name,
	full_name,
	status,
    user_id value,
    full_name name,
    'userinfo' type
from sso_user_info
where (@source = '' or source =  @source) and
	  ((@source_id = '' or @source_id=0) or source_id = @source_id)
`
