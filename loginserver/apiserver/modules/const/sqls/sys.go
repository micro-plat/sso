package sqls

//QuerySystemInfo 获取系统信息
const QuerySystemInfo = `
select 
	t.id,
	t.name,
	'' as index_url,
	t.logo,
	t.theme,
	t.layout,
	t.ident,
	t.wechat_status,
	t.login_url,
	t.enable,
	t.secret
from sso_system_info t
where t.ident=@ident and
	  t.enable=1
`

//QuerySystemWechantStatus 获取系统微信登录状态
const QuerySystemWechantStatus = `
select 
	t.id,t.wechat_status 
from 
	sso_system_info t 
where 
	t.id=@sys_id`

//QueryAllSystemInfo .
const QueryAllSystemInfo = `
select 
	t.sys_id,s.name,s.index_url,s.ident 
from 
	sso_user_role t 
inner join 
	sso_system_info s on s.id=t.sys_id 
where 
	t.user_id=@user_id`
