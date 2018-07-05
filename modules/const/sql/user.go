package sql

//QueryUserInfoList 查询用户信息列表
const QueryUserInfoList = `select TAB1.*
from (select L.*
		from (select rownum LINENUM, R.*
				from (select t.user_id,
							 t.user_name,
							 t.status,
							 t.mobile,
							 to_char(t.create_time, 'yyyy/mm/dd hh24:mi:ss') create_time,
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
			 to_char(t.create_time, 'yyyy/mm/dd hh24:mi:ss') create_time,
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
