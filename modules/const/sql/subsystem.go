package sql

const QuerySubSystemListCount = `select count(*) from sso_system_info  where id >= @id`


const QuerySubSystemList = `select t.* from sso_system_info t where t.id >= @id`

const DeleteSubSystemById = `delete from sso_system_info where id = @id`

const QuerySubSystemListWithField = `select t.* from sso_system_info t where t.name like '%'||@name||'%' and t.enable=@enable`

const AddSubSystem  = `insert into sso_system_info(id,name,index_url,login_timeout,logo) 
values(seq_system_info_id.nextval,@name,@addr,@time_out,@logo)`

const UpdateEnable = `update sso_system_info t
set  t.enable = @enable
where t.id=@id
`

const UpdateEdit = `update sso_system_info t
set  t.enable = @enable,t.index_url=@index_url,t.login_timeout=@login_timeout,t.logo=@logo,t.name=@name
where t.id=@id
`