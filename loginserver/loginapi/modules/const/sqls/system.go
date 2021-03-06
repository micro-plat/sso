package sqls

const SearchUserSystemInfo = `
select
  sys.id,
  sys.name,
  sys.index_url,
  sys.logo
from sso_system_info sys
inner join sso_user_role ur on ur.sys_id = sys.id
inner join sso_role_info role on role.role_id = ur.role_id
where role.status = 0 AND
    ur.enable=1 and 
    sys.index_url is not null and
    sys.index_url <> '' and
    ur.user_id = @user_id and
    sys.enable=1;
`
const QuerySysInfoByIdent = `
select 
  id, name, index_url, enable, ident 
from sso_system_info 
where ident = @ident 
limit 1;
`
