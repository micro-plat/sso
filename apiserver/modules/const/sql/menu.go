package sql

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
where 
	r.user_id=@user_id 
	and r.sys_id=@sys_id 
	and r.enable=1 
	and  m.enable=1 
	and s.enable=1
	and rinfo.status = 0
order by s.parent asc,s.level_id desc,s.sortrank asc`

//QueryUserPopularMenus 查询常用菜单
//TODO:
const QueryUserPopularMenus = `select * from( select * from(select p.used_cnt, s.id,s.name,s.parent,s.sys_id,s.level_id,s.icon,s.path,to_char(s.create_time,'yyyy-mm-dd hh24:mi:ss') create_time,s.sortrank
from sso_user_role  r
inner join sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
inner join sso_user_popular p on p.user_id=r.user_id and p.sys_id=r.sys_id and p.parent_id=s.id
where r.user_id=@user_id and r.sys_id=@sys_id and r.enable=1 and  m.enable=1 and s.enable=1)
union all (
select p.used_cnt, s.id,s.name,s.parent,s.sys_id,s.level_id,s.icon,s.path,to_char(s.create_time,'yyyy-mm-dd hh24:mi:ss') create_time,s.sortrank
from sso_user_role  r
inner join sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
inner join sso_user_popular p on p.user_id=r.user_id and p.sys_id=r.sys_id and p.menu_id=s.id
where r.user_id=@user_id and r.sys_id=@sys_id and r.enable=1 and  m.enable=1 and s.enable=1
))
order by used_cnt desc`

//CheckUserPopularMenu 检查用户菜单是否存在
const CheckUserPopularMenu = `
select 
	count(1) 
from 
	sso_user_popular t 
where 
	t.user_id=@user_id
	and t.menu_id=@menu_id
	and t.sys_id=@sys_id`

//SaveUserPopularMenu 保存用户常用菜单
const SaveUserPopularMenu = `
insert into 
	sso_user_popular
	( user_id, sys_id, parent_id, menu_id, used_cnt)
values
	(@user_id,
	@sys_id,
	@parent_id,
	@menu_id,
	@used_cnt)
`

//UpdateUserPopularMenu 累加用户使用次数
const UpdateUserPopularMenu = `
update 
	sso_user_popular t
set  
	t.used_cnt = t.used_cnt+1
where 
	t.user_id=@user_id
	and t.menu_id=@menu_id
	and t.sys_id=@sys_id
`

//QueryUserMenu 查询用户菜单
const QueryUserMenu = `
select 
	count(1)
from 
	sso_user_role  r
inner join 
	sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join 
	sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
where 
	r.user_id=@user_id 
	and r.sys_id=@sys_id 
	and s.path in (#path) 
	and r.enable=1 
	and  m.enable=1 
	and s.enable=1`
