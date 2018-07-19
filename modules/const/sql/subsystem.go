package sql

const QuerySubSystemTotalCount = `select count(*) from sso_system_info  where 1=1  ?t.name  &t.enable`


const QuerySubSystemList = `select t.* from sso_system_info t where t.id >= @id`

const QuerySubSystemPageList = `
select 
t2.* 
from 
(select t1.*,rownum as rn from sso_system_info t1 where 1=1 ?t.name  &t.enable and rownum < @page * @pageSize) t2 
where 
t2.rn > (@page - 1) * @pageSize`


const DeleteSubSystemById = `delete from sso_system_info where id = @id`


const AddSubSystem  = `insert into sso_system_info(id,name,index_url,login_timeout,logo,theme,layout,ident) 
values(seq_system_info_id.nextval,@name,@addr,@time_out,@logo,@theme,@style,@ident)`

const UpdateEnable = `update sso_system_info t
set  t.enable = @enable
where t.id=@id
`

const UpdateEdit = `update sso_system_info t
set  t.enable = @enable,t.index_url=@index_url,t.login_timeout=@login_timeout,t.logo=@logo,t.name=@name,t.theme=@theme,t.layout=@layout
where t.id=@id
`