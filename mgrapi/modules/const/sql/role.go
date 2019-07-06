package sql

//QueryRoleInfoList 查询角色信息列表
const QueryRoleInfoList = `
select 
	t.role_id,
	t.name role_name,
	t.status,
	case when t.status = 0 then '正常' when t.status = 2 then '禁用' end status_label,
	t.create_time
from sso_role_info t
where 
	1 = 1 #role_sql
order by 
	t.role_id
limit @start, @ps
`

//QueryRoleInfoListCount 获取角色信息列表数量
const QueryRoleInfoListCount = `
select 
	count(1)
from 
	sso_role_info t
where 
	1 = 1 #role_sql
order by 
	t.role_id
`

//UpdateRoleStatus 修改角色状态
const UpdateRoleStatus = `
update 
	sso_role_info t
set 
	t.status = @status
where 
	t.role_id = @role_id
`

//DeleteRole 删除角色
const DeleteRole = `
delete from 
	sso_role_info
where 
	role_id = @role_id
`

//DeleteRoleMenu 删除角色
const DeleteRoleMenu = `
delete from 
	sso_role_menu
where 
	role_id = @role_id
`

//EditRoleInfo 编辑角色信息
const EditRoleInfo = `
update 
	sso_role_info t
set 
	t.status = @status, 
	t.name = @role_name
where 
	t.role_id = @role_id
`

//AddRoleInfo 添加角色信息
const AddRoleInfo = `
insert into 
	sso_role_info
	( name, status)
values
	( @role_name, @status)
`

//AddRoleAuth 添加角色权限
const AddRoleAuth = `
insert into 
	sso_role_menu
	( sys_id, role_id, menu_id, enable, sortrank)
values
	( @sys_id, @role_id, @menu_id, 1, @sortrank)
`

//DelRoleAuth 删除角色权限
const DelRoleAuth = `
delete from 
	sso_role_menu
where 
	sys_id = @sys_id
	and role_id = @role_id
`

//QuerySysMenucList 系统菜单获取
const QuerySysMenucList = `select t.id, 
t.name title, 
t.parent, 
t.sys_id, 
t.level_id,  
(case
	when t.id in (select menu_id
					from sso_role_menu rm
				   where rm.role_id = @role_id
					 and rm.sys_id = @sys_id) then
	 1
	else
	 0
  end) checked,
t.icon, 
t.path, 
t.enable, 
t.create_time, 
t.sortrank 
from sso_system_menu t 
where t.sys_id = @sys_id 
`

//GetPageAuth 获取页面授权tag
const GetPageAuth = `select t1.id,t1.path,t2.enable 
from sso_system_menu t1 
left join sso_role_menu t2 on t1.id = t2.menu_id
where t1.parent = (select id from sso_system_menu where path=@path) 
	and t2.sys_id=@sys_id 
	and t2.role_id=@role_id
`
