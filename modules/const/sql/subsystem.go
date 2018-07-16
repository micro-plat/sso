package sql

const QuerySubSystemListCount = `select count(*) from sso_system_info  where id >= @id`


const QuerySubSystemList = `select t.* from sso_system_info t where t.id >= @id`

const QuerySubSystemPageList = `
select 
t2.* 
from 
(select t1.*,rownum as rn from sso_system_info t1 where 1=1 and rownum < @page * @pageSize) t2 
where 
t2.rn > (@page - 1) * @pageSize`


const DeleteSubSystemById = `delete from sso_system_info where id = @id`

const QuerySubSystemListWithField = `select t.* from sso_system_info t 
where t.name like '%'||@name||'%' and t.enable=@enable`

const QuerySubSystemListAll = `select * from sso_system_info order by id`

const QuerySubSystemListByName = `select t.* from sso_system_info t where t.name like '%'||@name||'%'`

const QuerySubSystemListByEnable = `select t.* from sso_system_info t where t.enable=@enable`


const AddSubSystem  = `insert into sso_system_info(id,name,index_url,login_timeout,logo,theme,layout) 
values(seq_system_info_id.nextval,@name,@addr,@time_out,@logo,@theme,@style)`

const UpdateEnable = `update sso_system_info t
set  t.enable = @enable
where t.id=@id
`

const UpdateEdit = `update sso_system_info t
set  t.enable = @enable,t.index_url=@index_url,t.login_timeout=@login_timeout,t.logo=@logo,t.name=@name,t.theme=@theme,t.layout=@layout
where t.id=@id
`