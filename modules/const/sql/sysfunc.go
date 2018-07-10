package sql

const QuerySysFuncList = `select t.* from sso_system_menu t where t.sys_id=@sysid`

const EnableSysFunc = `update sso_system_menu t
set  t.enable = @enable
where t.id=@id
`

const DeleteSysFunc  = `delete from sso_system_menu where id = @id`

const EditSysFunc = `update sso_system_menu t
set  t.name=@name,t.icon=@icon,t.path=@path
where t.id=@id
`
