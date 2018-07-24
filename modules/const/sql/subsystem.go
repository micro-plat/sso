package sql

const QuerySubSystemTotalCount = `select count(1) from sso_system_info where 1=1 ?name &enable`

const QuerySubSystemList = `select t.* from sso_system_info t where t.id >= @id`

const QuerySubSystemPageList = `
select 
t2.* 
from 
(select t.*,rownum as rn from sso_system_info t where t.name like '%'||@name||'%' &t.enable and rownum < @pi * @ps) t2 
where 
t2.rn > (@pi - 1) * @ps`

const DeleteSubSystemById = `delete from sso_system_info where id = @id`

const AddSubSystem = `insert into sso_system_info(id,name,index_url,login_timeout,logo,theme,layout,ident) 
values(seq_system_info_id.nextval,@name,@addr,@time_out,@logo,@theme,@style,@ident)`

const UpdateEnable = `update sso_system_info t
set  t.enable = @enable
where t.id=@id
`

const UpdateEdit = `update sso_system_info t
set  t.enable = @enable,t.index_url=@index_url,t.login_timeout=@login_timeout,t.logo=@logo,t.name=@name,t.theme=@theme,t.layout=@layout,t.ident=@ident
where t.id=@id
`
