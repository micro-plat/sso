-- 增加表字段
alter table sso_user_info add source varchar(36) NOT NULL DEFAULT '' COMMENT '来源';
alter table sso_user_info add source_id varchar(120) NOT NULL DEFAULT '' COMMENT '来源id';
alter table sso_user_info add last_login_time datetime DEFAULT NULL COMMENT '最近登陆时间';

-- 系统信息
delete from sso_system_info where ident = 'crp-psms';
INSERT INTO sso_system_info ( 
	name, index_url, enable, login_timeout, 
	logo, theme, layout, ident, login_url, 
	wechat_status, secret
) VALUES (
'惠捷油站管理系统', '', 1, '3000', 
'', 'bg-info|bg-info|bg-light light-info', 'container', 'crp-psms', 'http://member/login', 
1, '311124b57e468ff88e4f1c8743354314');



-- 菜单数据
delete from sso_system_menu where sys_id  = (select id from sso_system_info where ident = 'crp-psms' limit 1);


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '加油系统', 0, s.id , '1', ' ', '-', '1', '2020-02-28 10:15:19', '1', '0'
from sso_system_info s
where s.ident = 'crp-psms';

INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)
select '油站管理', m.id, m.sys_id, '2', 'fa fa-folder-open-o text-success', '-', '1', '2020-03-02 09:54:49', '3', '1'
from sso_system_menu m 
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 1 and s.ident = 'crp-psms' and m.name = '加油系统';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)
select '统计报表', m.parent , m.sys_id, '2', 'fa fa-folder-open-o text-success', '-', '1', '2020-03-02 10:00:45', '5', '1'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油站管理' and path = '-' and s.ident = 'crp-psms';

 
INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '订单管理', m.parent , m.sys_id, '2', 'fa fa-folder-open-o text-success', '-', '1', '2020-02-28 10:16:52', '2', '1'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油站管理' and path = '-' and s.ident = 'crp-psms';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '油价管理', m.parent , m.sys_id, '2', 'fa fa-folder-open-o text-success', '-', '1', '2020-03-02 09:57:05', '4', '1'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油站管理' and path = '-' and s.ident = 'crp-psms';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '订单记录', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/oms/order/info', '1', '2020-02-28 10:17:35', '1', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '订单管理' and path = '-' and s.ident = 'crp-psms';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '退款历史', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/oms/refund/info', '1', '2020-02-28 10:18:17', '2', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '订单管理' and path = '-' and s.ident = 'crp-psms';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '油站信息', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/crp/station/info', '1', '2020-03-02 09:55:27', '1', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油站管理' and path = '-' and s.ident = 'crp-psms';

INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '油枪管理', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/crp/station/gun', '1', '2020-03-02 09:56:03', '2', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油站管理' and path = '-' and s.ident = 'crp-psms';



INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '活动油价', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/crp/station/activity', '1', '2020-03-02 09:58:41', '2', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油价管理' and path = '-' and s.ident = 'crp-psms';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '调价历史',  m.id, m.sys_id, '3', 'fa fa-circle text-success', '/crp/station/price/history', '1', '2020-03-02 09:59:10', '3', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油价管理' and path = '-' and s.ident = 'crp-psms';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '油价管理', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/crp/station/price/manage', '1', '2020-03-02 09:57:51', '1', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '油价管理' and path = '-' and s.ident = 'crp-psms';


INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '班结报表', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/crp/rep/statistical', '1', '2020-03-02 10:01:57', '1', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '统计报表' and path = '-' and s.ident = 'crp-psms';

INSERT INTO sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '交易日报', m.id, m.sys_id, '3', 'fa fa-circle text-success', '/crp/rep/daily', '1', '2020-03-02 10:02:21', '2', '0'
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id
where m.level_id = 2 and m.name = '统计报表' and path = '-' and s.ident = 'crp-psms';




-- 增加角色数据
delete from sso_role_info where name = '惠捷加油';
INSERT INTO sso_role_info (name, status, create_time) VALUES ('惠捷加油', '0', '2020-02-28 10:30:07');


-- 角色与菜单的关联
delete 
from sso_role_menu
where sys_id  = (select id from sso_system_info where ident = 'crp-psms' limit 1);

insert into sso_role_menu(
  sys_id,
  role_id,
  menu_id,
  enable,
  sortrank
)
select  
  m.sys_id,
  (select role_id from sso_role_info where name = '惠捷加油' LIMIT 1) as role_id,
  m.id,
  1,
  0
from sso_system_menu m
inner join sso_system_info s on m.sys_id = s.id 
where s.ident = 'crp-psms';


