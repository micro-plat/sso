package sql

const QuerySubSystemTotalCount = `select count(*) from sso_system_info`

//const QuerySubSystem = `select t.* from sso_system_info t`

const QuerySubSystemPageList = `
select 
t2.* 
from 
(select t1.*,rownum as rn from sso_system_info t1 where 1=1 and rownum < @page * @pageSize) t2 
where 
t2.rn > (@page - 1) * @pageSize`

const DeleteSubSystemById = `delete from sso_system_info where id = @id`

const QuerySubSystemList = `select t.* from sso_system_info t 
where 1=1  ?t.name  &t.enable`

const AddSubSystem = `insert into sso_system_info(id,name,index_url,login_timeout,logo,theme,layout) 
values(seq_system_info_id.nextval,@name,@addr,@time_out,@logo,@theme,@style)`

const UpdateEnable = `update sso_system_info t
set  t.enable = @enable
where t.id=@id
`

const UpdateEdit = `update sso_system_info t
set  t.enable = @enable,t.index_url=@index_url,t.login_timeout=@login_timeout,t.logo=@logo,t.name=@name,t.theme=@theme,t.layout=@layout
where t.id=@id
`
