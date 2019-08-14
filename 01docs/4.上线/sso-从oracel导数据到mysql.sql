/*mysql表要先创建*/

/*1: 第一部查询oracle数据
2: 导出oracle的查询数据(导出成文件)
3: 将insert 文件放到mysql中运行
*/

/*oracle*/ 导出文件的语句
SELECT * from sso_system_info;
select role_id, name, status, TO_CHAR(create_time, 'YYYY-MM-DD HH:MM:SS') create_time from sso_role_info;
select user_id, user_name, password, email, status, to_char(mobile) mobile, wx_openid, to_char(create_time, 'YYYY-MM-DD HH:MM:SS') create_time, changepwd_times, ext_params from sso_user_info;
select id, name, parent, sys_id, level_id, icon, path, enable, to_char(create_time, 'YYYY-MM-DD HH:MM:SS') create_time, sortrank, is_open  from sso_system_menu;
select sys_id, role_id,menu_id, enable, to_char(create_time, 'YYYY-MM-DD HH:MM:SS') create_time, sortrank from sso_role_menu;
select user_id, sys_id, role_id, enable from sso_user_role;



------mysql操作------

/* 以下为历史原因，原来对主键设置了0,而且是用来表示特殊意义，因此要作修改*/

/*更新系统表为 当ident = sso 表示为 用户权限系统 */
update sso_system_info set id = 0 where ident = 'sso'

update sso_system_info set index_url = '' 

/*更新角色表中的 管理员role_id 为0 */
update sso_role_info set role_id = 0 where name = '管理员';

/*更新某个人的user_id为0，可能是作为管理吧，不清楚，这个要在线上看一下*/
update sso_user_info set user_id = 0 where user_name = 'yanglei'



------------------------------------------------------------------------
#select id, user_id, sys_id, parent_id, menu_id, used_cnt, to_char(create_time, 'YYYY-MM-DD HH:MM:SS') create_time from sso_user_popular;
#select id, company_name, account, pwd from  sso_city_info;
#select id, user_id, sys_id, level_id, keywords, status, to_char(create_time,  'YYYY-MM-DD HH:MM:SS') create_time from sso_notify_subscribe;
#select id, user_id, sys_id, level_id, title, keywords, content, status, to_char(create_time,  'YYYY-MM-DD HH:MM:SS') create_time, send_count, scan_batch_id, to_char(flow_timeout,  'YYYY-MM-DD HH:MM:SS') flow_timeout, to_char(finish_time,  'YYYY-MM-DD HH:MM:SS') finish_time from sso_notify_user;



操作日志不用导出，不管
#select type, sys_id, user_id, to_char(create_time, 'YYYY-MM-DD HH:MM:SS') create_time, content from sso_operate_log;
#select id, sys_id, level_id, title, keywords, content, status, to_char(create_time, 'YYYY-MM-DD HH:MM:SS') create_time, case when finish_time = null then null else to_char(finish_time,'YYYY-MM-DD HH:MM:SS') end finish_time  from sso_notify_records;
