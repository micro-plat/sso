
/*角色与数据权限数据的关联关系 */
create table sso_role_datapermission(
    id bigint primary key auto_increment   not null    comment '功能编号' ,
    sys_id bigint  not null    comment '系统编号' ,
    role_id bigint  not null    comment '角色编号' ,
    name VARCHAR(128) not null comment '名称',
    table_name VARCHAR(128) not null comment '表名',
    operate_action varchar(64)  not null    comment '操作动作',
    permissions VARCHAR(600)  not null    comment '数据规则信息json, sso_data_permission',
    status      tinyint  not null DEFAULT 0 comment '启用:0 禁用:1',
    create_time datetime DEFAULT CURRENT_TIMESTAMP not null    comment '创建时间' 
)COMMENT='功能信息';

-- alter table sso_role_datapermission add unique(sys_id,role_id,permission_id);
alter table sso_role_datapermission add index index_role_datapermission_sys_id(sys_id,role_id, table_name);