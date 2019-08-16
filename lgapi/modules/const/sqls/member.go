package sqls

//QueryUserByUserName 根据用户名获取用户信息
const QueryUserByUserName = `
select 
	user_id,user_name,password,status,wx_openid,ext_params 
from sso_user_info 
where user_name=@user_name
limit 1;`

//AddUserOpenID 给用户绑定openid
const AddUserOpenID = `
update sso_user_info set 
	wx_openid = @openid 
where user_name = @username and 
	  wx_openid is not null 
limit 1;`

//QueryUserInfoByUID 查询用户信息
const QueryUserInfoByUID = `
select 
	u.user_id,u.user_name,u.mobile,u.wx_openid,u.status 
from sso_user_info u
where u.user_id=@user_id;`

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
