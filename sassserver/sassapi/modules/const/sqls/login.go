package sqls

//UnLockMember 解锁被锁定的用户
const UnLockMember = `update sso_user_info set status = 0 where status = 1 and user_name = @user_name`

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
	and i.status=0`

//QueryUserByMobile 根据手机号获取用户信息
const QueryUserByMobile = `
select 
	user_id,
	user_name,
	password,
	status,
	wx_openid,
	ext_params,
	belong_id,
	belong_type 
from sso_user_info 
where mobile=@mobile
limit 1;`

//QueryUserMenus 获取用户菜单信息
const QueryUserMenus = `
select 
	s.id,
	s.name,
	s.parent,
	s.sys_id,
	s.level_id,
	s.icon,
	s.path,
	s.is_open,
	s.create_time,
	s.sortrank
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
order by s.parent asc,s.level_id desc,s.sortrank asc`

//QuerySystemInfo 获取系统信息
const QuerySystemInfo = `
select 
	t.id,
	t.name,
	t.index_url,
	t.logo,
	t.theme,
	t.layout,
	t.ident,
	t.wechat_status,
	t.login_url,
	t.secret 
from 
	sso_system_info t
where 
	t.ident=@ident
`
