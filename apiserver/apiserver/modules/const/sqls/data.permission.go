package sqls

//AddDataPermission 新增 [数据权限] 数据
const AddDataPermission = `
insert into sso_data_permission(
	sys_id,
	ident,
	name,
	type,
	value,
	remark
)
VALUES(
	@sys_id,
	@ident,
	@name,
	@type,
	@value,
	@remark
)
`

//QueryUserDataPermission 获取某个用户的某个类型下的 [数据权限] 信息
const QueryUserDataPermission = `
select 
	d.id,
	d.name,
	d.sys_id,
	d.type,
	d.value,
	d.remark
from sso_user_role  r
inner join sso_role_info rinfo on rinfo.role_id = r.role_id
inner join sso_role_datapermission p on p.sys_id = r.sys_id and p.role_id = r.role_id 
inner join sso_data_permission d on d.sys_id = p.sys_id and d.id = p.permission_id
inner join sso_system_info sys on sys.id = r.sys_id
where r.user_id=@user_id 
	and sys.ident=@ident
  	and d.type=@type
	and r.enable=1 
	and rinfo.status = 0
order by d.id
`
