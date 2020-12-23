package sqls

//QueryDataPermissionTotalCount 数据权限分页总条数
const QueryDataPermissionTotalCount = `
select
	count(1) as count
from sso_data_permission
where sys_id = @sys_id 
	  #name
	  #table_name
`

//QueryDataPermissionList 数据权限分页　数据
const QueryDataPermissionList = `
select
	id,
	sys_id,
	ident,
	name,
	table_name,
	operate_action,
	rules,
	remark,
	status
from sso_data_permission
where sys_id = @sys_id
	  #name
	  #table_name
limit @start, @ps
`

const DeletePermissionInfoById = `delete from sso_data_permission where id=@id limit 1`

const GetNotDefaultPermissionCount = `select count(1) as count from sso_data_permission where sys_id = @sys_id and type=@type and isall = 0`

const DeletePermissionDefaultData = `delete from sso_data_permission where sys_id = @sys_id and type=@type and isall = 1 limit 1`

const AddDataPermission = `
insert into sso_data_permission(
	sys_id,
	ident,
	name,
	table_name,  
	operate_action,
	rules,
	remark
)
VALUES(
	@sys_id,
	@ident,
	@name,
	@table_name,  
	@operate_action,
	@rules,
	@remark
)
`

//UpdateDataPermission 更新数据权限数据
const UpdateDataPermission = `
update sso_data_permission set 
	name = @name,
	operate_action = @operate_action,
	rules = @rules,
	remark = @remark
where id=@id
limit 1
`

const ChangePermissionConfigStatus = `
update sso_data_permission set 
	status = @status
where id=@id
limit 1
`
