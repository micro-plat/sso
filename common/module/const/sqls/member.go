package sqls

//QueryUserByUserName 根据用户名获取用户信息
const QueryUserByUserName = `
select 
	user_id,
	user_name,
	full_name,
	password,
	status,
	wx_openid,
	ext_params,
	source,
	source_id,
	last_login_time 
from sso_user_info 
where user_name=@user_name
limit 1;`

//UnLockMember 解锁被锁定的用户
const UnLockMember = `update sso_user_info set status = 0 where status = 1 and user_name = @user_name;`

//AddUserOpenID 给用户绑定openid
const AddUserOpenID = `
update sso_user_info set 
	wx_openid = @openid
where user_id = @user_id and 
	  (wx_openid is null or  wx_openid = '') 
limit 1;`

//OpenIDIsExists 判断当前的openid是否已绑定过用户
const OpenIDIsExists = `SELECT 1 FROM sso_user_info u WHERE u.wx_openid = @openid limit 1;`

//QueryUserInfoByUID 查询用户信息
const QueryUserInfoByUID = `
select 
	u.user_id,u.user_name,u.mobile,u.wx_openid,u.status 
from sso_user_info u
where u.user_id=@user_id;`

//UpdateUserStatus 更新用户状态
const UpdateUserStatus = `
update sso_user_info set status = @status where user_id = @user_id limit 1;`

//QueryUserInfoByOpenID 查询用户信息
const QueryUserInfoByOpenID = `
select 
	u.user_id,u.user_name,u.password,u.status 
from sso_user_info u
where u.wx_openid=@open_id; `

//QueryUserRoleCount 查询yonghu角色count
const QueryUserRoleCount = `
select 
	count(1)
from sso_user_role r
inner join sso_system_info s on r.sys_id=s.id
inner join sso_role_info i on i.role_id=r.role_id 
where 
	r.user_id=@user_id 
	and (@ident is null or @ident=s.ident) 
	and r.enable=1 
	and s.enable=1
	and i.status=0;`

//QueryUserRole 查询系统角色列表
const QueryUserRole = `
select 
	r.role_id,i.name role_name,s.index_url,s.login_url,s.login_timeout,r.sys_id 
from sso_user_role r
inner join sso_system_info s on r.sys_id=s.id
inner join sso_role_info i on i.role_id=r.role_id 
where 
	r.user_id=@user_id 
	and (@ident is null or @ident=s.ident)  
	and r.enable=1 
	and s.enable=1
	and i.status=0;`

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

const UpdateUserLoginTime = `
update sso_user_info set 
	last_login_time = now() 
where user_id = @user_id
limit 1
`

const QueryUserHasRoleMenuCount = `
select 
	count(s.id) as count
from sso_user_role  r
inner join sso_role_info rinfo on rinfo.role_id = r.role_id
inner join sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
inner join sso_system_info sys on sys.id = r.sys_id
where r.user_id=@user_id 
	and sys.ident=@ident
	and r.enable=1 
	and m.enable=1 
	and s.enable=1
	and rinfo.status = 0
`
