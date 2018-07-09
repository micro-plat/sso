package sql

const QuerySubSystemListCount = `select count(*) from sso_system_info  where id >= @id`


const QuerySubSystemList = `select t.* from sso_system_info t where t.id >= @id`

const DeleteSubSystemById = `delete from sso_system_info where id = @id`

const QuerySubSystemListWithField = `select t.* from sso_system_info t where t.name like '%'||@name||'%' or t.enable=@status`

const AddSubSystem  = `insert into sso_system_info(id,name,index_url,login_timeout,logo) 
values(seq_system_info_id.nextval,@name,@addr,@time_out,@logo)`