package province

//GetProvince 获取第一级省市
const GetProvince = `
select 
	canton_code,
	chinese_name,
	parent_code,
	grade,
	full_spell,
	simple_spell
from dds_area_info
where parent_code = @parent_code 
order by canton_code
`

//GetCityByProvinceID 根据省获取市信息
const GetCityByProvinceID = `
select 
	canton_code,
	chinese_name,
	parent_code,
	grade,
	full_spell,
	simple_spell
from dds_area_info
where 
and if(isnull(@parent_code)||@parent_code='',1=1,parent_code=@parent_code)
and grade='2'
order by canton_code
`

//GetAll 根据省获取市信息
const GetAll = `
select 
	canton_code,
	chinese_name,
	parent_code,
	grade,
	full_spell,
	simple_spell
from dds_area_info
order by canton_code
`
