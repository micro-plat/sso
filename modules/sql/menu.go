package sql

//QueryUserMenus 获取用户菜单信息
const QueryUserMenus = `select s.*
from sso_user_role  r
inner join sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
where r.user_id=@user_id and r.sys_id=@sys_id and r.enable=0 and  m.enable=1 and s.enable=1`

//QueryUserMenu 查询用户菜单
const QueryUserMenu = `select count(1)
from sso_user_role  r
inner join sso_role_menu m on r.role_id=m.role_id and r.sys_id=m.sys_id
inner join sso_system_menu s on s.sys_id=m.sys_id and s.id=m.menu_id
where r.user_id=@user_id and r.sys_id=@sys_id and and s.path=@path r.enable=1 and  m.enable=1 and s.enable=1 and rownum<=1`
