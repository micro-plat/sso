package sql

//QuerySubSystemTotalCount .
const QuerySubSystemTotalCount = `
select 
  count(1) 
from 
  sso_system_info 
where 
  #name_sql
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
  t.*
from 
<<<<<<< HEAD:apiserver/modules/const/sql/subsystem.go
(select t.*,rownum as rn from sso_system_info t 
  where t.name like '%'||@name||'%' &t.enable and rownum <= @pi * @ps) t2 
=======
  sso_system_info t 
>>>>>>> 750f5c63baeb3b4a71bc53caecd154a8e0ed6969:flowserver/modules/const/sql/subsystem.go
where 
  #name_sql
  &enable 
limit 
	#pageSize offset #currentPage
`

//DeleteSubSystemById .
const DeleteSubSystemById = `
delete from  
  sso_system_info 
where 
  id = @id
`

<<<<<<< HEAD:apiserver/modules/const/sql/subsystem.go
const AddSubSystem = `
insert into 
  sso_system_info
  (
    id,
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
    seq_system_info_id.nextval,
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
=======
//AddSubSystem .
const AddSubSystem = `
insert into 
  sso_system_info
  (name,index_url,login_timeout,logo,theme,layout,ident,login_url) 
values
  (@name,@addr,@time_out,@logo,@theme,@style,@ident,@login_url)`
>>>>>>> 750f5c63baeb3b4a71bc53caecd154a8e0ed6969:flowserver/modules/const/sql/subsystem.go

//UpdateEnable .
const UpdateEnable = `
update 
  sso_system_info t
set  
  t.enable = @enable
where 
  t.id=@id
`

<<<<<<< HEAD:apiserver/modules/const/sql/subsystem.go
=======
//UpdateEdit .
>>>>>>> 750f5c63baeb3b4a71bc53caecd154a8e0ed6969:flowserver/modules/const/sql/subsystem.go
const UpdateEdit = `
update 
  sso_system_info t
set  
  t.enable = @enable,
<<<<<<< HEAD:apiserver/modules/const/sql/subsystem.go
  t.index_url = @index_url,
  t.login_timeout = @login_timeout,
  t.logo = @logo,
  t.name = @name,
  t.theme = @theme,
  t.layout = @layout,
  t.ident = @ident,
  t.wechat_status = @wechat_status,
  t.secret = @secret 
where 
  t.id = @id
=======
  t.index_url=@index_url,
  t.login_timeout=@login_timeout,
  t.logo=@logo,
  t.name=@name,
  t.theme=@theme,
  t.layout=@layout,
  t.ident=@ident
where 
  t.id=@id
>>>>>>> 750f5c63baeb3b4a71bc53caecd154a8e0ed6969:flowserver/modules/const/sql/subsystem.go
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
