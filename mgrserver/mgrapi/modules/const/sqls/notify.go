package sqls

//QueryUserNotifyCount .
const QueryUserNotifyCount = `
select 
	count(1) 
from 
	sso_notify_user 
where 
	1=1 ?title &sys_id &user_id`

//QueryUserNotifyPageList .
const QueryUserNotifyPageList = `
select 
	t.*,
	t.create_time create_times,
	t.finish_time finish_times
from 
	sso_notify_user t 
	where t.title like '%#title%' 
	and t.user_id=@user_id 
	and t.sys_id=@sys_id
limit 
	#pageSize offset #currentPage
`

//QueryUserNotifySetCount .
const QueryUserNotifySetCount = `
select 
	count(1) 
from 
	sso_notify_subscribe 
where 
	user_id=@user_id 
	and sys_id=@sys_id`

//QueryUserNotifySetPageList .
const QueryUserNotifySetPageList = `
select 
	t.*,
	t.create_time create_times
from 
	sso_notify_subscribe t 
where 
	t.user_id=@user_id 
	and t.sys_id=@sys_id
limit 
	#pageSize offset #currentPage
`

// AddNotifySettings .
const AddNotifySettings = `
insert into 
	sso_notify_subscribe
	(user_id,sys_id,level_id,keywords) 
values
	(@user_id,@sys_id,@level_id,@keywords)`

//DelNotifySettings .
const DelNotifySettings = `
delete from 
	sso_notify_subscribe 
where 
	id = @id 
	and user_id=@uid
`

//DelNotify .
const DelNotify = `
delete from 
	sso_notify_user 
where 
	id = @id 
	and user_id=@uid
`

//EditNotifySettings .
const EditNotifySettings = `
update 
	sso_notify_subscribe t
set  
	t.keywords = @keywords,t.level_id=@level_id,t.status=@status
where 
	t.id=@id
`

//InsertNotify .
const InsertNotify = `
insert into 
	sso_notify_records
	(sys_id,level_id,title,keywords,content) 
values
	(@sys_id,@level_id,@title,@keywords,@content)`

// InsertNotifyUser .
const InsertNotifyUser = `
insert into 
	sso_notify_user 
	(user_id,sys_id,level_id,title,keywords,content,flow_timeout) 
select 
	user_id,sys_id,level_id,@title,@keywords,@content,now() 
from 
	sso_notify_subscribe
where 
	keywords=@keywords 
	and sys_id=@sys_id 
	and (level_id=0 or level_id=@level_id)`

//UpdateNotifyUser .
const UpdateNotifyUser = `
update 
	sso_notify_user t
set 
	t.status = 2,
	t.scan_batch_id = @guid,
	t.send_count=t.send_count +1,
	t.flow_timeout = now() + 1/24/60
where 
	t.send_count <= 3 
	and t.flow_timeout < now() 
	and t.flow_timeout > now() - 5/24/60 
	and t.status in (1,2)`

//QueryToUserNotify .
const QueryToUserNotify = `
select 
	t.id,
	t.title,
	t.content,
	t.create_time create_times,
	u.wx_openid,
	s.name
from 
	sso_notify_user t
left join 
	sso_user_info u ON u.user_id = t.user_id
left join 
	sso_system_info s ON s.id = t.sys_id 
where 
	t.scan_batch_id=@guid 
	and to_seconds(t.flow_timeout) > to_seconds(now())`

//SendNotifyUserSucc .
const SendNotifyUserSucc = `
update 
	sso_notify_user t
set  
	t.status = 0,
	t.finish_time=now()
where 
	t.id=@id
`
