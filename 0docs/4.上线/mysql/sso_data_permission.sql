
 /*drop table sso_data_permission;*/

CREATE TABLE sso_data_permission (
  id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '功能编号',
  sys_id bigint(20) NOT NULL COMMENT '系统编号',
  ident varchar(32) NOT NULL DEFAULT '' COMMENT '系统标识',
  name varchar(128) NOT NULL COMMENT '名称',
  table_name varchar(128) NOT NULL COMMENT '表名',
  operate_action varchar(64) NOT NULL COMMENT '操作动作',
  rules varchar(8000) DEFAULT NULL COMMENT '规则json',
  remark varchar(256) NOT NULL COMMENT '说明',
  status tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (id),
  UNIQUE KEY sys_id_name (sys_id,name)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COMMENT='数据权限规则数据';