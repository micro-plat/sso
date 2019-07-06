package sql

const QueryUserInfoList = `
select 
	t.user_id,
	t.user_name,
	t.status,
	case when t.status = 0 then '正常' when t.status=1 then '锁定' when t.status = 2 then '禁用' end status_label,
	t.mobile,
	t.create_time,
	t.email,
	t.ext_params
from sso_user_info t
left join sso_user_role r on r.user_id = t.user_id
where 
	1=1 
	and if(@role_id <> '',r.role_id=@role_id,1=1) 
	#user_name
group by t.user_id,
			t.user_name,
			t.status,
			t.mobile,
			t.email,
			t.create_time,
			t.ext_params
order by 
	t.user_id
limit limit @start, @ps
`

//QueryUserRoleList 查询用户角色信息列表
const QueryUserRoleList = `
select 
	a.user_id, 
	a.sys_id, 
	a.role_id, 
	s.name sys_name, 
	r.name role_name
from 
	sso_user_role a
inner join 
	sso_system_info s on s.id = a.sys_id
inner join 
	sso_role_info r on r.role_id = a.role_id
where 
	a.user_id in ( #user_id_string)
order by a.user_id, a.sys_id, a.role_id
`

//QueryUserInfoListCount 获取用户信息列表数量
const QueryUserInfoListCount = `
select 
	count(1)
from (select 
		t.user_id,
		t.user_name,
		t.status,
		case when t.status = 0 then '正常' when t.status=1 then '锁定' when t.status = 2 then '禁用' end status_label,
		t.mobile,
		t.create_time,
		t.email
	from 
		sso_user_info t
	left join 
		sso_user_role r on r.user_id = t.user_id
	where 
		1=1 
		and if(@role_id <> '',r.role_id=@role_id,1=1) 
		#user_name
	group by 
		t.user_id,
		t.user_name,
		t.status,
		t.mobile,
		t.email,
		t.create_time
	order by t.user_id) R`

//UpdateUserStatus 获取用户信息列表数量
const UpdateUserStatus = `
update 
	sso_user_info t
set 
	t.status = @status
where 
	t.user_id = @user_id
`

//DeleteUser 删除用户
const DeleteUser = `
delete from 
	sso_user_info
where 
	user_id = @user_id`

//QueryUserInfo 查询用户信息列表
const QueryUserInfo = `
select 
	t.user_id,
	t.user_name,
	t.mobile,
	t.email 
from 
	sso_user_info t 
where 
	t.user_id=@user_id`

//EditUserInfo 编辑用户信息
const EditUserInfo = `
update 
	sso_user_info t
set 
	t.status = @status, 
	t.user_name = @user_name, 
	t.mobile = @mobile,
	t.email = @email,
	t.ext_params = @ext_params
where 
	t.user_id = @user_id
`

//DelUserRole 删除用户角色
const DelUserRole = `
delete from 
	sso_user_role 
where 
	user_id=@user_id`

//EditUserRole 编辑用户角色
const EditUserRole = `
update 
	sso_user_role t 
set 
	t.role_id = @role_id 
where 
	t.user_id = @user_id`

/* //GetNewUserID 获取新用户ID 这个不用,mysql中处理方式不一样
//GetNewUserID 获取新用户ID
const GetNewUserID = `
select
	seq_user_info_id.nextval
from
	dual`
*/

//AddUserInfo 添加用户信息
const AddUserInfo = `
insert 
	into sso_user_info 
	(user_name, status, password, mobile, email, ext_params)
values
	(@user_name, @status, @password, @mobile, @email, @ext_params)
`

//AddUserRole 添加用户角色
const AddUserRole = `
insert into 
	sso_user_role
	( user_id, role_id, sys_id, enable)
values
	( @user_id, @role_id, @sys_id, 1)
`

//QueryUserByName .
const QueryUserByName = `
select 
	t.user_name,
	t.wx_openid 
from 
	sso_user_info t 
where 
	t.user_name=@user_name`

//QueryUserPswd 查询用户密码
const QueryUserPswd = `
select 
	count(1)
from 
	sso_user_info t
where 
	t.user_id=@user_id
 	&password
`

const EditInfo = `
update 
	sso_user_info t
set  
	t.mobile = @tel, t.email = @email
where 
	t.user_name = @username`

//QueryOldPwd .
const QueryOldPwd = `
select 
	t.password,
	t.changepwd_times 
from 
	sso_user_info t 
where 
	t.user_id=@user_id`

//SetNewPwd .
const SetNewPwd = `
update 
	sso_user_info t
set 
	t.password = @password,
	t.changepwd_times = t.changepwd_times + 1
where 
	t.user_id = @user_id`

//QueryUserBind .
const QueryUserBind = `
select 
	t.email,
	t.wx_openid 
from 
	sso_user_info t 
where 
	t.email=@email`

//ExecUserBind .
const ExecUserBind = `
update 
	sso_user_info t
set 
	t.wx_openid = @wx_openid
where 
	t.email = @email`

//QueryUserBySysCount .
const QueryUserBySysCount = `
select 
	count(1) 
from 
	sso_user_role t 
inner join 
	sso_user_info i on i.user_id=t.user_id 
where 
	t.sys_id=@sys_id`

//QueryUserBySysList .
const QueryUserBySysList = `
select 
	i.*
from 
	sso_user_role t 
inner join 
	sso_user_info i on i.user_id=t.user_id 
where 
	t.sys_id=@sys_id 
limit 
	#pageSize offset #currentPage
`
