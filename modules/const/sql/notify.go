package sql


const QueryUserNotifyCount = `select count(1) from sso_notify_user where 1=1 ?title &sys_id &user_id`

const QueryUserNotifyPageList  = `
select 
t2.* 
from 
(select t.*,to_char(t.create_time, 'yyyy-MM-dd hh24:mi:ss') create_times,to_char(t.finish_time, 'yyyy-MM-dd hh24:mi:ss') finish_times,rownum as rn from sso_notify_user t 
	where t.title like '%'||@title||'%' 
	and t.user_id=@user_id 
	and t.sys_id=@sys_id
	and rownum < @pi * @ps) t2 
where 
t2.rn > (@pi - 1) * @ps`


const QueryUserNotifySetCount = `select count(1) from sso_notify_subscribe where user_id=@user_id and sys_id=@sys_id`

const QueryUserNotifySetPageList = `
select 
t2.* 
from 
(select t.*,to_char(t.create_time, 'yyyy-MM-dd hh24:mi:ss') create_times,rownum as rn from sso_notify_subscribe t 
	where t.user_id=@user_id 
	and t.sys_id=@sys_id
	and rownum < @pi * @ps) t2 
where 
t2.rn > (@pi - 1) * @ps`

const AddNotifySettings = `insert into sso_notify_subscribe(id,user_id,sys_id,level_id,keywords,status) 
values(seq_notify_subscribe_id.nextval,@user_id,@sys_id,@level_id,@keywords,@status)`

const DelNotifySettings = `delete from sso_notify_subscribe where id = @id`

const DelNotify = `delete from sso_notify_user where id = @id`

const EditNotifySettings = `update sso_notify_subscribe t
set  t.keywords = @keywords,t.level_id=@level_id,t.status=@status
where t.id=@id
`