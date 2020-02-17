package model

//DataPermissionReq 数据规则配置信息
type DataPermissionReq struct {
	ID            string `form:"id" json:"id" `
	Ident         string `form:"ident" json:"ident" `
	SysID         string `form:"sys_id" json:"sys_id" valid:"required"`
	Name          string `form:"name" json:"name" valid:"required"`             //名称
	TableName     string `form:"table_name" json:"table_name" valid:"required"` //表名
	OperateAction string `form:"operate_action" json:"operate_action"`          //操作动作
	Rules         string `form:"rules" json:"rules" valid:"required"`           //值
	Remark        string `form:"remark" json:"remark"`
}

//RolePermissionReq 角色规则关联信息
type RolePermissionReq struct {
	//ID     int    `form:"id" json:"id" `
	SysID  string `form:"sys_id" json:"sys_id" valid:"required"`
	RoleID string `form:"role_id" json:"role_id" valid:"required"`
	// Name   string `form:"name" json:"name" valid:"required"` //名称
	// TableName     string `form:"table_name" json:"table_name" valid:"required"`   //表
	// OperateAction string `form:"operate_action" json:"operate_action"`            //操作动作
	Permissions string `form:"permissions" json:"permissions" valid:"required"` //规则主键信息(sso_data_permission)
}

//RolePermissionQueryReq 查询角色规则关联信息
type RolePermissionQueryReq struct {
	SysID  string `form:"sys_id" json:"sys_id" valid:"required"`
	RoleID string `form:"role_id" json:"role_id" valid:"required"`
	// PageIndex int    `form:"pi" json:"pi" valid:"required"`
	// PageSize  int    `form:"ps" json:"ps" valid:"required"`
}
