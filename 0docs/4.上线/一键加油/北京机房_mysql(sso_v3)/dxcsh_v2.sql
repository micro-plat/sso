-- 系统信息
delete from sso_system_info where ident = 'dxcsh_v2';
INSERT INTO sso_system_info ( 
	name, index_url, enable, login_timeout, 
	logo, theme, layout, ident, login_url, 
	wechat_status, secret
) VALUES (
'大象生活网-v2', 'http://bss.carlife.18jiayou.com:9199/ssocallback', 1, '3000', 
'http://bj.images.cdqykj.cn/sso/4cd16a1b5b36ac980713ce25c3c7ab0f.png', 'bg-primary|bg-primary|bg-dark dark-primary	app-header-fixed', 'container', 'dxcsh_v2', 'http://member/login', 
1, 'B128F779D5741E701923346F7FA9F95C');


delete from sso_system_menu where sys_id  = (select id from sso_system_info where ident = 'dxcsh_v2' limit 1);

-- 增加一级菜单

insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '交易承载', 0, s.id, '1', ' ', '-', '1', now(), '1', '0'  
from sso_system_info s  
where s.ident = 'dxcsh_v2';

insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '基础配置', 0, s.id, '1', ' ', '-', '1', now(), '2', '0'  
from sso_system_info s  
where s.ident = 'dxcsh_v2';


insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '商户配置', m.id, m.sys_id, '2', 'fa fa-cog ', '-', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'dxcsh_v2' and m.name = '交易承载';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '下游商户', m.id, m.sys_id, '3', ' ', '/merchant/index', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '商户配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '行业终端活动', m.id, m.sys_id, '2', 'fa fa-paper-plane-o text-primary', '-', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'dxcsh_v2' and m.name = '交易承载';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '活动管理', m.id, m.sys_id, '3', ' ', '/activity/index', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端活动';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '活动产品', m.id, m.sys_id, '3', ' ', '/product/index', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端活动';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '商户产品', m.id, m.sys_id, '3', ' ', '/merchant/product/index', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端活动';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '行业终端订单', m.id, m.sys_id, '2', 'fa fa-tags text-info', '-', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'dxcsh_v2' and m.name = '交易承载';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '订单列表', m.id, m.sys_id, '3', ' ', '/order/list', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '支付列表', m.id, m.sys_id, '3', ' ', '/pay/list', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '退款列表', m.id, m.sys_id, '3', ' ', '/refund/list', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '单笔查询', m.id, m.sys_id, '3', ' ', '/trd/order/main/single', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '行业终端配置', m.id, m.sys_id, '2', 'fa fa-share-alt text-primary', '-', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'dxcsh_v2' and m.name = '基础配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '行业终端用户', m.id, m.sys_id, '3', ' ', '/user', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '渠道错误码', m.id, m.sys_id, '3', ' ', '/channel/errorcode', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '渠道分组', m.id, m.sys_id, '3', ' ', '/channel/groupcode', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '渠道管理', m.id, m.sys_id, '3', ' ', '/channel/index', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '芯片卡白名单', m.id, m.sys_id, '3', ' ', '/chipcard/list', '1', now(), '5', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '操作日志', m.id, m.sys_id, '3', ' ', '/operator/log/index', '1', now(), '6', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '数据字典', m.id, m.sys_id, '3', ' ', '/base/dictionary/info', '1', now(), '7', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '优惠券列表', m.id, m.sys_id, '3', ' ', '/coupon/modul/list', '0', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '商户配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '优惠券列表', m.id, m.sys_id, '3', ' ', '/coupon/list', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端活动';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '发货列表', m.id, m.sys_id, '3', ' ', '/trd/order/delivery', '1', now(), '5', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '关键字', m.id, m.sys_id, '3', ' ', '/wechat/keyword', '1', now(), '8', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '消息模板', m.id, m.sys_id, '3', ' ', '/message/template/index', '1', now(), '9', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '自动回复', m.id, m.sys_id, '3', ' ', '/wechat/scene', '1', now(), '10', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '提现列表', m.id, m.sys_id, '3', ' ', '/cash/info', '1', now(), '5', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '行业终端活动';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '新增页面测试', m.id, m.sys_id, '2', 'fa fa-file-image-o text-info', '-', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'dxcsh_v2' and m.name = '交易承载';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '油号订单', m.id, m.sys_id, '3', ' ', '/order/oil', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '新增页面测试';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '退款申请', m.id, m.sys_id, '3', ' ', '/order/refund/apply', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'dxcsh_v2' and m.name = '新增页面测试';

