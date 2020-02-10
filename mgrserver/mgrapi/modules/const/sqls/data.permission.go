package sqls

//QueryDataPermissionTotalCount 数据权限分页总条数
const QueryDataPermissionTotalCount = `
select
	count(1) as count
from sso_data_permission
where isall =0
      and sys_id = @sys_id
	  &type
`

//QueryDataPermissionList 数据权限分页　数据
const QueryDataPermissionList = `
select
	id,
	sys_id,
	ident,
	name,
	type,
	type_name,
	value,
	remark
from sso_data_permission
where isall =0
      and sys_id = @sys_id
	  &type
limit @start, @ps
`

const GetPermissionInfoByType = `
select 
	sys_id,
	type
from sso_data_permission 
where id = @id
limit 1 `

const DeletePermissionInfoById = `delete from sso_data_permission where id=@id and isall = 0 limit 1`

const GetNotDefaultPermissionCount = `select count(1) as count from sso_data_permission where sys_id = @sys_id and type=@type and isall = 0`

const DeletePermissionDefaultData = `delete from sso_data_permission where sys_id = @sys_id and type=@type and isall = 1 limit 1`

const AddDataPermission = `
insert into sso_data_permission(
	sys_id,
	ident,
	name,
	type,
	type_name,
	value,
	remark
)
VALUES(
	@sys_id,
	@ident,
	@name,
	@type,
	@type_name,
	@value,
	@remark
)
`

//AddDefaultDataPermissionInfo 增加一个默认全部
const AddDefaultDataPermissionInfo = `
insert into sso_data_permission(
	sys_id,
	ident,
	name,
	type,
	type_name,
	value,
	isall,
	remark
)
select 
	@sys_id,
	@ident,
	'全部',
	@type,
	@type_name,
	'*',
	1,
	'全部'
from DUAL
where NOT EXISTS (SELECT 1 FROM sso_data_permission WHERE sys_id=@sys_id and type=@type and value='*')
`

//UpdateDataPermission 更新数据权限数据
const UpdateDataPermission = `
update sso_data_permission set 
	name = @name,
	value = @value,
	remark = @remark
where id=@id
limit 1
`
