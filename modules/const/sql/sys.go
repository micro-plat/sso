package sql

//QuerySystemInfo 获取系统信息
const QuerySystemInfo = `select t.id,t.name,t.index_url,t.logo,t.theme,t.layout,t.rowid from sso_system_info t
where t.id=@sys_id`
