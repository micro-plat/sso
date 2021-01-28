package oracle

const dds_dictionary_info = `
create table dds_dictionary_info(
    id number(10)  not null ,
    name varchar2(128)  not null ,
    value varchar2(4000)  not null ,
    type varchar2(32)  not null ,
    sort_no number(2) default 0 not null ,
    status number(1)  not null ,
    group_code varchar2(32) default '*' not null 
    );


comment on table dds_dictionary_info is '字典表';
comment on column dds_dictionary_info.id is '编号';	
comment on column dds_dictionary_info.name is '名称';	
comment on column dds_dictionary_info.value is '值';	
comment on column dds_dictionary_info.type is '类型';	
comment on column dds_dictionary_info.sort_no is '排序值';	
comment on column dds_dictionary_info.status is '状态';	
comment on column dds_dictionary_info.group_code is '分组编号';	



alter table dds_dictionary_info
add constraint pk_dds_dictionary_info primary key(id);
create index IDX_DICTIONARY_INFO_TYPE on dds_dictionary_info(type);

create sequence seq_dictionary_info_id
increment by 1
minvalue 1
maxvalue 99999999999
start with 1
cache 20;
`
