package sql

//QuerySysFuncList .
const QuerySysFuncList = `
select 
	t.* 
from 
	sso_system_menu t 
where 
	t.sys_id=@sysid 
order by 
	t.sortrank,t.id`

//EnableSysFunc .
const EnableSysFunc = `
update 
	sso_system_menu t
set  
	t.enable = @enable
where 
	t.id=@id
`

//DeleteSysFunc .
const DeleteSysFunc = `
delete from 
	sso_system_menu 
where 
	id = @id
`

//EditSysFunc .
const EditSysFunc = `
update 
	sso_system_menu t
set  
	t.name=@name,
	t.icon=@icon,
	t.path=@path,
	t.is_open=@is_open
where 
	t.id=@id
`

// AddSysFunc .
const AddSysFunc = `
insert into 
	sso_system_menu
	(name,parent,sys_id,level_id,icon,path,sortrank,is_open) 
values
	(@name,@parent,@sys_id,@level_id,@icon,@path,@@IDENTITY +1,@is_open)`
