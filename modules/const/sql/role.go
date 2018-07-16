package sql

//QueryRoleInfoList 查询角色信息列表
const QueryRoleInfoList = `select TAB1.*
from (select L.*
		from (select rownum LINENUM, R.*
				from (select t.role_id,
							 t.name role_name,
							 t.status,
							 decode(t.status, 0, '正常', 2, '禁用') status_label,
							 to_char(t.create_time, 'yyyy/mm/dd hh24:mi') create_time
						from sso_role_info t
					   where 1 = 1
					   		 #role_sql
					   order by t.role_id) R
			   where rownum <= @pi * @ps) L
	   where L.LINENUM > @ps * (@pi - 1)) TAB1

`

//QueryRoleInfoListCount 获取角色信息列表数量
const QueryRoleInfoListCount = `select count(1)
from (select t.role_id,
	t.name role_name,
	t.status,
	decode(t.status, 0, '正常', 2, '禁用') status_label,
	to_char(t.create_time, 'yyyy/mm/dd hh24:mi') create_time
from sso_role_info t
where 1 = 1
	   #role_sql
order by t.role_id) R`

//UpdateRoleStatus 修改角色状态
const UpdateRoleStatus = `update sso_role_info t
set t.status = @status
where t.role_id = @role_id
and t.status = @ex_status
`

//DeleteRole 删除角色
const DeleteRole = `delete from sso_role_info t where t.role_id = @role_id`

//EditRoleInfo 编辑角色信息
const EditRoleInfo = `update sso_role_info t
set t.status = @status, t.name = @role_name
where t.role_id = @role_id
`

//GetNewRoleID 获取新角色ID
const GetNewRoleID = `select seq_role_info_id.nextval from dual`

//AddRoleInfo 添加角色信息
const AddRoleInfo = `insert into sso_role_info t
(role_id, name, status)
values
(seq_role_info_id.nextval, @role_name, @status)
`

//AddRoleAuth 添加角色权限
const AddRoleAuth = `insert into sso_role_menu t
(id, sys_id, role_id, menu_id, enable, sortrank)
values
(seq_role_menu_id.nextval, @sys_id, @role_id, @menu_id, 1, @sortrank)
`

//DelRoleAuth 添加角色权限
const DelRoleAuth = `delete from sso_role_menu t
where t.sys_id = @sys_id
and t.role_id = @role_id
`

//QuerySysMenucList 系统菜单获取
const QuerySysMenucList = `select t.id, 
t.name title, 
t.parent, 
t.sys_id, 
t.level_id, 
'true' as expanded, 
t.icon, 
t.path, 
t.enable, 
to_char(t.create_time, 'yyyy/mm/dd hh:mi') create_time, 
t.sortrank 
from sso_system_menu t 
where t.sys_id = @sys_id 
`
