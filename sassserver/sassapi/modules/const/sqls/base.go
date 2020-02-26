package sqls

//GetUserRoleList 获取用户角色列表
const GetUserRoleList = `
select 
  t.* 
from sso_role_info t
where belong_type = @belong_type
`

//GetSysList 获取系统列表
const GetSysList = `select * from sso_system_info t`
