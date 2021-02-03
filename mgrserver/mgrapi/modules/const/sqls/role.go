package sqls

//QueryRoleInfoList 查询角色信息列表
const QueryRoleInfoList = `
select 
	t.role_id,
	t.name role_name,
	t.status,
	case when t.status = 0 then '启用' when t.status = 2 then '禁用' end status_label,
	t.create_time
from sso_role_info t
where 
	1 = 1 #role_sql
	and if(@status <> -1, t.status=@status,1=1)
order by 
	t.status,t.role_id
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
	and if(@status <> -1, t.status=@status,1=1)
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

// //AddRoleDataPermissionAuth 添加角色与数据权限数据的关系
// const AddRoleDataPermissionAuth = `
// insert into sso_role_datapermission
// (
// 	role_id,
// 	sys_id,
// 	table_name,
// 	operate_action,
// 	name,
// 	permissions,
// 	status
// )
// values
// (
// 	@role_id,
// 	@sys_id,
// 	@table_name,
// 	@operate_action,
// 	@name,
// 	@permissions,
// 	1
// )`

//AddRoleDataPermissionAuth 添加角色与数据权限数据的关系
const AddRoleDataPermissionAuth = `
insert into sso_role_datapermission(
	sys_id, 
	role_id, 
	permission_config_id)
values(
	@sys_id, 
	@role_id, 
	@permission_config_id)`

//CheckSysMeun
const CheckSysMeun = `
SELECT
t.id
FROM sso_system_menu t
where sys_id = @sys_id
and t.parent in (#select_auth)
`

//DelRoleAuth 删除角色权限
const DelRoleAuth = `
delete from 
	sso_role_menu
where 
	sys_id = @sys_id
	and role_id = @role_id
`

//DelDataPermissionRoleAuth 删除数据权限的关联关系
const DelDataPermissionRoleAuth = `
delete from sso_role_datapermission 
where sys_id = @sys_id and 
	  role_id = @role_id`

// //QuerySysMenucList 系统菜单获取
const QuerySysMenucList = `

select  t.id,
t.name title,
t.parent,
t.sys_id,
t.level_id,
rm.id rid ,
if(rm.id is null, 0,1) checked,
t.icon,
t.path,
t.enable,
t.create_time,
t.sortrank 
from  sso_system_menu t
left join  sso_role_menu rm
				   on rm.role_id = @role_id
					 and rm.sys_id = t.sys_id
                     and rm.menu_id = t.id
where t.sys_id = @sys_id and t.enable = 1     

`

const QueryRoleDataPermission = `
select 
	t.id, 
	t.name,
	t.table_name,
	t.operate_action,
	t.remark,
	t.sys_id, 
	t.status,
	(case when t.id in (select permission_config_id from sso_role_datapermission rm where rm.role_id = @role_id and rm.sys_id = @sys_id) then 1 else 0 end) checked
from sso_data_permission t 
where t.sys_id = @sys_id and
      t.status = 0
`

//GetPageAuth 获取页面授权tag
const GetPageAuth = `select t1.id,t1.path,t2.enable 
from sso_system_menu t1 
left join sso_role_menu t2 on t1.id = t2.menu_id
where t1.parent = (select id from sso_system_menu where path=@path) 
	and t2.sys_id=@sys_id 
	and t2.role_id=@role_id
`

//QueryRoleInfoByName 通过名称查询角色信息
const QueryRoleInfoByName = `select role_id, name, status, create_time from sso_role_info where name=@role_name`

// const QueryRoleDataPermission = `
// SELECT
//   id,
//   sys_id,
//   role_id,
//   table_name,
//   operate_action,
//   permissions,
//   status,
//   name,
//   DATE_FORMAT(create_time, '%y-%m-%d %h:%i:%s') as create_time
// from  sso_role_datapermission
// where sys_id = @sys_id and
// 	  role_id = @role_id
// limit @start, @ps
// `

const QueryRoleDataPermissionCount = `
SELECT  
  count(1) as count
from  sso_role_datapermission
where sys_id = @sys_id and
	  role_id = @role_id
`
const ChangeRolePermissionStatus = `
update sso_role_datapermission set 
	status = @status 
where id=@id 
limit 1`

const DeleteRolePermission = `
delete from sso_role_datapermission 
where id=@id 
limit 1`

const UpdateRolePermission = `
update sso_role_datapermission set 
	name = @name,
	permissions=@permissions
where id=@id
limit 1
`
