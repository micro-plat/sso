package sql

const QuerySysFuncList = `select t.* from sso_system_menu t where t.sys_id=@sysid`

const EnableSysFunc = `update sso_system_menu t
set  t.enable = @enable
where t.id=@id
`

const DeleteSysFunc = `delete from sso_system_menu where id = @id`

const EditSysFunc = `update sso_system_menu t
set  t.name=@name,t.icon=@icon,t.path=@path,t.is_open=@is_open
where t.id=@id
`
const AddSysFunc = `insert into sso_system_menu(id,name,parent,sys_id,level_id,icon,path,sortrank,is_open) 
values(seq_system_menu_id.nextval,@name,@parent,@sys_id,@level_id,@icon,@path,seq_system_menu_id.nextval,@is_open)`
