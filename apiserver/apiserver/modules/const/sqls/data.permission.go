package sqls

//QueryUserDataPermission 获取某个用户的某个类型下的 [数据权限] 信息
const QueryUserDataPermission = `
select 
  s.id,
  s.name,
  s.rules
from sso_user_role  r
inner join sso_role_info rinfo on rinfo.role_id = r.role_id
inner join sso_role_datapermission p on p.sys_id = r.sys_id and p.role_id = r.role_id 
inner join sso_data_permission s on s.sys_id = p.sys_id and s.id=p.permission_config_id
inner join sso_system_info sys on sys.id = r.sys_id
where r.user_id=@user_id
	and sys.ident=@ident
	and r.enable=1 
	and rinfo.status = 0
  and s.table_name = @table_name
  and s.operate_action=@operate_action
  and s.status = 0
`

const QueryPermissionConfig = `
select 
   id, name, rules
from sso_data_permission
where id in ( #ids )
`

const GetAllUserInfoByUserRole = `
select 
  DISTINCT
  user_id
from sso_user_role o
inner join (
		select 
			r.role_id,r.sys_id
		from sso_user_role  r
		inner join sso_role_info rinfo on rinfo.role_id = r.role_id
		inner join sso_system_info sys on sys.id = r.sys_id
		where r.user_id=@user_id 
			and sys.ident=@ident
			and r.enable=1 
			and rinfo.status = 0
) t on o.role_id = t.role_id and o.sys_id = t.sys_id
`
