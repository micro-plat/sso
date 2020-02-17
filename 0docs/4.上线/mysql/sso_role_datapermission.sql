
/*角色与数据权限数据的关联关系 */
CREATE TABLE sso_role_datapermission (
  id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '功能编号',
  sys_id bigint(20) NOT NULL COMMENT '系统编号',
  role_id bigint(20) NOT NULL COMMENT '角色编号',
  permission_config_id bigint(20) NOT NULL COMMENT '规则id',
  create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id),
  KEY index_role_datapermission_sys_id (sys_id,role_id)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COMMENT='角色与规则的关联信息表';