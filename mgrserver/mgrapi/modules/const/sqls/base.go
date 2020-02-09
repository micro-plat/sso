package sqls

//GetUserRoleList 获取用户角色列表
const GetUserRoleList = `select * from sso_role_info t`

//GetSysList 获取系统列表
const GetSysList = `select * from sso_system_info t`

//GetPermissTypes types list
const GetPermissTypes = `
select 
	DISTINCT
	type,
  	type_name
from sso_data_permission 
where sys_id = @sys_id
`
