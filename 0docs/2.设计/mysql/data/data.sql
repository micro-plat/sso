INSERT INTO dds_dictionary_info(id,name,value,type,sort_no,group_code,status) VALUES (1, '新增', '新增', 'operate_action', 0,'*', 0);
INSERT INTO dds_dictionary_info(id,name,value,type,sort_no,group_code,status) VALUES (2, '修改', '修改', 'operate_action', 0,'*', 0);
INSERT INTO dds_dictionary_info(id,name,value,type,sort_no,group_code,status) VALUES (3, '启用', '启用', 'operate_action', 0,'*', 0);
INSERT INTO dds_dictionary_info(id,name,value,type,sort_no,group_code,status) VALUES (4, '禁用', '禁用', 'operate_action', 0,'*', 0);
INSERT INTO dds_dictionary_info(id,name,value,type,sort_no,group_code,status) VALUES (5, '锁定', '锁定', 'operate_action', 0,'*', 0);
INSERT INTO dds_dictionary_info(id,name,value,type,sort_no,group_code,status) VALUES (6, '启用', '0', 'role_status', 0,'*', 0);
INSERT INTO dds_dictionary_info(id,name,value,type,sort_no,group_code,status) VALUES (7, '禁用', '2', 'role_status', 0,'*', 0);

INSERT INTO sso_role_info (role_id, name, status, create_time) VALUES (1, '管理员', 0, '2018-6-27 03:06:54');

INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (1, 1, 1, 1, 1, '2019-12-17 14:31:33', 1);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (2, 1, 1, 2, 1, '2019-12-17 14:31:33', 2);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (3, 1, 1, 4, 1, '2019-12-17 14:31:33', 3);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (4, 1, 1, 7, 1, '2019-12-17 14:31:33', 4);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (5, 1, 1, 8, 1, '2019-12-17 14:31:33', 5);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (6, 1, 1, 9, 1, '2019-12-17 14:31:33', 6);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (7, 1, 1, 5, 1, '2019-12-17 14:31:33', 7);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (8, 1, 1, 3, 1, '2019-12-17 14:31:33', 8);
INSERT INTO sso_role_menu(id, sys_id, role_id, menu_id, enable, create_time, sortrank) VALUES (9, 1, 1, 6, 1, '2019-12-17 14:31:33', 9);

INSERT INTO sso_system_info (id, name, index_url, enable, login_timeout, logo, theme, layout, ident, login_url, wechat_status, secret) VALUES (1, '用户系统', 'http://ssov4.100bm.com', 1, 3000, 'http://ssov4.100bm.com/sso/4cd16a1b5b36ac980713ce25c3c7ab0f.png', 'bg-danger|bg-danger|bg-dark dark-danger', 'app-header-fixed', 'sso', '/user/index', 0, 'B128F779D5741E701923346F7FA9F95C');


INSERT INTO sso_system_menu(id, name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)VALUES (1, '用户权限', 0, 1, 1, '-', '-', 1, '2018-6-29 2:06:10', 2, 1);
INSERT INTO sso_system_menu(id, name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)VALUES (2, '用户角色', 1, 1, 2, 'fa fa-users text-info', '-', 1, '2018-7-13 02:07:31', 2, 1);
INSERT INTO sso_system_menu(id, name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)VALUES (3, '系统功能', 1, 1, 2, 'fa fa-tasks text-info-lter', '-', 1, '2018-6-28 10:06:19', 3, 1);
INSERT INTO sso_system_menu(id, name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)VALUES (4, '用户管理', 2, 1, 3, 'fa fa-user-circle text-primary', '/user/index', 1, '2018-6-28 10:06:39', 1, 1);
INSERT INTO sso_system_menu(id, name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)VALUES (5, '角色权限', 2, 1, 3, 'fa fa-users text-danger', '/role/index', 1, '2018-6-28 10:06:52', 2, 1);
INSERT INTO sso_system_menu(id, name, parent, sys_id, level_id, icon, path, enable, create_time, sortrank, is_open)VALUES (6, '系统管理', 3, 1, 3, 'fa fa-envelope text-success', '/sys/index', 1, '2018-6-28 11:06:03', 1, 1); 



INSERT INTO sso_user_info (user_id, full_name, user_name, password,  status, mobile ) VALUES (1, '管理员', 'guanly', 'e10adc3949ba59abbe56e057f20f883e',  0, '18000000000');

INSERT INTO sso_user_role ( id, user_id, sys_id, role_id, enable) VALUES (1, 1, 1, 1, 1);
