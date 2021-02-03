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

const GetRoleMenus = `
select 
	s.id,
	s.name,
	s.parent,
	s.sys_id,
	s.level_id,
 	s.path  ,
    m.menu_id,
    s2.path tag   
 from sso_role_info rinfo 
inner join sso_system_info sys on sys.ident  = @ident  and rinfo.role_id = @role_id and rinfo.status = 0
inner join sso_role_menu m on m.role_id = @role_id   and sys.id=m.sys_id and  m.enable=1 
inner join sso_system_menu s on s.sys_id=sys.id and s.id=m.menu_id and s.enable=1 and s.level_id in (3)
left join (
select sm.path,sm.parent,sm.id from sso_system_menu sm 
inner join sso_system_info si on sm.sys_id = si.id and si.ident =  @ident and sm.enable = 1  
inner join sso_role_menu rm on sm.id = rm .menu_id  and rm.role_id = @role_id 
) s2  on s2.parent = s.id 

order by s.parent asc,s.level_id desc,s.sortrank asc
`
