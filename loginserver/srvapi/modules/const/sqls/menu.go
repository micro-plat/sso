package sqls

//QueryUserMenus 获取用户菜单信息
const QueryUserMenus = `
select 
s.id,
s.name,
s.parent,
s.level_id,
s.icon,
s.path,
s.is_open
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
	and s.menu_type in (0,1)
	and rinfo.status = 0
order by s.parent asc,s.level_id desc,s.sortrank asc`

//QueryUserMenuTags 返回用户按钮级的权限信息
const QueryUserMenuTags = `
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
    and s.level_id = 4
order by s.parent asc,s.level_id desc,s.sortrank asc
`
