package sqls

const SearchUserSystemInfo = `
select
  sys.id,
  sys.name,
  sys.index_url as indexurl,
  sys.logo,
  sys.callback_url as callbackurl
from sso_system_info sys
inner join sso_user_role ur on ur.sys_id = sys.id
inner join sso_role_info role on role.role_id = ur.role_id
where role.status = 0 AND
	  ur.user_id = @user_id; 
`
