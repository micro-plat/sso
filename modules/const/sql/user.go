package sql

//QueryUserInfoList 查询用户信息列表
const QueryUserInfoList = `select TAB1.*
from (select L.*
		from (select rownum LINENUM, R.*
				from (select to_char(t.user_id) user_id,
							 t.user_name,
							 t.status,
							 decode(t.status,0,'正常',1,'锁定',2,'禁用') status_label,
							 t.mobile,
							 to_char(t.create_time, 'yyyy/mm/dd hh24:mi') create_time,
							 t.email
						from sso_user_info t
						left join sso_user_role r on r.user_id = t.user_id
						where 1=1 
						and r.role_id = nvl(@role_id, r.role_id)
						#user_name
						group by t.user_id,
                                  t.user_name,
                                  t.status,
                                  t.mobile,
                                  t.email,
                                  t.create_time
					   order by t.user_id) R
			   where rownum <= @pi * @ps) L
	   where L.LINENUM > @ps * (@pi - 1)) TAB1
`

//QueryUserRoleList 查询用户角色信息列表
const QueryUserRoleList = `select to_char(a.user_id) user_id, to_char(a.sys_id) sys_id, to_char(a.role_id) role_id, s.name sys_name, r.name role_name
from sso_user_role a
inner join sso_system_info s on s.id = a.sys_id
inner join sso_role_info r on r.role_id = a.role_id
where a.user_id in
	 (select TAB1.user_id
		from (select L.*
				from (select rownum LINENUM, R.*
						from (select to_char(t.user_id) user_id
								from sso_user_info t
								left join sso_user_role r on r.user_id = t.user_id
							   where 1 = 1
								 and r.role_id = nvl(@role_id, r.role_id)
								 #user_name
							   group by t.user_id
							   order by t.user_id) R
					   where rownum <= @pi * @ps) L
			   where L.LINENUM > @ps * (@pi - 1)) TAB1)
order by a.user_id, a.sys_id, a.role_id
`

//QueryUserInfoListCount 获取用户信息列表数量
const QueryUserInfoListCount = `select count(1)
from (select to_char(t.user_id) user_id,
t.user_name,
t.status,
decode(t.status,0,'正常',1,'锁定',2,'禁用') status_label,
t.mobile,
to_char(t.create_time, 'yyyy/mm/dd hh24:mi') create_time,
t.email
from sso_user_info t
left join sso_user_role r on r.user_id = t.user_id
where 1=1 
and r.role_id = nvl(@role_id, r.role_id)
#user_name
group by t.user_id,
                                  t.user_name,
                                  t.status,
                                  t.mobile,
                                  t.email,
                                  t.create_time
order by t.user_id) R`

//UpdateUserStatus 获取用户信息列表数量
const UpdateUserStatus = `update sso_user_info t
set t.status = @status
where t.user_id = @user_id
`

//DeleteUser 删除用户
const DeleteUser = `delete from sso_user_info t where t.user_id = @user_id`

//QueryUserInfo 查询用户信息列表
const QueryUserInfo = `select t.user_id,t.user_name,t.mobile,t.email from sso_user_info t where t.user_id=@user_id`

//EditUserInfo 编辑用户信息
const EditUserInfo = `update sso_user_info t
set t.status = @status, t.user_name = @user_name, t.mobile = @mobile, t.email = @email
where t.user_id = @user_id
`

//DelUserRole 删除用户角色
const DelUserRole = `delete from sso_user_role t where t.user_id = @user_id`

//EditUserRole 编辑用户角色
const EditUserRole = `update sso_user_role t set t.role_id = @role_id where t.user_id = @user_id`

//GetNewUserID 获取新用户ID
const GetNewUserID = `select seq_user_info_id.nextval from dual`

//AddUserInfo 添加用户信息
const AddUserInfo = `insert into sso_user_info t
(user_id, user_name, status, password, mobile, email)
values
(@user_id, @user_name, @status, @password, @mobile, @email)
`

//AddUserRole 添加用户角色
const AddUserRole = `insert into sso_user_role
(id, user_id, role_id, sys_id, enable)
values
(seq_user_role_id.nextval, @user_id, @role_id, @sys_id, 1)
`

//QueryUserPswd 查询用户密码
const QueryUserPswd = `select count(1)
  from sso_user_info t
 where t.user_id=@user_id
 &password
`
const EditInfo = `update sso_user_info t
set  t.mobile = @tel, t.email = @email
where t.user_name = @username`

const QueryOldPwd = `select t.password from sso_user_info t where t.user_id=@user_id`

const SetNewPwd = `update sso_user_info t
set t.password = @password
where t.user_id = @user_id`

const QueryUserBind = `select t.email,t.wx_openid from sso_user_info t where t.email=@email`

const ExecUserBind = `update sso_user_info t
set t.wx_openid = @wx_openid
where t.email = @email`