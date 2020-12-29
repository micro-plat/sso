package dictionary

//Get 获取某个类型下的字典信息
const Get = `
select 
	id,
	name,
	value,
	type,
	sort_no,
	group_code,
	status
from dds_dictionary_info
where type = @type and
	  status = 0
order by sort_no,id
`
