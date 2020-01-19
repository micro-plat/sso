package sqls

//QuerySystemMenuInfo 获取要导出的系统菜单
const QuerySystemMenuInfo = `
select 
	id,
	name,
	parent,
	sys_id,
	level_id,
	icon,
	path,
	enable,
	sortrank,
	is_open
from sso_system_menu
where sys_id = @sys_id
`

const ExistsSystemMenu = `
select  
  count(1)
from sso_system_menu 
where sys_id = @sys_id
`

//AddSystemMenu 增加菜单数据
const AddSystemMenu = `
insert into sso_system_menu(
	name,
	parent,
	sys_id,
	level_id,
	icon,
	path,
	enable,
	sortrank,
	is_open
)
values(
	@name,
	@parent,
	@sys_id,
	@level_id,
	@icon,
	@path,
	@enable,
	@sortrank,
	@is_open
)
`
