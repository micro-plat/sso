package oracle

const dds_area_info = `
create table dds_area_info(
    canton_code varchar2(32)  not null ,
    chinese_name varchar2(128)  not null ,
    parent_code varchar2(32)   ,
    grade number(2)   ,
    full_spell varchar2(64)   ,
    simple_spell varchar2(16)   ,
    sort_id number(11) default 0 not null,
    status number(2) default 0 not null 
    );


comment on table dds_area_info is '地区表';
comment on column dds_area_info.canton_code is '区域编号';	
comment on column dds_area_info.chinese_name is '中文名称';	
comment on column dds_area_info.parent_code is '父级编号';	
comment on column dds_area_info.grade is '行政级别';	
comment on column dds_area_info.full_spell is '英文/全拼';	
comment on column dds_area_info.simple_spell is '简拼';	
comment on column dds_area_info.sort_id is '排序';	
comment on column dds_area_info.status is '状态';	



alter table dds_area_info
add constraint pk_dds_area_info primary key(canton_code);
create index IDX_AREA_PARENT_CODE on dds_area_info(parent_code);
`
