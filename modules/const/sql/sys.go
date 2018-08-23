package sql

//QuerySystemInfo 获取系统信息
const QuerySystemInfo = `select t.id,t.name,t.index_url,t.logo,t.theme,t.layout,t.ident,t.wechat_status from sso_system_info t
where t.ident=@ident`

const QueryAllSystemInfo = `select t.sys_id,s.name,s.index_url,s.ident from sso_user_role t 
left join sso_system_info s on s.id=t.sys_id where t.user_id=@user_id`
