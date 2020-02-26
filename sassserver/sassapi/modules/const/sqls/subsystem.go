package sqls

//QuerySubSystemTotalCount .
const QuerySubSystemTotalCount = `
select 
  count(1) 
from 
  sso_system_info 
where 
  1 = 1
  #name
  &enable
`

//QuerySubSystemList .
const QuerySubSystemList = `
select 
  t.* 
from 
  sso_system_info t 
where 
  t.id >= @id`

//QuerySubSystemPageList .
const QuerySubSystemPageList = `
select 
  id, name, index_url as callbackurl, enable, login_timeout, 
  logo, theme, layout, ident, login_url, wechat_status 
from 
  sso_system_info t 
where 1 = 1 
  #name
  &enable
limit @start, @ps;
`

//DeleteSubSystemById .
const DeleteSubSystemById = `
delete from  
  sso_system_info 
where 
  id = @id
`

const ExistsNameOrIdent = `select 
  count(1) as count 
from sso_system_info 
where name=@name or ident=@ident`

const AddSubSystem = `
insert into 
  sso_system_info
  (
    name,
    index_url,
    login_timeout,
    logo,
    theme,
    layout,
    ident,
    wechat_status,
    login_url,
    secret
  ) 
values
  (
    @name,
    @addr,
    @time_out,
    @logo,
    @theme,
    @style,
    @ident,
    @wechat_status,
    @login_url,
    @secret
  )`

//UpdateEnable .
const UpdateEnable = `
update 
  sso_system_info t
set  
  t.enable = @enable
where 
  t.id=@id
`

//UpdateEdit .
const UpdateEdit = `
update 
  sso_system_info t
set  
  t.enable = @enable,
  t.index_url = @index_url,
  t.login_timeout = @login_timeout,
  t.logo = @logo,
  t.name = @name,
  t.theme = @theme,
  t.layout = @layout,
  t.ident = @ident,
  t.wechat_status = @wechat_status
where 
  t.id = @id
`

const QuerySsoSystemMenu = `
select t.*
  from sso_system_menu t
where t.sys_id = @sys_id
   and t.level_id = @level_id
   and t.parent =@parent
   #sortrank
#orderby
limit 1;`

const UpSsoSystemMenu = `
update sso_system_menu t
   set t.sortrank = @sortrank
 where t.sys_id = @sys_id
   and t.level_id = @level_id
   and t.id = @id
`

// GetUsers .
const GetUsers = `
select
  r.user_id,u.user_name
from 
  sso_system_info i
inner join 
  sso_user_role r on r.sys_id = i.id
inner join 
  sso_user_info u on u.user_id = r.user_id
where 
  i.ident=@system_name 
order by 
  r.user_id
`

// GetAllUser .
const GetAllUser = `
select 
  distinct r.user_id,u.user_name
from 
  sso_system_info i
inner join 
  sso_user_role r on r.sys_id = i.id
inner join 
  sso_user_info u on u.user_id = r.user_id
where 
  i.id>=0  
order by 
  r.user_id
`

//ChangeSecret 修改秘钥
const ChangeSecret = `update sso_system_info set secret = @secret  where id = @id limit 1`

const GetSystemInfoByIdent = `
SELECT 
  id,
  name,
  enable
from sso_system_info
where ident = @ident
`
