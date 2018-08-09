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

const AddNotifySettings = `insert into sso_notify_subscribe(id,user_id,sys_id,level_id,keywords) 
values(seq_notify_subscribe_id.nextval,@user_id,@sys_id,@level_id,@keywords)`

const DelNotifySettings = `delete from sso_notify_subscribe where id = @id`

const DelNotify = `delete from sso_notify_user where id = @id`

const EditNotifySettings = `update sso_notify_subscribe t
set  t.keywords = @keywords,t.level_id=@level_id,t.status=@status
where t.id=@id
`

const InsertNotify = `insert into 
sso_notify_records(id,sys_id,level_id,title,keywords,content) 
values(seq_notify_records_id.nextval,@sys_id,@level_id,@title,@keywords,@content)`

const InsertNotifyUser = `INSERT INTO SSO_NOTIFY_USER (ID,USER_ID,SYS_ID,LEVEL_ID,TITLE,KEYWORDS,CONTENT,FLOW_TIMEOUT) 
SELECT SEQ_NOTIFY_USER_ID.nextval,USER_ID,SYS_ID,LEVEL_ID,@title,@keywords,@content,SYSDATE 
FROM SSO_NOTIFY_SUBSCRIBE
where KEYWORDS=@keywords AND SYS_ID=@sys_id AND LEVEL_ID=@level_id`


const UpdateNotifyUser = `update sso_notify_user t
set t.status = 2,t.scan_batch_id = @guid,t.send_count=t.send_count +1,t.flow_timeout = SYSDATE+1/1440
where t.send_count <= 3 and t.flow_timeout < SYSDATE and t.status in (1,2)
`

const QueryToUserNotify = `select t.id,t.title,t.content,t.create_time,u.wx_openid,s.name
from sso_notify_user t
left join sso_user_info u ON u.user_id = t.user_id
left join sso_system_info s ON s.id = t.sys_id 
where t.scan_batch_id=@guid`


const SendNotifyUserSucc = `update sso_notify_user t
set  t.status = 0,t.finish_time=SYSDATE
where t.id=@id
`