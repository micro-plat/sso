package sql

//QueryUserInfoList 查询用户信息列表
const QueryUserInfoList = `select TAB1.*
from (select L.*
		from (select rownum LINENUM, R.*
				from (select t.user_id,
							 t.user_name,
							 t.status,
							 decode(t.status,0,'正常',1,'锁定',2,'禁用') status_label,
							 t.mobile,
							 to_char(t.create_time, 'yyyy/mm/dd hh24:mi') create_time,
							 ri.name role_name,
							 r.sys_id,
							 r.role_id,
							 r.enable
						from sso_user_info t
						left join sso_user_role r on r.user_id = t.user_id
						left join sso_role_info ri on ri.role_id = r.role_id
						where 1=1 
						&r.role_id
						&user_name
					   order by t.user_id, r.role_id) R
			   where rownum <= @pi * @ps) L
	   where L.LINENUM > @ps * (@pi - 1)) TAB1
`

//QueryUserInfoListCount 获取用户信息列表数量
const QueryUserInfoListCount = `select count(1)
from (select t.user_id,
			 t.user_name,
			 t.status,
			 t.mobile,
			 to_char(t.create_time, 'yyyy/mm/dd hh24:mi') create_time,
			 ri.name role_name,
			 r.sys_id,
			 r.role_id,
			 r.enable
		from sso_user_info t
		left join sso_user_role r on r.user_id = t.user_id
		left join sso_role_info ri on ri.role_id = r.role_id
		 where 1=1 
			 &r.role_id
			 &user_name
	   order by t.user_id, r.role_id) R`

//UpdateUserStatus 获取用户信息列表数量
const UpdateUserStatus = `update sso_user_info t
set t.status = @status
where t.user_id = @user_id
and t.status = @ex_status
`

//DeleteUser 删除用户
const DeleteUser = `delete from sso_user_info t where t.user_id = @user_id`

//QueryUserInfo 查询用户信息列表
const QueryUserInfo = `select t.user_id,
       t.user_name,
       t.status,
       decode(t.status, 0, '正常', 1, '锁定', 2, '禁用') status_label,
       t.mobile,
       to_char(t.create_time, 'yyyy/mm/dd hh24:mi') create_time,
       ri.name role_name,
       r.sys_id,
       r.role_id,
       r.enable
  from sso_user_info t
  left join sso_user_role r on r.user_id = t.user_id
  left join sso_role_info ri on ri.role_id = r.role_id
 where 1 = 1 
       &user_id
`

//EditUserInfo 编辑用户信息
const EditUserInfo = `update sso_user_info t
set t.status = @status, t.user_name = @user_name, t.mobile = @mobile
where t.user_id = @user_id
`

//EditUserRole 编辑用户角色
const EditUserRole = `update sso_user_role t set t.role_id = @role_id where t.user_id = @user_id`

//GetNewUserID 获取新用户ID
const GetNewUserID = `select seq_user_info_id.nextval from dual`

//AddUserInfo 添加用户信息
const AddUserInfo = `insert into sso_user_info t
(user_id, user_name, status, password, mobile)
values
(@user_id, @user_name, @status, @password, @mobile)
`

//AddUserRole 添加用户角色
const AddUserRole = `insert into sso_user_role
(id, user_id, role_id, sys_id, enable)
values
(seq_user_role_id.nextval, @user_id, @role_id, 0, 1)
`

//QueryUserPswd 查询用户密码
const QueryUserPswd = `select t.password
  from sso_user_info t
 where 1 = 1 
       &user_id
`
