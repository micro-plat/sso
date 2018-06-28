package sql

//QueryUserMenus 获取用户菜单信息
const QueryUserMenus = `select s.id,s.name,s.parent,s.sys_id,s.level_id,s.icon,s.path,to_char(s.create_time,'yyyy-mm-dd hh24:mi:ss') create_time,s.sortrank
from sso_user_role  r
inner join sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
where r.user_id=@user_id and r.sys_id=@sys_id and r.enable=1 and  m.enable=1 and s.enable=1
order by s.parent asc,s.level_id desc,s.sortrank asc`

//QueryUserPopularMenus 查询常用菜单
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

//QueryUserMenu 查询用户菜单
const QueryUserMenu = `select count(1)
from sso_user_role  r
inner join sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
where r.user_id=@user_id and r.sys_id=@sys_id and and s.path=@path r.enable=1 and  m.enable=1 and s.enable=1 and rownum<=1`
