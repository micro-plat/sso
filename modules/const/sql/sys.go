package sql

//QuerySystemInfo 获取系统信息
const QuerySystemInfo = `select t.* from sso_system_info t where t.id=@sys_id`
