package model

//DataPermissionGetReq 获取用户有权限的　[数据权限]　数据
type DataPermissionGetReq struct {
	UserID        int    `form:"user_id" json:"user_id" valid:"required"`
	Ident         string `form:"ident" json:"ident" valid:"required"`
	TableName     string `form:"table_name" json:"table_name" valid:"required"`
	OperateAction string `form:"operate_action" json:"operate_action"`
}
