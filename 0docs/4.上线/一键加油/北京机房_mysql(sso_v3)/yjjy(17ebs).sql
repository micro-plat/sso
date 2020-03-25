-- 系统信息
delete from sso_system_info where ident = 'yjjy';
INSERT INTO sso_system_info ( 
	name, index_url, enable, login_timeout, 
	logo, theme, layout, ident, login_url, 
	wechat_status, secret
) VALUES (
'17EBS', 'http://bss.17ebs.18jiayou.com/ssocallback', 1, '3000', 
'', 'bg-info|bg-info|bg-light light-info', 'container', 'yjjy', 'http://member/login', 
1, 'B128F779D5741E701923346F7FA9F95C');



delete from sso_system_menu where sys_id  = (select id from sso_system_info where ident = 'yjjy' limit 1);

insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '系统配置', 0, s.id, '1', ' ', '-', '1', now(), '7', '0'  
from sso_system_info s  
where s.ident = 'yjjy';

insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '一键加油交易', 0, s.id, '1', ' ', '-', '1', now(), '3', '0'  
from sso_system_info s  
where s.ident = 'yjjy';

insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) 
select '一键加油渠道', 0, s.id, '1', ' ', '-', '1', now(), '5', '0'  
from sso_system_info s  
where s.ident = 'yjjy';


insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '交易订单', m.id, m.sys_id, '3', ' ', '/oms/order/info', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '发货管理', m.id, m.sys_id, '3', ' ', '/oms/order/delivery', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '下游通知', m.id, m.sys_id, '3', ' ', '/oms/notify/info', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '退款管理', m.id, m.sys_id, '3', ' ', '/oms/refund/info', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '退款管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '上游退货', m.id, m.sys_id, '3', ' ', '/oms/refund/up/return', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '退款管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '人工审核', m.id, m.sys_id, '3', ' ', '/oms/audit/info', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单审核';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '发货明细', m.id, m.sys_id, '3', ' ', '/ebs/order/delivery/coupon', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '优惠券订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '核销记录', m.id, m.sys_id, '3', ' ', '/ebs/coupon/consume', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '优惠券订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '冲正记录', m.id, m.sys_id, '3', ' ', '/ebs/coupon/reverse', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '优惠券订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '核销对账', m.id, m.sys_id, '3', ' ', '/ebs/coupon/compare', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '优惠券订单';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '基础数据', m.id, m.sys_id, '2', 'fa fa-cog text-danger', '-', '1', now(), '1', '1'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'yjjy' and m.name = '系统配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '产品线', m.id, m.sys_id, '3', ' ', '/oms/product/line', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '基础数据';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '省市信息', m.id, m.sys_id, '3', ' ', '/oms/canton/info', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '基础数据';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '字典信息', m.id, m.sys_id, '3', ' ', '/dds/dictionary/info', '1', now(), '5', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '基础数据';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '利润统计', m.id, m.sys_id, '3', ' ', '/report/profit', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '报表统计';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '上游报表', m.id, m.sys_id, '3', ' ', '/report/up/channel', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '报表统计';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '下游报表', m.id, m.sys_id, '3', ' ', '/report/down/channel', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '报表统计';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '电子发票申请记录', m.id, m.sys_id, '3', ' ', '/crp/invoice/apply', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '订单管理', m.id, m.sys_id, '2', 'fa fa-tags text-success', '-', '1', now(), '1', '1'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'yjjy' and m.name = '一键加油交易';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '订单管理', m.id, m.sys_id, '3', ' ', '/crp/order/info/extend', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '发货管理', m.id, m.sys_id, '3', ' ', '/crp/order/delivery/extend', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '通知管理', m.id, m.sys_id, '3', ' ', '/crp/notify/info', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '销券渠道', m.id, m.sys_id, '3', ' ', '/ebs/consume/channel', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '销券管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '产品管理', m.id, m.sys_id, '3', ' ', '/ebs/consume/up/channel/map', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '销券管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '油站配置', m.id, m.sys_id, '2', 'fa fa-share-alt-square text-info', '-', '1', now(), '2', '1'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'yjjy' and m.name = '系统配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '活动计划', m.id, m.sys_id, '3', ' ', '/crp/station/activity/schedule', '1', now(), '6', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '油站配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '油价计划', m.id, m.sys_id, '3', ' ', '/crp/up/product/price/schedule', '1', now(), '7', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '油站配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '变更记录', m.id, m.sys_id, '3', ' ', '/crp/up/product/price/history', '1', now(), '8', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '油站配置';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '退款管理', m.id, m.sys_id, '2', 'fa fa-reply text-success', '-', '1', now(), '2', '1'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'yjjy' and m.name = '一键加油交易';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '预下单管理', m.id, m.sys_id, '3', ' ', '/crp/pre/order/info', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '退款申请', m.id, m.sys_id, '3', ' ', '/crp/refund/apply', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '退款管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '退款管理', m.id, m.sys_id, '3', ' ', '/crp/refund/info/extend', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '退款管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '渠道管理', m.id, m.sys_id, '3', ' ', '/oms/down/channel', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '货架管理', m.id, m.sys_id, '3', ' ', '/oms/down/shelf', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '产品配置', m.id, m.sys_id, '3', ' ', '/oms/down/product', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '渠道管理', m.id, m.sys_id, '3', ' ', '/oms/up/channel', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '货架管理', m.id, m.sys_id, '3', ' ', '/oms/up/shelf', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '产品配置', m.id, m.sys_id, '3', ' ', '/oms/up/product', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '下游管理', m.id, m.sys_id, '2', 'fa fa-sitemap text-primary', '-', '1', now(), '1', '1'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'yjjy' and m.name = '一键加油渠道';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '上游管理', m.id, m.sys_id, '2', 'fa fa-paper-plane text-primary', '-', '1', now(), '2', '1'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'yjjy' and m.name = '一键加油渠道';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '价格配置', m.id, m.sys_id, '3', ' ', '/crp/down/product/extend', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '价格明细', m.id, m.sys_id, '3', ' ', '/crp/down/product/detail', '1', now(), '5', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '货架管理', m.id, m.sys_id, '3', ' ', '/oil/line/up/shelf', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '成本规则', m.id, m.sys_id, '3', ' ', '/crp/up/price/calc/config', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '价格配置', m.id, m.sys_id, '3', ' ', '/crp/up/product/extend', '1', now(), '5', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '下游账户', m.id, m.sys_id, '3', ' ', '/beanpay/downaccount/info', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '资金管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '下游资金变动', m.id, m.sys_id, '3', ' ', '/beanpay/downaccount/record', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '资金管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '开票申请', m.id, m.sys_id, '3', ' ', '/crp/invoice/apply', '1', now(), '5', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '订单管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '渠道管理', m.id, m.sys_id, '3', ' ', '/crp/down/channel', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '渠道管理', m.id, m.sys_id, '3', ' ', '/oil/line/up/channel', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '上游账户', m.id, m.sys_id, '3', ' ', '/beanpay/upaccount/info', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '资金管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '上游资金变动', m.id, m.sys_id, '3', ' ', '/beanpay/upaccount/record', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '资金管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '货架管理', m.id, m.sys_id, '3', ' ', '/crp/down/shelf', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '产品管理', m.id, m.sys_id, '3', ' ', '/crp/down/product', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '下游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '产品管理', m.id, m.sys_id, '3', ' ', '/oil/line/up/product', '1', now(), '4', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '发货配置', m.id, m.sys_id, '3', ' ', '/oil/line/vds/channel', '1', now(), '6', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '发货错误码', m.id, m.sys_id, '3', ' ', '/oil/line/vds/channel/error', '1', now(), '7', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '上游管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '油站管理', m.id, m.sys_id, '2', 'fa fa-sitemap text-success', '-', '1', now(), '3', '1'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 1  and s.ident = 'yjjy' and m.name = '一键加油渠道';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '油站信息', m.id, m.sys_id, '3', ' ', '/crp/station/info', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '油站管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '油站油枪', m.id, m.sys_id, '3', ' ', '/crp/station/gun', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '油站管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '油站油号', m.id, m.sys_id, '3', ' ', '/crp/station/oil', '1', now(), '3', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '油站管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '销券渠道', m.id, m.sys_id, '3', ' ', '/ebs/consume/channel', '1', now(), '1', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '销券管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '产品管理', m.id, m.sys_id, '3', ' ', '/ebs/consume/up/channel/map', '1', now(), '2', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '销券管理';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '国标价管理', m.id, m.sys_id, '3', ' ', '/crp/oil/standard/price', '1', now(), '6', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '基础数据';
insert into sso_system_menu (name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open) select '管理公司配置', m.id, m.sys_id, '3', ' ', '/crp/station/manage/company', '1', now(), '7', '0'  from sso_system_menu m  inner join sso_system_info s on m.sys_id = s.id where m.level_id = 2  and s.ident = 'yjjy' and m.name = '基础数据';
