
/*角色与数据权限数据的关联关系 */
create table sso_role_datapermission(
    id bigint primary key auto_increment   not null    comment '功能编号' ,
    sys_id bigint  not null    comment '系统编号' ,
    role_id bigint  not null    comment '角色编号' ,
    permission_id bigint  not null    comment '数据权限编号' ,
    create_time datetime DEFAULT CURRENT_TIMESTAMP not null    comment '创建时间' 
)COMMENT='功能信息';

alter table sso_role_datapermission add unique(sys_id,role_id,permission_id);
alter table sso_role_datapermission add index index_role_datapermission_sys_id(sys_id,role_id);